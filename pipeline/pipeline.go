package pipeline

import "github.com/deryrahman/optimus-etl-plugin/processor"

type Pipeline interface {
	AddProcessor(processor.Processor)
	Run()
}
