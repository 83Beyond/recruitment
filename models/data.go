package models

type PositionInfo struct {
	Name string
	Link string
	City string
}

var (
	PositionList []PositionInfo
)

func InitData() {
	PositionList = make([]PositionInfo, 0)
	PositionList = append(PositionList, PositionInfo{
		Name: "【2025校招】Android开发工程师-直播体验",
		Link: "https://www.baidu.com",
		City: "北京",
	})
	PositionList = append(PositionList, PositionInfo{
		Name: "【2025校招】Android开发工程师-直播体验【2025校招】Android开发工程师-直播体验",
		Link: "https://www.baidu.com",
		City: "长沙",
	})
}
