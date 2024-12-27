package ChinaLife

import (
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
	"strings"
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
	prefix := "https://chinalife.zhiye.com/custom/zwxq"
	var r Resp
	_ = json.Unmarshal(resp.Content, &r)
	hideMenu := "1"
	for _, position := range r.Data {
		var payload = make([]string, 0)
		payload = append(payload, "hideMenu"+"="+hideMenu)
		payload = append(payload, "jobAdId"+"="+position.Id)
		payload = append(payload, "jc"+"="+position.Category)
		payload = append(payload, "ClassificationTwo"+"="+position.ClassificationTwo)
		payload = append(payload, "ClassificationOne"+"="+position.ClassificationOne)

		positionInfo := models.PositionInfo{
			Name:   position.Name,
			City:   position.LocNames[0],
			Link:   prefix + "?" + strings.Join(payload, "&"),
			Source: resp.Source,
		}
		fmt.Println(positionInfo)
		models.PositionInfoChan <- positionInfo
	}
}
