package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	line := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go tail("file.log", line, ctx)

	log.Print(<-line)
}

func tail(filename string, line chan string, ctx context.Context) {
	//File, err := os.OpenFile(filename, os.O_RDONLY, 644)
	var Info os.FileInfo
	var err error

	for true {
		Info, err = os.Stat(filename)

		if err != nil {
			log.Panic()
		}
		log.Printf("%d\n", Info.Size())
		time.Sleep(time.Second * 2)
	}
	line <- "End"
	ctx.Done()
}
