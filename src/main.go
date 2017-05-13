package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
)

var (
	workers int = 2
)

func main() {
	fmt.Println("Hello World")
	
	flag.Parse()
	
	grab()
}

func init() {
	flag.IntVar(&workers, "w", workers, "количество потоков")
}

func grab() {
	
	response, err := http.Get("https://market.yandex.ru/product/1712214246?nid=56179&track=main_page_endless_snippet");
	
	if err == nil {
		
		defer response.Body.Close()
		
		body, _ := ioutil.ReadAll(response.Body)
		
		fmt.Println(body)
	}
	// test
	
}
