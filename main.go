package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/u2takey/go-utils/strings"
)

type Tail struct {
	filename  string
	line      string
	change    chan bool
	FileObjet *os.File
	time      time.Duration
	size      uint64
	posi      uint64
}

// Read last line
func (tail *Tail) ReadLine() {
	var err error
	tail.FileObjet, err = os.OpenFile(tail.filename, os.O_RDONLY, 600)
	if err != nil {
		log.Panic(err)
	}
	size2read := int64(tail.posi) - int64(tail.size)
	ret, _ := tail.FileObjet.Seek(size2read, 2)

	//log.Print(pos, err2)

	var b = make([]byte, -1*size2read)
	read, e := tail.FileObjet.Read(b)
	s := strings.BytesToString(b)
	tail.posi = tail.size

	log.Printf("bytes: %v\nerr %v\n ret: %v\n string: %v\n", read, e, ret, s)
	tail.FileObjet.Close()
}

func (tail *Tail) Check() {
	var Info os.FileInfo
	var err error
	tail.change = make(chan bool)
	//Atencion *****
	Info, err = os.Stat(tail.filename)
	if err != nil {
		log.Panic(err)
	}

	tail.size = uint64(Info.Size())
	tail.posi = tail.size

	for _loop := true; _loop == true; {
		Info, err = os.Stat(tail.filename)
		if err != nil {
			log.Panic(err)
		}
		//log.Printf("%d\n", Info.Size())
		_size := uint64(Info.Size())
		if tail.size != _size {
			tail.size = _size

			tail.change <- true
			tail.ReadLine()

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
		case <-FileWatch.change:
			log.Printf("change %v size %v", FileWatch.change, FileWatch.size)
			//FileWatch.change = false
		default:

		}
	}
}
