package pipeline

import (
	"log"

	"github.com/deryrahman/optimus-etl-plugin/processor"
	"github.com/deryrahman/optimus-etl-plugin/sink"
	"github.com/deryrahman/optimus-etl-plugin/source"
)

type pipeline struct {
	source     source.Source
	processors []processor.Processor
	sink       sink.Sink
}

func NewPipeline(source source.Source, sink sink.Sink) Pipeline {
	return &pipeline{
		source:     source,
		processors: []processor.Processor{},
		sink:       sink,
	}
}

func (p *pipeline) AddProcessor(processor processor.Processor) {
	p.processors = append(p.processors, processor)
}

func (p *pipeline) Run() {
	log.Default().Println("start executing")
	out := p.source.Process()
	for _, processor := range p.processors {
		out = processor.Process(out)
	}
	<-p.sink.Process(out)
	log.Default().Println("done executing")
}
