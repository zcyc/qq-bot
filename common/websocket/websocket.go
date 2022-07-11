package websocket

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"qq-bot/integration/dto"
	"qq-bot/integration/qq"
)

func parseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")

	return json.Unmarshal([]byte(data.String()), target)
}

// Session websocket 链接所需要的会话信息
type Session struct {
	ID      string // 鉴权完成后由 QQ 服务器下发
	LastSeq uint32 // 消息序列号, 保存用于发送心跳时使用
	URL     string
	Token   qq.Token
	Intent  dto.Intent
}
