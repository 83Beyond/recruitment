package Kwai

import (
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
)

type Resp struct {
	Data    Data   `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	List []PositionInfo `json:"job_post_list"`
}

type PositionInfo struct {
	Name     string   `json:"title"`
	City     string   `json:"workplace"`
	Id       string   `json:"id"`
	CityInfo CityInfo `json:"city_info"`
}

type CityInfo struct {
	Name string `json:"name"`
}

func Parse(resp network.Response) {
	prefix := "https://jobs.bytedance.com/campus/position/"
	var r Resp
	_ = json.Unmarshal(resp.Content, &r)
	for _, position := range r.Data.List {
		positionInfo := models.PositionInfo{
			Name:   position.Name,
			City:   position.CityInfo.Name,
			Link:   prefix + position.Id + "/detail",
			Source: resp.Source,
		}
		fmt.Println(positionInfo)
		models.PositionInfoChan <- positionInfo
	}
}
