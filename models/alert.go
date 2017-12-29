package models

import (
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
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Elastic        Elastic     `json:"elastic"`
	Query          string      `json:"query"`
	MatchType      string      `json:"match_type"`
	MatchFrequency int         `json:"match_frequency"`
	MatchPeriod    string      `json:"match_period"`
	Targets        AlertTarget `json:"targets"`
	LastSent       time.Time   `json:"last_sent"`
}
//type AlertStore interface {
//	AllAlerts() ([]Alert, error)
//}
