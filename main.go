package main

import (
	"action_status/action"
	"log"
	"sync"
)

func callAdd(str string, wg *sync.WaitGroup) {
	error := action.AddAction(str)
	if error != nil {
		log.Fatalln(error.Error())
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	log.Println("Processing ......")
	go callAdd("{\"action\":\"jump\", \"time\":100}", &wg)
	wg.Add(1)
	go callAdd("{\"action\":\"run\", \"time\":75}", &wg)
	wg.Add(1)
	go callAdd("{\"action\":\"jump\", \"time\":200}", &wg)
	wg.Add(1)
	wg.Wait()
	result := action.GetStats()
	log.Println("\n", result)
}
