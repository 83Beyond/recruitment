package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/83Beyond/recruitment/models"
	"image/color"
)

func newTop() *fyne.Container {
	top := make([]fyne.CanvasObject, 0)

	searchBtn := widget.NewButtonWithIcon("搜索", getSearchIcon(), searchBtnCallback)
	searchBtn.Importance = widget.HighImportance
	searchBtn.Resize(fyne.NewSize(BtnNormalWidth, BtnNormalHeight))
	top = append(top, searchBtn)

	citySelect := widget.NewSelect(models.CitySelectList, citySelectCallback)
	citySelect.Selected = models.CitySelectList[0]
	citySelect.Resize(fyne.NewSize(SelectNormalWidth, SelectNormalHeight))
	citySelect.Move(fyne.NewPos(BtnNormalWidth+NormalSpacing, MoveZero))
	top = append(top, citySelect)

	typeSelect := widget.NewSelect(models.TypeSelectList, typeSelectCallback)
	typeSelect.Selected = models.TypeSelectList[0]
	typeSelect.Resize(fyne.NewSize(SelectNormalWidth, SelectNormalHeight))
	typeSelect.Move(fyne.NewPos(BtnNormalWidth+SelectNormalWidth+2*NormalSpacing, MoveZero))
	top = append(top, typeSelect)

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("请输入关键词")
	//positionEntry.Bind(models.Keyword)
	positionEntry.OnChanged = positionEntryCallback
	positionEntry.Resize(fyne.NewSize(EntryNormalWidth, EntryNormalHeight))
	positionEntry.Move(fyne.NewPos(BtnNormalWidth+2*SelectNormalWidth+3*NormalSpacing, MoveZero))
	top = append(top, positionEntry)

	line := canvas.NewLine(Gray)
	line.Move(fyne.NewPos(MoveZero, BtnNormalHeight+NormalSpacing))
	spacer := canvas.NewRectangle(color.Transparent)
	spacer.Resize(fyne.NewSize(0, 5))

	return container.NewVBox(container.NewWithoutLayout(top...), spacer, line)
}
