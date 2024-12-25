package models

import "github.com/83Beyond/recruitment/utils"

type PositionInfo struct {
	Name   string
	Link   string
	City   string
	Source string
}

var (
	PositionList     []PositionInfo
	ShowPositionList []PositionInfo
	PositionInfoChan chan PositionInfo
	SelectedCity     string
	SelectedType     string
	Keyword          string
)

var (
	CitySelectList = []string{"全部", "北京", "广州", "上海", "深圳", "杭州", "成都", "乌鲁木齐"}
	TypeSelectList = []string{"校园招聘", "日常实习"}
	AllCity        = []string{"全部"}
	AllSource      = []string{"全部"}
	AllCitySet     = utils.NewSet()
	AllSourceSet   = utils.NewSet()

	FirstCity     = []string{"全部"}
	CityMap       = map[string][]string{}
	FirstCitySet  = utils.NewSet()
	SecondCitySet = utils.NewSet()

	FirstCitySelected string

	ScrapPage = 1
)

func InitData() {
	PositionList = make([]PositionInfo, 0)
	ShowPositionList = make([]PositionInfo, 0)
	PositionInfoChan = make(chan PositionInfo, 1000)
}

func RefreshData() {
	PositionList = make([]PositionInfo, 0)
	ShowPositionList = make([]PositionInfo, 0)
	AllCity = []string{"全部"}
	AllSource = []string{"全部"}
	AllCitySet = utils.NewSet()
	AllSourceSet = utils.NewSet()
}
