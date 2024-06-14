package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"

	"github.com/defiants-co/perpstream-go-2/clients/hegic"
)

func main() {
	// var Rpcs []string = []string{
	// 	"https://arbitrum.llamarpc.com",
	// 	"https://rpc.ankr.com/arbitrum",
	// 	"https://arbitrum-one-rpc.publicnode.com",
	// 	"https://arb-pokt.nodies.app",
	// 	"https://arb-mainnet-public.unifra.io",
	// 	"https://arbitrum.meowrpc.com",
	// 	"https://arbitrum-one.publicnode.com",
	// 	"https://arbitrum.rpc.subquery.network/public",
	// 	"https://1rpc.io/arb",
	// }

	// var p []models.Future
	// gc, _ := gmx.NewGmxClient(Rpcs, 10)

	// l, _ := gc.GetLeaderboard(100000)

	// var mu sync.Mutex
	// var wg sync.WaitGroup

	// for _, s := range l {
	// 	if s.PnlUsd > 15000 {
	// 		wg.Add(1)
	// 		go func(id string) {
	// 			defer wg.Done()
	// 			fmt.Println("getting " + id)
	// 			a, _ := gc.FetchPositions(id)
	// 			mu.Lock()
	// 			p = append(p, a...)
	// 			mu.Unlock()
	// 		}(s.Id)
	// 	}
	// }

	// wg.Wait()

	h, _ := hegic.NewHegicClient()

	p, _ := h.FetchPositions("0xFDE316cc38F7b9F0711EcD4E9797AACdE4FaB7d0")
	WriteStructToFile("options.csv", p)
}

func WriteStructToFile[T any](filename string, data []T) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("data is not a slice")
	}

	if val.Len() == 0 {
		return nil
	}

	firstElem := val.Index(0).Interface()
	headers := getStructFields(firstElem)
	if err := writer.Write(headers); err != nil {
		return err
	}

	for i := 0; i < val.Len(); i++ {
		row := getStructValues(val.Index(i).Interface())
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func getStructFields(data interface{}) []string {
	val := reflect.ValueOf(data)
	t := val.Type()
	fields := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fields[i] = t.Field(i).Name
	}

	return fields
}

func getStructValues(data interface{}) []string {
	val := reflect.ValueOf(data)
	values := make([]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		values[i] = fmt.Sprintf("%v", val.Field(i).Interface())
	}

	return values
}
