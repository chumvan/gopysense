package main

import (
	"context"
	"fmt"
	"math"
	"reflect"
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

	done := make(chan bool)

	tickerHighFreq := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-done:
				client.Close()
				return
			case t := <-tickerHighFreq.C:
				m, err := sensehat.GetOrientation()
				if err != nil {
					panic(err)
				}
				fmt.Println(m.String())
				value := m.Pitch
				p := influxdb2.NewPoint("oriHighRes",
					map[string]string{"unit": "Degree"},
					map[string]interface{}{"avg": value},
					time.Now())
				writeAPI.WritePoint(context.Background(), p)
				fmt.Printf("Packet sent at: %s\n", t)
				fmt.Printf("High frequency value: %f of type: %s\n", value, reflect.TypeOf(value))
				fmt.Println("===============")
			}
		}
	}()

	tickerLowFreq := time.NewTicker(3000 * time.Millisecond)
	go func() {
		for {
			select {
			case <-done:
				client.Close()
				return
			case t := <-tickerLowFreq.C:
				m, err := sensehat.GetOrientation()
				if err != nil {
					panic(err)
				}
				fmt.Println(m.String())
				value := float32(math.Round(m.Pitch))
				p := influxdb2.NewPoint("oriLowRes",
					map[string]string{"unit": "Degree"},
					map[string]interface{}{"avg": value},
					time.Now())
				err = writeAPI.WritePoint(context.Background(), p)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Packet sent at: %s\n", t)
				fmt.Printf("Low frequency value: %f of type: %s\n", value, reflect.TypeOf(value))
				fmt.Println("===============")

			}
		}
	}()

	time.Sleep(30000 * time.Millisecond)
	tickerHighFreq.Stop()
	tickerLowFreq.Stop()
	fmt.Println("Finished all sending")
	done <- true
}
