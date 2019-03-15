package main

import (
	_ "yeetong_zhny_platform_job/routers"
	"yeetong_zhny_platform_job/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	defer toolbox.StopTask()

	weathercron := beego.AppConfig.String("weathercron")
	weatherupdateapiurl := beego.AppConfig.String("weatherupdateapiurl")
	logs.Info("weather cron=" + weathercron)
	logs.Info("weather update api url=" + weatherupdateapiurl)

	weatherupdatetask := toolbox.NewTask("weatherupdatetask", weathercron, func() error {
		err, _ := utils.FetchPost(nil, weatherupdateapiurl)
		if err != nil {
			logs.Error(err)
		}
		return nil
	})
	toolbox.AddTask("weatherupdatetask", weatherupdatetask)

	toolbox.StartTask()
	beego.Run()
}
