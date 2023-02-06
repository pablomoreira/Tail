package tail

import (
	"log"
	"os"
	"time"

	"github.com/u2takey/go-utils/strings"
)

type Tail struct {
	Filename  string
	Buffer    string
	Change    chan bool
	FileObjet *os.File
	Time      time.Duration
	Size      int64
	posi      int64
}

// Read last line
func (tail *Tail) Read() {
	var err error
	tail.FileObjet, err = os.OpenFile(tail.Filename, os.O_RDONLY, 600)
	if err != nil {
		log.Panic(err)
	}

	size2read := int64(tail.posi+1) - int64(tail.Size)

	if size2read < 0 {
		tail.posi, _ = tail.FileObjet.Seek(size2read, 2)

		//log.Print(pos, err2)

		var b = make([]byte, -1*size2read)
		tail.FileObjet.Read(b)
		//tail.posi, _ = tail.FileObjet.Seek(0, 2)
		tail.posi = tail.Size - 1

		tail.Buffer = strings.BytesToString(b)
	}

	tail.FileObjet.Close()
}

func (tail *Tail) Check() {
	var Info os.FileInfo
	var err error

	for _loop := true; _loop == true; {
		Info, err = os.Stat(tail.Filename)
		if err != nil {
			//log.Fatal(err)
			log.Print("No such file..")
			tail.Size = 0
			tail.posi = -1
		} else {
			//log.Printf("%d\n", Info.Size())
			_size := int64(Info.Size())
			if tail.Size < _size {
				tail.Size = _size
				tail.Change <- true
			} else {
				if tail.Size > _size {
					tail.Size = _size
					tail.posi = tail.Size - 1
					//tail.change <- true
				}
			}
		}

		time.Sleep(tail.Time)
	}
}
