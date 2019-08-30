package main

import (
	"encoding/json"
	"fmt"

	client "github.com/influxdata/influxdb/client/v2"
)

func main() {
	httpClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println(err)
	}
	q := client.NewQuery("select mean(v1), mean(v2) from cpu", "test", "s")
	if resp, err := httpClient.Query(q); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("resp.Results:", resp.Results)
		fmt.Println("resp.Results[0]", resp.Results[0])
		fmt.Println(resp.Results[0].Series[0].Values[0])
		if s, ok := (resp.Results[0].Series[0].Values[0][0]).(json.Number); ok {
			fmt.Println(s.String())
		} else {
			fmt.Println("error")
		}
	}
}
