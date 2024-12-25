package ui

import (
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/spider"
	"strings"
)

func searchBtnCallback() {
	fmt.Println("刷新数据")
	models.RefreshData()
	scheduler := spider.Scheduler{}
	fmt.Println("开始调度")
	scheduler.Schedule()
	scheduler.Start()
}

func citySelectCallback(c string) {
	models.SelectedCity = c
}

func typeSelectCallback(t string) {
	models.SelectedType = t
}

func positionEntryCallback(keyword string) {
	models.Keyword = keyword
}

func sourceHeaderSelectCallback(source string) {
	if source == "全部" {
		models.ShowPositionList = models.PositionList
	} else {
		models.ShowPositionList = make([]models.PositionInfo, 0)
		for _, position := range models.PositionList {
			if position.Source == source {
				models.ShowPositionList = append(models.ShowPositionList, position)
			}
		}
	}
	ShowList.Refresh()
}

func firstCityHeaderCallback(first string) {

	if first == "全部" {
		SecondCityHeader.Options = []string{"全部"}
		models.ShowPositionList = models.PositionList
	} else {
		models.FirstCitySelected = first
		SecondCityHeader.Options = models.CityMap[first]
		models.ShowPositionList = make([]models.PositionInfo, 0)
		for _, position := range models.PositionList {
			city := strings.Split(position.City, "-")
			if city[0] == first {
				models.ShowPositionList = append(models.ShowPositionList, position)
			}
		}
	}
	SecondCityHeader.Selected = "全部"
	SecondCityHeader.Refresh()
	ShowList.Refresh()
}

func secondCityHeaderCallback(second string) {
	models.ShowPositionList = make([]models.PositionInfo, 0)
	if second == "全部" {
		for _, position := range models.PositionList {
			city := strings.Split(position.City, "-")
			if city[0] == models.FirstCitySelected {
				models.ShowPositionList = append(models.ShowPositionList, position)
			}
		}
	} else {
		for _, position := range models.PositionList {
			city := strings.Split(position.City, "-")
			if city[0] == models.FirstCitySelected && city[1] == second {
				models.ShowPositionList = append(models.ShowPositionList, position)
			}
		}
	}

	ShowList.Refresh()
}
