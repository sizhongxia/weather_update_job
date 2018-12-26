package main

import (
	_ "update_weather_job/routers"
	"update_weather_job/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	defer toolbox.StopTask()
	tk1 := toolbox.NewTask("tk1", beego.AppConfig.String("cron"), func() error {
		err, _ := utils.FetchPost(nil, beego.AppConfig.String("apiurl"))
		if err != nil {
			logs.Error(err)
		}
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	beego.Run()
}
