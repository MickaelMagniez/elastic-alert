package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"time"
	"encoding/json"
	"github.com/mickaelmagniez/elastic-alert/config"
	"context"
	"github.com/mickaelmagniez/elastic-alert/store/datastore"
	"github.com/mickaelmagniez/elastic-alert/store"
)

func main() {

	config.InitConfiguration()

	datastore.Init()

	configuration := config.GetConfiguration()

	fmt.Println("worker")
	alerts, err := store.AllAlerts()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	} else {

		fmt.Println(len(alerts))

		for _, alert := range alerts {

			ctx := context.Background()

			client, _ := elastic.NewClient(
				elastic.SetURL(alert.Elastic.Url),
				elastic.SetSniff(false),

			)

			//client := es.GetES()

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
				Index(alert.Elastic.Index).
				Type(alert.Elastic.Index).
			//Query(elastic.RawStringQuery(alert.Query)).
				Query(query).
				Do(ctx)
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
				store.UpdateAlert(alert)
				fmt.Println("match frequency ok")
				for _, email := range alert.Targets.Emails {
					fmt.Printf("email %s\n", email.Recipient)
					fmt.Printf("sender %s\n", configuration.Targets.Email.Sender)

					m := gomail.NewMessage()

					m.SetHeader("From", configuration.Targets.Email.Sender)
					m.SetHeader("To", email.Recipient)
					//m.SetHeader("From", "alex@example.com")
					//m.SetHeader("To", "bob@example.com", "cora@example.com")
					m.SetHeader("Subject", fmt.Sprintf("Alert '%s' triggered!", alert.Name))
					//fmt.Println(string((*(res.Hits.Hits[0].Source))[:]))
					a, _ := json.MarshalIndent(res.Hits.Hits[0].Source, "", "\t")
					fmt.Println(string(a[:]))
					m.SetBody("text/html", fmt.Sprintf("<pre>%s</pre>", string(a[:])))

					d := gomail.NewDialer(configuration.Targets.Email.Smtp.Host, configuration.Targets.Email.Smtp.Port, configuration.Targets.Email.Smtp.Username, configuration.Targets.Email.Smtp.Password)
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
