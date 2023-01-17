package main

import (
	"context"
	"log"
	"os"
	"time"
)

type Tail struct {
	filename  string
	line      string
	change    bool
	FileObjet os.File
	time      time.Duration
	size      uint64
	test      chan bool
}

func (tail *Tail) Check() {
	var Info os.FileInfo
	var err error
	tail.test = make(chan bool)
	for _loop := true; _loop == true; {
		Info, err = os.Stat(tail.filename)
		if err != nil {
			log.Panic(err)
		}
		//log.Printf("%d\n", Info.Size())
		_size := uint64(Info.Size())
		if tail.size != _size {
			tail.size = _size
			tail.change = true
			tail.test <- true

		} else {
		}
		time.Sleep(tail.time)

	}
}

func main() {
	FileWatch := Tail{filename: "file.log", time: time.Second * 1, size: 0}

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
		case <-FileWatch.test:
			log.Printf("change %v size %v", FileWatch.change, FileWatch.size)
			FileWatch.change = false
		default:

		}

	}

}

//func check
