package models

import (
	"fmt"
	"github.com/mickaelmagniez/elastic-alert/es"
	"github.com/olivere/elastic"
	"encoding/json"
	"time"
)

type AlertModel struct{}

type AlertEmail struct {
	Recipient string `json:"recipient"`
}
type AlertTarget struct {
	Emails []AlertEmail `json:"emails"`
}
type Elastic struct {
	Url   string `json:"url"`
	Index string `json:"index"`
	Type  string `json:"type"`
}
type Alert struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Elastic        Elastic `json:"elastic"`
	Query          string  `json:"query"`
	MatchType      string  `json:"match_type"`
	MatchFrequency int     `json:"match_frequency"`
	MatchPeriod    string  `json:"match_period"`
	//Query   string      `json:"query"`
	Targets  AlertTarget `json:"targets"`
	LastSent time.Time   `json:"last_sent"`
}

const ESIndex string = ".elastic-alert"
const ESType string = "alert"

func (m AlertModel) Create(alert Alert) (Alert, error) {
	client := es.GetES()

	put1, err := client.Index().
		Index(ESIndex).
		Type(ESType).
		BodyJson(alert).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	alert.ID = put1.Id

	return alert, nil

}
func (m AlertModel) Update(alert Alert) (Alert, error) {
	client := es.GetES()

	put1, err := client.Index().
		Index(ESIndex).
		Type(ESType).
		Id(alert.ID).
		BodyJson(alert).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	alert.ID = put1.Id

	return alert, nil

}

func (m AlertModel) Delete(id string) (string, error) {
	client := es.GetES()

	put1, err := client.Delete().
		Index(ESIndex).
		Type(ESType).
		Id(id).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}
	_, err = client.Flush().
		Index(ESIndex).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	return id, nil

}

func (m AlertModel) Get(id string) (Alert, error) {
	client := es.GetES()

	res, err := client.Search().
		Index(ESIndex).
		Type(ESType).
		Query(elastic.NewIdsQuery().Ids(id)).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Got alert %s %s\n", res.Hits.Hits[0].Id, res.Hits.Hits[0])

	var alert Alert

	err = json.Unmarshal(*res.Hits.Hits[0].Source, &alert)

	if err != nil {
		// Handle error
		panic(err)
	}

	return alert, nil

}

func (m AlertModel) All() ([]Alert, error) {
	client := es.GetES()

	searchResult, err := client.Search().
		Index(ESIndex).
		Type(ESType).
		Query(elastic.NewMatchAllQuery()).
		From(0).Size(10).
		Do(*es.GetContext())
	if err != nil {
		// Handle error
		panic(err)
	}
	alerts := make([]Alert, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {
		var alert Alert

		fmt.Println(hit.Id)
		err := json.Unmarshal(*hit.Source, &alert)
		if err != nil {
			// Deserialization failed
		}
		alert.ID = hit.Id
		alerts[i] = alert
	}
	return alerts, nil
}

func (m AlertModel) GetServers() ([]string, error) {
	client := es.GetES()

	searchResult, err := client.Search().
		Index(ESIndex).
		Type(ESType).
		Query(elastic.NewMatchAllQuery()).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.url")).
		From(0).Size(0).
		Do(*es.GetContext())

	fmt.Println("===")
	fmt.Println(searchResult.TotalHits())
	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			fmt.Println(bucket.Key.(string))
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	fmt.Println(len(servers))
	return servers, nil
}


func (m AlertModel) GetIndices(url string) ([]string, error) {
	client := es.GetES()

	searchResult, err := client.Search().
		Index(ESIndex).
		Type(ESType).
		Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery("elastic.url", url))).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.index")).
		From(0).Size(0).
		Do(*es.GetContext())

	fmt.Println("===")
	fmt.Println(url)
	fmt.Println(searchResult.TotalHits())
	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			fmt.Println(bucket.Key.(string))
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	fmt.Println(len(servers))
	return servers, nil
}



func (m AlertModel) GetTypes(url string, index string) ([]string, error) {
	client := es.GetES()

	searchResult, err := client.Search().
		Index(ESIndex).
		Type(ESType).
			//
		Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery("elastic.url", url)).Must(elastic.NewTermQuery("elastic.index", index))).
		Aggregation("elastics", elastic.NewTermsAggregation().Field("elastic.type")).
		From(0).Size(0).
		Do(*es.GetContext())

	fmt.Println("===")
	fmt.Println(index)
	fmt.Println(searchResult.TotalHits())
	if err != nil {
		// Handle error
		panic(err)
	}
	servers := make([]string, 0, searchResult.TotalHits())

	if agg, found := searchResult.Aggregations.Terms("elastics"); found {
		elastics := make(map[string]int64)
		for _, bucket := range agg.Buckets {
			fmt.Println(bucket.Key.(string))
			elastics[bucket.Key.(string)] = bucket.DocCount
			servers = append(servers, bucket.Key.(string))
		}
	}
	fmt.Println(len(servers))
	return servers, nil
}
