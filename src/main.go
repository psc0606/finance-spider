package main

import (
	"finance-spider/src/model"
	"finance-spider/src/service"
	"github.com/b3log/gulu"
)

// init function when package loaded.
func init() {
	gulu.Log.SetLevel("warn")
	model.LoadConf()
}

func main() {
	service.ConnectDB()
	_ = service.FundDownload()
	defer service.DisconnectDB()
}
