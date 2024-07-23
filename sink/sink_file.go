package sink

import (
	"io/fs"
	"log"
	"os"
)

type file struct {
	filepath string
}

func NewFile(filepath string) Sink {
	return &file{
		filepath: filepath,
	}
}

func (c *file) Process(in <-chan []byte) <-chan byte {
	termination := make(chan byte)
	go func(<-chan byte) {
		log.Default().Println("ready to write to", c.filepath)
		_ = os.WriteFile(c.filepath, <-in, fs.ModePerm)
		log.Default().Println("done writing to", c.filepath)
		termination <- byte('0')
	}(termination)

	return termination
}
