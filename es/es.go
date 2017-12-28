package es

import (
	"github.com/olivere/elastic"
	"context"
	"github.com/mickaelmagniez/elastic-alert/config"
	"fmt"
)

type ES struct {
	*elastic.Client
}

var es *elastic.Client
var ctx *context.Context

func Init() {
	configuration := config.GetConfiguration()

	ctx1 := context.Background()
	ctx = &ctx1

	es1, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%d", configuration.Elastic.Host, configuration.Elastic.Port)),
		elastic.SetSniff(false),

	)
	es = es1
	if err != nil {
		panic(err)
	}

}

func GetES() *elastic.Client {
	return es
}

func GetContext() *context.Context {
	return ctx
}
