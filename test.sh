curl -X DELETE "http://127.0.0.1:9200/test"
curl -X PUT -H "Content-Type: application/json" "http://127.0.0.1:9200/test" -d @test.json
DATE=`date +%Y-%m-%dT%H:%M:%S%z`

curl -X POST -H "Content-Type: application/json" "http://127.0.0.1:9200/test/test" -d "{\"date\":\"`date +%Y-%m-%dT%H:%M:%S%z`\", \"message\": \"error\"}"