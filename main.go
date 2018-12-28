package main

import (
	_ "weather_update_job/routers"
	"weather_update_job/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	defer toolbox.StopTask()
	cron := beego.AppConfig.String("cron")
	apiurl := beego.AppConfig.String("apiurl")
	logs.Info("cron=" + cron)
	logs.Info("apiurl=" + apiurl)
	tk1 := toolbox.NewTask("tk1", cron, func() error {
		err, res := utils.FetchPost(nil, apiurl)
		if err != nil {
			logs.Error(err)
		} else {
			logs.Info(res)
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	beego.Run()
}
