package main

import (
	"github.com/deryrahman/optimus-etl-plugin/pipeline"
	"github.com/deryrahman/optimus-etl-plugin/sink"
	"github.com/deryrahman/optimus-etl-plugin/source"
)

func main() {
	// source
	r := source.NewFile("./example.json")
	// processor
	// s := processor.NewSelector()
	// sink
	w := sink.NewFile("./output.json")

	p := pipeline.NewPipeline(r, w)
	// p.AddProcessor(s)

	p.Run()
	// <-c
}
