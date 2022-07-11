package hitokoto

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func Get() string {
	uri := "https://v1.hitokoto.cn"
	res, err := resty.New().R().Get(uri)
	if err != nil {
		return "没什么好说的。"
	}
	body := &hitokotoBody{}
	_ = json.Unmarshal(res.Body(), body)
	return fmt.Sprintf("『%s』\n来自：%s", body.Hitokoto, body.From)
}

type hitokotoBody struct {
	Id         int         `json:"id"`
	Uuid       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUid int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}
