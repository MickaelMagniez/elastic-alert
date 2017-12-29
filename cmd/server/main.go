package main

import (
	"github.com/mickaelmagniez/elastic-alert/store/datastore"
	"github.com/mickaelmagniez/elastic-alert/config"
	"github.com/mickaelmagniez/elastic-alert/api"
)


func main() {
	//SetupES()

	config.InitConfiguration()

	datastore.Init()

	api.Run()


}
