package main

import (
	"github.com/83Beyond/recruitment/conf"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/ui"
)

func main() {
	models.InitData()
	conf.InitConfig()
	ui.InitUi()
}
