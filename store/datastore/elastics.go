package datastore

import (
	//"github.com/mickaelmagniez/elastic-alert/es"
	"github.com/olivere/elastic"
	"github.com/mickaelmagniez/elastic-alert/config"
)

func GetServerUrls() ([]string, error) {
	configuration := config.GetConfiguration()

	searchResult, err := EsStore.Search().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Query(elastic.NewMatchAllQuery()).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.url")).
		From(0).Size(0).
		Do(*EsStore.ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	return servers, nil
}

func GetIndices(url string) ([]string, error) {
	configuration := config.GetConfiguration()

	searchResult, err := EsStore.Search().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery("elastic.url", url))).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.index")).
		From(0).Size(0).
		Do(*EsStore.ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	return servers, nil
}

func GetTypes(url string, index string) ([]string, error) {

	configuration := config.GetConfiguration()

	searchResult, err := EsStore.Search().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
	//
		Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery("elastic.url", url)).Must(elastic.NewTermQuery("elastic.index", index))).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.type")).
		From(0).Size(0).
		Do(*EsStore.ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	return servers, nil
}
