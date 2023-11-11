package main

import (
	"flag"
	"fmt"
	"one-day-offer/my-office-25-11-2023/internal/collector"
	"one-day-offer/my-office-25-11-2023/internal/readers"
	"time"
)

const (
	maxChanSize    = 32
	requestTimeOut = time.Second * 10
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("error there should be only 1 argument")
	}
	fileName := flag.Args()[0]

	buf, err := readers.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	var (
		results = make(chan string, maxChanSize)
		urls    = make(chan string, maxChanSize)
	)
	go func() {
		if err = readers.ReadUrls(buf, urls); err != nil {
			panic(err)
		}
		close(urls)
	}()

	col := collector.NewCollector(requestTimeOut)
	go func() {
		col.FetchData(urls, results)
		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}
}
