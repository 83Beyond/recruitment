package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/83Beyond/recruitment/models"
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
	positionData := newShowList()
	dataContent := container.NewBorder(positionHeader, nil, nil, nil, positionData)
	return dataContent
}

func newHeader() *fyne.Container {
	nameHeader := widget.NewLabel("职位名称")
	nameHeader.Resize(fyne.NewSize(NameHeaderWidth, NormalHeight))
	nameHeader.Move(fyne.NewPos(MoveZero, MoveZero))
	cityHeader := widget.NewLabel("工作地点")
	cityHeader.Resize(fyne.NewSize(CityHeaderWidth, NormalHeight))
	cityHeader.Move(fyne.NewPos(NameHeaderWidth, MoveZero))
	//pubTimeHeader := widget.NewLabel("发布时间")
	//pubTimeHeader.Resize(fyne.NewSize(ContentWidth/4, 50))
	//pubTimeHeader.Move(fyne.NewPos(ContentWidth/2+ContentWidth/4, 0))

	header := container.NewWithoutLayout(nameHeader, cityHeader)
	line := canvas.NewLine(Gray)

	return container.NewVBox(header, line)
}

func newShowList() *widget.List {
	showList := widget.NewList(
		func() int {
			return len(models.PositionList)
		},
		func() fyne.CanvasObject {
			name := widget.NewHyperlink("职位名称", nil)
			name.Alignment = fyne.TextAlignCenter

			nameBox := container.NewHBox(name)
			nameBox.Resize(fyne.NewSize(NameHeaderWidth, NormalHeight))
			nameBox.Move(fyne.NewPos(MoveZero, MoveZero))

			city := widget.NewLabel("工作地点")
			cityBox := container.NewHBox(city)
			cityBox.Resize(fyne.NewSize(CityHeaderWidth, NormalHeight))
			cityBox.Move(fyne.NewPos(NameHeaderWidth, MoveZero))
			//
			//pubTime := NewCenterLabel("pubTime")
			//pubTimeBox := container.NewHBox(pubTime)
			//pubTimeBox.Resize(fyne.NewSize(ContentWidth/4, 50))
			//pubTimeBox.Move(fyne.NewPos(ContentWidth/2+ContentWidth/4, 0))

			return container.NewWithoutLayout(
				nameBox,
				cityBox)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			grid := item.(*fyne.Container)

			titleBox := grid.Objects[0].(*fyne.Container)
			cityBox := grid.Objects[1].(*fyne.Container)
			//pubTimeBox := grid.Objects[2].(*fyne.Container)

			link := titleBox.Objects[0].(*widget.Hyperlink)

			// 截断显示
			maxLength := 26
			title := models.PositionList[id].Name
			fullTile := []rune(models.PositionList[id].Name)
			if len(fullTile) > maxLength {
				title = string(fullTile[:maxLength]) + "..."
			}
			link.SetText(title)
			link.SetURLFromString(models.PositionList[id].Link)

			city := cityBox.Objects[0].(*widget.Label)
			city.SetText(models.PositionList[id].City)
			//
			//pubTime := pubTimeBox.Objects[0].(*widget.Label)
			//pubTime.SetText(models.PositionList[id].PubTime)
		},
	)
	return showList
}
