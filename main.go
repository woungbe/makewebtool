package main

import (
	"fmt"
	"makewebtool/configs"
	"makewebtool/router/httpd"
	"os"
)

func serviceStart() {
	cnf := configs.GetConfig()
	er := cnf.InitConfig("config.json")
	if er != nil {
		fmt.Println("설정정보로드에러")
		os.Exit(1)
	}

	httpd.SetRouters(cnf.E) // http 등록하기 !!
	cnf.E.Start(":8000")

}

func main() {
	serviceStart()
	// cnf := model.Init()
	// cnf.MakeFileFor()
	// fmt.Println(cnf)
}
