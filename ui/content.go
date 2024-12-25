package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/83Beyond/recruitment/models"
)

var (
	ShowList         *widget.List
	FirstCityHeader  *widget.Select
	SecondCityHeader *widget.Select
	SourceHeader     *widget.Select
)

func newContent() *container.AppTabs {
	checkCompany := []string{"字节跳动", "美团", "百度", "快手", "小红书"}
	// 创建多个复选框
	checks := []fyne.CanvasObject{}
	for _, c := range checkCompany {
		check := widget.NewCheck(c, func(bool) {})
		checks = append(checks, check)
	}
	company := container.NewBorder(nil, nil, nil, nil, checks...)
	positionList := ShowPositionList()

	tab1 := container.NewTabItem("职位信息", positionList)
	tab2 := container.NewTabItem("已选公司", company)
	tabs := container.NewAppTabs(tab1, tab2)
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}

func ShowPositionList() *fyne.Container {
	positionHeader := newHeader()
	ShowList = newShowList()
	dataContent := container.NewBorder(positionHeader, nil, nil, nil, ShowList)
	return dataContent
}

func newCenterLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.Alignment = fyne.TextAlignCenter
	return label
}

func newHeader() *fyne.Container {
	nameHeader := newCenterLabel("职位名称")
	nameHeader.Resize(fyne.NewSize(NameHeaderWidth, NormalHeight))
	nameHeader.Move(fyne.NewPos(MoveZero, MoveZero))

	FirstCityHeader = widget.NewSelect(models.AllCity, firstCityHeaderCallback)
	FirstCityHeader.Selected = "工作"
	FirstCityHeader.Alignment = fyne.TextAlignCenter
	FirstCityHeader.Resize(fyne.NewSize(FirstCityHeaderWidth, NormalHeight))
	FirstCityHeader.Move(fyne.NewPos(NameHeaderWidth, MoveZero))

	SecondCityHeader = widget.NewSelect(models.AllCity, secondCityHeaderCallback)
	SecondCityHeader.Selected = "地点"
	SecondCityHeader.Alignment = fyne.TextAlignCenter
	SecondCityHeader.Resize(fyne.NewSize(SecondCityHeaderWidth, NormalHeight))
	SecondCityHeader.Move(fyne.NewPos(NameHeaderWidth+FirstCityHeaderWidth, MoveZero))

	SourceHeader = widget.NewSelect(models.AllSource, sourceHeaderSelectCallback)
	SourceHeader.Selected = "来源"
	SourceHeader.Alignment = fyne.TextAlignCenter
	SourceHeader.Resize(fyne.NewSize(SourceHeaderWidth, NormalHeight))
	SourceHeader.Move(fyne.NewPos(NameHeaderWidth+FirstCityHeaderWidth+SecondCityHeaderWidth, MoveZero))

	header := container.NewWithoutLayout(nameHeader, FirstCityHeader, SecondCityHeader, SourceHeader)
	line := canvas.NewLine(Gray)

	return container.NewVBox(header, line)
}

func newShowList() *widget.List {
	showList := widget.NewList(
		func() int {
			return len(models.ShowPositionList)
		},
		func() fyne.CanvasObject {
			name := widget.NewHyperlink("职位名称", nil)
			name.Alignment = fyne.TextAlignCenter

			nameBox := container.NewHScroll(name)
			nameBox.Resize(fyne.NewSize(NameHeaderWidth, NormalHeight))
			nameBox.Move(fyne.NewPos(MoveZero, MoveZero))

			city := newCenterLabel("工作地点")
			cityBox := container.NewHScroll(city)
			cityBox.Resize(fyne.NewSize(FirstCityHeaderWidth+SecondCityHeaderWidth, NormalHeight))
			cityBox.Move(fyne.NewPos(NameHeaderWidth, MoveZero))

			source := newCenterLabel("来源")
			sourceBox := container.NewHScroll(source)
			sourceBox.Resize(fyne.NewSize(SourceHeaderWidth, NormalHeight))
			sourceBox.Move(fyne.NewPos(NameHeaderWidth+FirstCityHeaderWidth+SecondCityHeaderWidth, MoveZero))

			return container.NewWithoutLayout(
				nameBox,
				cityBox,
				sourceBox)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			grid := item.(*fyne.Container)

			titleBox := grid.Objects[0].(*container.Scroll)
			cityBox := grid.Objects[1].(*container.Scroll)
			sourceBox := grid.Objects[2].(*container.Scroll)

			link := titleBox.Content.(*widget.Hyperlink)

			//// 截断显示
			//maxLength := 26
			title := models.ShowPositionList[id].Name
			//fullTile := []rune(models.ShowPositionList[id].Name)
			//if len(fullTile) > maxLength {
			//	title = string(fullTile[:maxLength]) + "..."
			//}
			link.SetText(title)
			link.SetURLFromString(models.ShowPositionList[id].Link)

			city := cityBox.Content.(*widget.Label)
			city.SetText(models.ShowPositionList[id].City)

			source := sourceBox.Content.(*widget.Label)
			source.SetText(models.ShowPositionList[id].Source)
		},
	)
	return showList
}
