package websocket

import "qq-bot/integration/dto"

const (
	EventReady           dto.EventType = "READY"
	EventAtMessageCreate dto.EventType = "AT_MESSAGE_CREATE"
)

func parseAndHandleEvent(event *dto.WSPayload) error {
	if event.OpCode == dto.OPCodeDispatch {
		switch event.Type {
		case EventAtMessageCreate:
			return atMessageHandler(event, event.RawMessage)
		}
	}

	return nil
}

func atMessageHandler(event *dto.WSPayload, message []byte) error {
	data := &dto.AtMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}

	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(event, data)
	}
	return nil
}
