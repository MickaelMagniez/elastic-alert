package main

import (
	"fmt"
	"github.com/mickaelmagniez/elastic-alert/models"
	"github.com/mickaelmagniez/elastic-alert/es"
	"github.com/olivere/elastic"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"time"
	"encoding/json"
)

var alertModel = new(models.AlertModel)

func main() {
	es.Init()

	fmt.Println("worker")
	alerts, err := alertModel.All()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	} else {
		client := es.GetES()
		fmt.Println(len(alerts))

		for _, alert := range alerts {
			fmt.Println(alert)
			query := elastic.NewBoolQuery()
			query = query.Must(
				//elastic.NewRangeQuery("date").From(alert.LastSent).To("now"),
				elastic.RawStringQuery(alert.Query),
			)
			src, err := query.Source()
			if err != nil {
				panic(err)
			}
			data, err := json.Marshal(src)
			if err != nil {
				panic(err)
			}
			s := string(data)

			fmt.Println(s)
			//test  := map[string]interface{}{}
			//json.Unmarshal([]byte(alert.Query), &test)
			//
			//query :=  map[string]interface{}{}
			//json.Unmarshal([]byte("{\"bool\": {\"must\": [{\"range\": {\"date\": {\"gte\": \"2015-01-01 00:00:00\", \"lte\": \"now\"}}}]}}"), &query)
			//
			//query["bool"]["must"] = append(query["bool"]["must"], test)
			res, err := client.Search().
			//Index(ESIndex).
			//Type(ESType).
			//Query(elastic.RawStringQuery(alert.Query)).
				Query(query).
				Do(*es.GetContext())
			if err != nil {
				// Handle error
				panic(err)
			}

			fmt.Printf("Got alert %d\n", res.TotalHits())

			var limit = 1
			if alert.MatchType != "once" {
				limit = alert.MatchFrequency
			}
			if int(res.TotalHits()) >= limit {
				alert.LastSent = time.Now()
				alertModel.Update(alert)
				fmt.Println("match frequency ok")
				for _, email := range alert.Targets.Emails {
					fmt.Printf("email %s\n", email.Recipient)

					m := gomail.NewMessage()

					m.SetHeader("From", "alert@example.com")
					m.SetHeader("To", email.Recipient)
					m.SetHeader("Subject", fmt.Sprintf("Alert '%s' triggered!", alert.Name))
					//fmt.Println(string((*(res.Hits.Hits[0].Source))[:]))
					a, _ := json.MarshalIndent(res.Hits.Hits[0].Source, "", "\t")
					fmt.Println(string(a[:]))
					m.SetBody("text/html", fmt.Sprintf("<pre>%s</pre>",string(a[:])))

					d := gomail.NewDialer("127.0.0.1", 1025, "", "")
					d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
					// Send the email to Bob, Cora and Dan.
					if err := d.DialAndSend(m); err != nil {
						panic(err)
					}

				}

			}

		}

	}
}
