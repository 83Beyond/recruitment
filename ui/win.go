package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/utils"
	"github.com/spf13/viper"
	"math/rand"
	"time"
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

	go func() {
		fmt.Println("更新UI的协程启动")
		rand.Seed(time.Now().UnixNano())
		for position := range models.PositionInfoChan {

			first, second := utils.ProcessCity(position.City)

			if !models.AllCitySet.Contains(first + "-" + second) {
				models.AllCity = append(models.AllCity, first+"-"+second)
				models.AllCitySet.Add(first + "-" + second)
			}
			position.City = first + "-" + second
			models.PositionList = append(models.PositionList, position)
			models.ShowPositionList = append(models.ShowPositionList, position)

			if !models.AllSourceSet.Contains(position.Source) {
				models.AllSource = append(models.AllSource, position.Source)
				models.AllSourceSet.Add(position.Source)
			}
			SourceHeader.Options = models.AllSource

			if !models.FirstCitySet.Contains(first) {
				models.FirstCity = append(models.FirstCity, first)
				models.CityMap[first] = []string{"全部"}
				models.CityMap[first] = append(models.CityMap[first], second)
				models.FirstCitySet.Add(first)
				models.SecondCitySet.Add(second)
			} else {
				if !models.SecondCitySet.Contains(second) {
					models.CityMap[first] = append(models.CityMap[first], second)
					models.SecondCitySet.Add(second)
				}
			}
			FirstCityHeader.Options = models.FirstCity

			//total := len(models.PositionList)
			interval := time.Duration(rand.Intn(10) + 40)
			time.Sleep(interval * time.Millisecond)
			ShowList.Refresh()
			//models.TotalCount.Set(total)
			ShowList.ScrollToBottom()
		}
	}()

	MainWindow.ShowAndRun()
}
