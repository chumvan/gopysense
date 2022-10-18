package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chumvan/gopysense/pkg/sensehat"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverURL := viper.Get("SERVER_URL").(string)
	authToken := viper.Get("WRITE_API_TOKEN").(string)
	org := viper.Get("ORG").(string)
	bucket := viper.Get("BUCKET").(string)

	fmt.Printf(`
	Url => %s,\n
	Token => %s,\n
	Org => %s,\n
	Bucket => %s,\n
	`, serverURL, authToken, org, bucket)

	client := influxdb2.NewClient(string(serverURL), string(authToken))
	fmt.Println("Successfully connected")
	writeAPI := client.WriteAPIBlocking(org, bucket)

	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				client.Close()
				return
			case t := <-ticker.C:
				m, err := sensehat.GetAllEnv()
				if err != nil {
					panic(err)
				}
				fmt.Println(m.String())

				p := influxdb2.NewPoint("sensory",
					map[string]string{"unit": "Celsius"},
					map[string]interface{}{"avg": m.Temperature},
					time.Now())
				writeAPI.WritePoint(context.Background(), p)
				fmt.Printf("Packet sent at: %s\n", t)

			}
		}
	}()
	time.Sleep(20000 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Finished all sending")
	done <- true
}
