package main

import (
	"context"
	"log"
	"time"
)

func main() {
	FileWatch := Tail{filename: "file.log", time: time.Second * 1, size: 0}
	FileWatch.change = make(chan bool)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//go tail("file.log", cancel, time.Millisecond*100)
	go FileWatch.Check()

	//log.Print(<-line)
	for _loop := true; _loop == true; {
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			log.Print("Done")
		case <-FileWatch.change:
			FileWatch.ReadLine()
			log.Printf("%s", FileWatch.line)
			//FileWatch.change = false
		default:

		}
	}
}
