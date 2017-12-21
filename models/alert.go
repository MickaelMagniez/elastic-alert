package models

import (
	"fmt"
	"github.com/mickaelmagniez/elastic-alert/es"
	"github.com/olivere/elastic"
	"encoding/json"
)

type AlertModel struct{}

type Alert struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
		alert.ID = hit.Id
		err := json.Unmarshal(*hit.Source, &alert)
		if err != nil {
			// Deserialization failed
		}
		alerts[i] = alert
	}
	return alerts, nil
}