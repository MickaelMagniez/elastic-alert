package datastore

import (
	"github.com/olivere/elastic"
	"github.com/mickaelmagniez/elastic-alert/config"
	"github.com/mickaelmagniez/elastic-alert/models"
	"encoding/json"
)

func CreateAlert(alert models.Alert) (models.Alert, error) {
	configuration := config.GetConfiguration()

	put1, err := EsStore.Index().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		BodyJson(alert).
		Do(*EsStore.ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	alert.ID = put1.Id

	return alert, nil

}
func UpdateAlert(alert models.Alert) (models.Alert, error) {
	configuration := config.GetConfiguration()

	put1, err := EsStore.Index().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Id(alert.ID).
		BodyJson(alert).
		Do(*EsStore.ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	alert.ID = put1.Id

	return alert, nil

}

func DeleteAlert(id string) (string, error) {
	configuration := config.GetConfiguration()

	_, err := EsStore.Delete().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Id(id).
		Do(*EsStore.ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	//_, err = EsStore.Flush().
	//	Index(configuration.Elastic.Index).
	//	Do(*EsStore.ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}

	return id, nil

}

func GetAlert(id string) (models.Alert, error) {
	configuration := config.GetConfiguration()

	searchResult, err := EsStore.Search().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Query(elastic.NewIdsQuery().Ids(id)).
		Do(*EsStore.ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	var alert models.Alert

	err = json.Unmarshal(*searchResult.Hits.Hits[0].Source, &alert)

	if err != nil {
		// Handle error
		panic(err)
	}

	return alert, nil

}

func AllAlerts() ([]models.Alert, error) {
	configuration := config.GetConfiguration()

	searchResult, err := EsStore.Search().
		Index(configuration.Elastic.Index).
		Type(configuration.Elastic.Type).
		Query(elastic.NewMatchAllQuery()).
		From(0).Size(10).
		Do(*EsStore.ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	alerts := make([]models.Alert, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {
		var alert models.Alert

		err := json.Unmarshal(*hit.Source, &alert)
		if err != nil {
			// Deserialization failed
		}
		alert.ID = hit.Id
		alerts[i] = alert
	}
	return alerts, nil
}
