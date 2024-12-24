package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newLeft() *fyne.Container {
	left := make([]fyne.CanvasObject, 0)

	bigTechBtn := widget.NewButtonWithIcon("大厂", GetDataIcon(), func() {})
	bigTechBtn.Importance = widget.WarningImportance
	left = append(left, bigTechBtn)

	bankInsuranceBtn := widget.NewButtonWithIcon("银行保险", GetDataIcon(), func() {})
	bankInsuranceBtn.Importance = widget.MediumImportance
	left = append(left, bankInsuranceBtn)

	return container.NewVBox(left...)
}
