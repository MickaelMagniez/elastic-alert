package store

import (
	"github.com/mickaelmagniez/elastic-alert/store/datastore"
	"github.com/mickaelmagniez/elastic-alert/models"
)

type Store interface {
	// Return list of known elastic servers
	GetElasticServers() ([]string, error)

	//Ping() error

}


func AllAlerts() ([]models.Alert, error) {
	return datastore.AllAlerts()
}

func CreateAlert(alert models.Alert) (models.Alert, error) {
	return datastore.CreateAlert(alert)
}

func UpdateAlert(alert models.Alert) (models.Alert, error) {
	return datastore.UpdateAlert(alert)
}
func GetAlert(id string) (models.Alert, error) {
	return datastore.GetAlert(id)
}
func DeleteAlert(id string) (string, error) {
	return datastore.DeleteAlert(id)
}



func GetElasticServers() ([]string, error) {
	return datastore.GetServerUrls()
}

func GetElasticIndicesOfServer(server string) ([]string, error) {
	return datastore.GetIndices(server)
}
func GetElasticTypesOfIndex(server string, index string) ([]string, error) {
	return datastore.GetTypes(server, index)
}
