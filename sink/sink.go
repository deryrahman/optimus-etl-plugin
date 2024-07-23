package sink

type Sink interface {
	Process(in <-chan []byte) (termination <-chan byte)
}
