package PAIC

import (
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
)

type Resp struct {
	Data         Data   `json:"data"`
	ResponseCode string `json:"responseCode"`
	ResponseMsg  int    `json:"responseMsg"`
}

type Data struct {
	List []PositionInfo `json:"list"`
}

type PositionInfo struct {
	Name     string   `json:"deptShowName"`
	LocNames []string `json:"LocNames"`
	City     string   `json:"workCity"`
	Id       string   `json:"idPosition"`
}

func Parse(resp network.Response) {
	prefix := "https://campus.pingan.com/positionDetail?"
	var r Resp
	_ = json.Unmarshal(resp.Content, &r)
	for _, position := range r.Data.List {
		positionInfo := models.PositionInfo{
			Name:   position.Name,
			City:   position.City,
			Link:   prefix + "positionId=" + position.Id,
			Source: resp.Source,
		}
		fmt.Println(positionInfo)
		models.PositionInfoChan <- positionInfo
	}
}
