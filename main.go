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
}

func (tail *Tail) Check() {
	var Info os.FileInfo
	var err error
	for _loop := true; _loop == true; {
		Info, err = os.Stat(tail.filename)

		if err != nil {
			log.Panic(err)
		}
		log.Printf("%d\n", Info.Size())
		time.Sleep(tail.time)

	}
}

func main() {
	FileWatch := Tail{filename: "file.log", time: time.Second * 1}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//go tail("file.log", cancel, time.Millisecond*100)
	go FileWatch.Check()

	//log.Print(<-line)
	for _loop := true; _loop == true; {
		select {
		case <-ctx.Done():
			log.Print("Done")
			//_loop = false
		default:
			log.Print(FileWatch.change)
		}
		time.Sleep(time.Millisecond * 1000)
	}

}

//func check
