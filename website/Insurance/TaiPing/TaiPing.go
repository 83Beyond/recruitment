package TaiPing

import (
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
)

type Resp struct {
	Data    []PositionInfo `json:"Data"`
	Message string         `json:"Message"`
	Code    int            `json:"Code"`
}

type PositionInfo struct {
	Name              string   `json:"JobAdName"`
	LocNames          []string `json:"LocNames"`
	City              string   `json:"City"`
	Id                string   `json:"Id"`
	Category          string   `json:"Category"`
	ClassificationTwo string   `json:"ClassificationTwo"`
	ClassificationOne string   `json:"ClassificationOne"`
}

func Parse(resp network.Response) {
	var prefix = "https://cntp.zhiye.com/intern/detail?"
	var r Resp
	_ = json.Unmarshal(resp.Content, &r)
	for _, position := range r.Data {
		positionInfo := models.PositionInfo{
			Name:   position.Name,
			City:   position.LocNames[0],
			Link:   prefix + "jobAdId=" + position.Id,
			Source: resp.Source,
		}
		fmt.Println(positionInfo)
		models.PositionInfoChan <- positionInfo
	}
}
