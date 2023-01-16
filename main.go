package main

import (
	"context"
	"log"
	"os"
	"time"
)
type person struct {
    filename string
	line string
    chang bool
	FileObjet os.File
}

func main() {
	//line := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	
	go tail("file.log", cancel, time.Millisecond*100)

	//log.Print(<-line)
	for _loop := true; _loop == true; {
		select {
		case <-ctx.Done():
			log.Print("Done")
			_loop = false
		default:
			
		}
		time.Sleep(time.Millisecond * 1)
	}
}

func tail(filename string, cancel context.CancelFunc, t_sleep time.Duration) {
	//File, err := os.OpenFile(filename, os.O_RDONLY, 644)
	var Info os.FileInfo
	var err error
	time.Sleep(t_sleep)
	for _loop := true; _loop == true; {
		Info, err = os.Stat(filename)

		if err != nil {
			log.Panic(err)
		}
		log.Printf("%d\n", Info.Size())
		time.Sleep(time.Second * 1)
		_loop = false
		cancel()
	}
}

func check