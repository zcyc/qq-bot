package at_message

import (
	"context"
	"log"
	"qq-bot/common/websocket"
	"qq-bot/integration/dto"
	"qq-bot/integration/hitokoto"
	"qq-bot/integration/qq"
	"strings"
)

func NewHandler(botClient *qq.Bot) dto.Intent {
	ctx := context.Background()
	var atMessage websocket.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.AtMessageData) error {
		if strings.HasSuffix(data.Content, "> 一言") {
			s := hitokoto.Get()
			_, _ = botClient.SendMessage(ctx, data.ChannelID, &dto.ChannelMessage{MsgID: data.ID, Content: s})
		}
		if strings.HasSuffix(data.Content, "> 图来") {
			//p := lolicon.Get()
			p := "https://c.wallhere.com/photos/bc/64/snow_snowflake_winter-626454.jpg!d"
			res, err := botClient.SendMessage(ctx, data.ChannelID, &dto.ChannelMessage{MsgID: data.ID, Image: p})
			if err != nil {
				log.Printf("PostMessage error %s", err.Error())
			}
			log.Printf("PostMessage result %+v", res)
		}
		return nil
	}
	intent := websocket.RegisterHandlers(atMessage) // 注册socket消息处理

	return intent
}
