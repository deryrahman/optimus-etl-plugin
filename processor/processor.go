package processor

type Processor interface {
	Process(in <-chan []byte) (out <-chan []byte)
}
