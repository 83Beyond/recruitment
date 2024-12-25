package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newLeft() *fyne.Container {
	left := make([]fyne.CanvasObject, 0)

	bigTechBtn := widget.NewButtonWithIcon("大厂", getBigTechIcon(), func() {})
	bigTechBtn.Importance = widget.WarningImportance
	left = append(left, bigTechBtn)

	bankBtn := widget.NewButtonWithIcon("银行", getFinanceIcon(), func() {})
	bankBtn.Importance = widget.MediumImportance
	left = append(left, bankBtn)

	insuranceBtn := widget.NewButtonWithIcon("保险", getFinanceIcon(), func() {})
	insuranceBtn.Importance = widget.MediumImportance
	left = append(left, insuranceBtn)

	return container.NewVBox(left...)
}
