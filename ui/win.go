package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/spf13/viper"
)

var MainApp fyne.App
var MainWindow fyne.Window

func InitUi() {
	MainApp = app.New()
	MainWindow = MainApp.NewWindow(viper.GetString("ui.title"))
	MainWindow.Resize(fyne.NewSize(float32(viper.GetInt("ui.width")), float32(viper.GetInt("ui.height"))))

	top := newTop()
	left := newLeft()
	content := newContent()

	mainLayout := container.NewBorder(top, nil, left, nil, content)
	MainWindow.SetContent(mainLayout)

	MainWindow.ShowAndRun()
}
