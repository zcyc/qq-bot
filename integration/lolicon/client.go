package lolicon

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

func Get() string {
	uri := "https://api.lolicon.app/setu/v2?r18=0"
	res, err := resty.New().R().Get(uri)
	if err != nil {
		return "年轻人要学会节制。"
	}
	body := &loliconBody{}
	_ = json.Unmarshal(res.Body(), body)
	log.Printf("%+v", body)
	return body.Data[0].Urls.Original
}

type loliconBody struct {
	Error string `json:"error"`
	Data  []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		Uid        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		R18        bool     `json:"r18"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Tags       []string `json:"tags"`
		Ext        string   `json:"ext"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original"`
		} `json:"urls"`
	} `json:"data"`
}
