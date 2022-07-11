package qq

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"qq-bot/integration/dto"
	"time"
)

// 接口域名
const (
	Domain        string = "api.sgroup.qq.com"         // 正式环境
	DomainSandbox string = "sandbox.api.sgroup.qq.com" // 沙盒环境
)

// 接口 URI
const (
	WebsocketGatewayUri string = "/gateway"                        // Websocket 接口
	SendMessageUri      string = "/channels/{channel_id}/messages" // 发送 ARK 模板消息接口
	UserInfoUri         string = "/users/@me"                      // 获取用户详情接口
)

type Bot struct {
	token       *Token
	sandbox     bool
	restyClient *resty.Client
}

func NewBotClient(sandBoxMode bool, token *Token) *Bot {
	return &Bot{
		token:   token,
		sandbox: sandBoxMode,
		restyClient: resty.New().
			SetTimeout(time.Second * 10).
			SetAuthScheme(token.Scheme).
			SetAuthToken(token.GetString()),
	}
}

func (c *Bot) request(ctx context.Context) *resty.Request {
	return c.restyClient.R().SetContext(ctx)
}

func (c *Bot) getURL(endpoint string) string {
	d := Domain
	if c.sandbox {
		d = DomainSandbox
	}

	return fmt.Sprintf("https://%s%s", d, endpoint)
}

// GetAccessPoint 获取 websocket 网关
func (c *Bot) GetAccessPoint(ctx context.Context) (*dto.AccessPoint, error) {
	resp, err := c.request(ctx).
		SetResult(dto.AccessPoint{}).
		Get(c.getURL(WebsocketGatewayUri))
	if err != nil {
		return nil, err
	}
	log.Println(resp)

	return resp.Result().(*dto.AccessPoint), nil
}

// SendMessage 发送消息
func (c *Bot) SendMessage(ctx context.Context, channelID string, msg *dto.ChannelMessage) (*dto.Message, error) {
	resp, err := c.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(c.getURL(SendMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.Message), nil
}

// Me 拉取当前用户的信息
func (c *Bot) Me(ctx context.Context) (*dto.User, error) {
	resp, err := c.request(ctx).
		SetResult(dto.User{}).
		Get(c.getURL(UserInfoUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.User), nil
}
