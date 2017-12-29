package datastore

import (
	"github.com/olivere/elastic"
	"context"
	"fmt"
	"github.com/mickaelmagniez/elastic-alert/config"
)

type datastore struct {
	*elastic.Client
	ctx *context.Context
}
var EsStore *datastore

func Init() {
	configuration := config.GetConfiguration()

	ctx := context.Background()

	es, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%d", configuration.Elastic.Host, configuration.Elastic.Port)),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	EsStore = &datastore{
		Client: es,
		ctx: &ctx,
	}
}

