package XHS

import (
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
	"strconv"
)

type Resp struct {
	Data      Data   `json:"data"`
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

type Data struct {
	List []PositionInfo `json:"list"`
}

type PositionInfo struct {
	Name string `json:"positionName"`
	City string `json:"workplace"`
	Id   int    `json:"positionId"`
}

func Parse(resp network.Response) {
	prefix := "https://job.xiaohongshu.com/campus/position/"
	var r Resp
	_ = json.Unmarshal(resp.Content, &r)
	for _, position := range r.Data.List {
		positionInfo := models.PositionInfo{
			Name:   position.Name,
			City:   position.City,
			Link:   prefix + strconv.Itoa(position.Id),
			Source: resp.Source,
		}
		fmt.Println(positionInfo)
		models.PositionInfoChan <- positionInfo
	}
}
