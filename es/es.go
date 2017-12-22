package es

import (
	"github.com/olivere/elastic"
	"context"
)

//DB ...
type ES struct {
	*elastic.Client
}

const (
	Host = "postgres"
)

var es *elastic.Client
var ctx *context.Context

//Init ...
func Init() {
	ctx1 := context.Background()
	ctx = &ctx1

	es1, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),

	)
	es = es1
	if err != nil {
		// Handle error
		panic(err)
	}

}

////ConnectDB ...
//func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
//	db, err := sql.Open("postgres", dataSourceName)
//	if err != nil {
//		return nil, err
//	}
//	if err = db.Ping(); err != nil {
//		return nil, err
//	}
//	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
//	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
//	return dbmap, nil
//}

func GetES() *elastic.Client {
	return es
}

func GetContext() *context.Context {
	return ctx
}
