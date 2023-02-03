package main

import (
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
	size      int64
	posi      int64
}

// Read last line
func (tail *Tail) ReadLine() {
	var err error
	tail.FileObjet, err = os.OpenFile(tail.filename, os.O_RDONLY, 600)
	if err != nil {
		log.Panic(err)
	}

	size2read := int64(tail.posi+1) - int64(tail.size)

	if size2read < 0 {
		tail.posi, _ = tail.FileObjet.Seek(size2read, 2)

		//log.Print(pos, err2)

		var b = make([]byte, -1*size2read)
		tail.FileObjet.Read(b)
		//tail.posi, _ = tail.FileObjet.Seek(0, 2)
		tail.posi = tail.size - 1

		tail.line = strings.BytesToString(b)
		//log.Printf("bytes: %v\nerr %v\nret: %v\nstring: %v\n", read, e, tail.posi, s)
	}

	tail.FileObjet.Close()
}

func (tail *Tail) Check() {
	var Info os.FileInfo
	var err error

	for _loop := true; _loop == true; {
		Info, err = os.Stat(tail.filename)
		if err != nil {
			//log.Fatal(err)
			log.Print("No such file..")
			tail.size = 0
			tail.posi = -1
		} else {
			//log.Printf("%d\n", Info.Size())
			_size := int64(Info.Size())
			if tail.size < _size {
				tail.size = _size
				tail.change <- true
			} else {
				if tail.size > _size {
					tail.size = _size
					tail.posi = tail.size - 1
					//tail.change <- true
				}
			}
		}

		time.Sleep(tail.time)
	}
}
