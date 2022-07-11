package websocket

import (
	"log"
	"qq-bot/integration/dto"
	"qq-bot/integration/qq"
	"runtime"
)

// SessionManager 用于启动和管理 websocket 链接
type SessionManager struct {
	sessionChan chan Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

func (m *SessionManager) Start(ap *dto.AccessPoint, token *qq.Token, intent dto.Intent) error {
	// TODO 支持分片
	m.sessionChan = make(chan Session, 1)
	m.sessionChan <- Session{
		URL:    ap.URL,
		Token:  *token,
		Intent: intent,
	}

	for session := range m.sessionChan {
		go m.newConnect(session)
	}

	return nil
}

func (m *SessionManager) newConnect(session Session) {
	defer func() {
		if err := recover(); err != nil {
			panicHandler(err, session)
		}
	}()

	client := NewWSClient(&session)
	if err := client.Connect(); err != nil {
		log.Println(err)
		return
	}

	var err error
	if session.ID == "" { // 初次连接
		err = client.Identify()
	} else {
		err = client.Resume()
	}
	if err != nil {
		log.Println(err)
		return
	}

	if err := client.Listening(); err != nil {
		curSession := client.session
		log.Printf("client.Listening has err=%v", err)

		if IsNeedReIdentifyError(err) { // 重新鉴权, 需要清空 session 和 lastSeq 信息
			curSession.LastSeq = 0
			curSession.ID = ""
		}

		if IsNeedPanicError(err) { // 无法重新鉴权的错误, 已经无法恢复了, 只能 panic
			log.Printf("the connect can't re-identify err=%v", err)
			panic(err)
		}

		// 将 session 发送回 session chan 中重新使用
		m.sessionChan <- *curSession
	}
}

func panicHandler(e interface{}, session Session) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]

	log.Printf("[PANIC]session=[%#v], err=[%v], stack=%s", session, e, buf)
}
