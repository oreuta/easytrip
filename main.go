package main

import (
	"os"
	"strconv"

	_ "github.com/Dimitriy14/easytrip/routers"
	"github.com/astaxie/beego"
)

func main() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	beego.Run()
}
