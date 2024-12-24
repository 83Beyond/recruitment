package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func newTop() *fyne.Container {
	top := make([]fyne.CanvasObject, 0)

	searchBtn := widget.NewButtonWithIcon("搜索", getSearchIcon(), func() {})
	searchBtn.Importance = widget.HighImportance
	searchBtn.Resize(fyne.NewSize(BtnNormalWidth, BtnNormalHeight))
	top = append(top, searchBtn)
	fmt.Println(searchBtn.MinSize().Width)

	citySelect := widget.NewSelect([]string{
		"北京", "广州", "上海", "深圳", "杭州", "成都", "乌鲁木齐",
	}, func(city string) {
		fmt.Println(city)
	})
	citySelect.Selected = "北京"
	citySelect.Resize(fyne.NewSize(SelectNormalWidth, SelectNormalHeight))
	citySelect.Move(fyne.NewPos(BtnNormalWidth+NormalSpacing, 0))
	top = append(top, citySelect)

	typeSelect := widget.NewSelect([]string{"校园招聘", "社会招聘", "日常实习"}, func(t string) {})
	typeSelect.Selected = "校园招聘"
	typeSelect.Resize(fyne.NewSize(SelectNormalWidth, SelectNormalHeight))
	typeSelect.Move(fyne.NewPos(BtnNormalWidth+SelectNormalWidth+2*NormalSpacing, 0))
	top = append(top, typeSelect)

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("请输入关键词")
	positionEntry.Resize(fyne.NewSize(EntryNormalWidth, EntryNormalHeight))
	positionEntry.Move(fyne.NewPos(BtnNormalWidth+2*SelectNormalWidth+3*NormalSpacing, 0))
	top = append(top, positionEntry)

	line := canvas.NewLine(color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	line.Move(fyne.NewPos(0, BtnNormalHeight+NormalSpacing))
	spacer := canvas.NewRectangle(color.Transparent)
	spacer.Resize(fyne.NewSize(0, 5))

	return container.NewVBox(container.NewWithoutLayout(top...), spacer, line)
}
