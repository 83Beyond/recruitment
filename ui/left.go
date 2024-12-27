package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LeftButtonGroup struct {
	bigTechBtn   *widget.Button
	bankBtn      *widget.Button
	insuranceBtn *widget.Button
}

func NewLeftButtonGroup() *LeftButtonGroup {
	bg := &LeftButtonGroup{}

	bg.bigTechBtn = widget.NewButtonWithIcon("大厂", getBigTechIcon(), bg.bigTechBtnCallback)
	bg.bigTechBtn.Importance = widget.HighImportance

	bg.bankBtn = widget.NewButtonWithIcon("银行", getBankIcon(), bg.bankBtnCallback)
	bg.bankBtn.Importance = widget.MediumImportance

	bg.insuranceBtn = widget.NewButtonWithIcon("保险", getInsuranceIcon(), bg.insuranceCallback)
	bg.insuranceBtn.Importance = widget.MediumImportance

	return bg
}

func (bg *LeftButtonGroup) bigTechBtnCallback() {
	bg.bigTechBtn.Importance = widget.HighImportance
	bg.bankBtn.Importance = widget.MediumImportance
	bg.insuranceBtn.Importance = widget.MediumImportance

	bg.refreshBtn()
}

func (bg *LeftButtonGroup) bankBtnCallback() {
	bg.bigTechBtn.Importance = widget.MediumImportance
	bg.bankBtn.Importance = widget.HighImportance
	bg.insuranceBtn.Importance = widget.MediumImportance

	bg.refreshBtn()
}

func (bg *LeftButtonGroup) insuranceCallback() {
	bg.bigTechBtn.Importance = widget.MediumImportance
	bg.bankBtn.Importance = widget.MediumImportance
	bg.insuranceBtn.Importance = widget.HighImportance

	bg.refreshBtn()
}

func (bg *LeftButtonGroup) refreshBtn() {
	bg.bigTechBtn.Refresh()
	bg.bankBtn.Refresh()
	bg.insuranceBtn.Refresh()
}

func newLeft() *fyne.Container {
	buttonGroup := NewLeftButtonGroup()
	left := make([]fyne.CanvasObject, 0)
	left = append(left, buttonGroup.bigTechBtn)
	left = append(left, buttonGroup.bankBtn)
	left = append(left, buttonGroup.insuranceBtn)

	return container.NewVBox(left...)
}
