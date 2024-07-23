package source

import (
	"log"
	"os"
)

type file struct {
	filepath string
}

func NewFile(filepath string) Source {
	return &file{
		filepath: filepath,
	}
}

func (c *file) Process() <-chan []byte {
	out := make(chan []byte)
	go func(chan []byte) {
		log.Default().Println("ready to read from", c.filepath)
		b, _ := os.ReadFile(c.filepath)
		log.Default().Println("done reading from", c.filepath)
		out <- b
		close(out)
	}(out)
	return out
}
