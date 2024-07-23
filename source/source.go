package source

type Source interface {
	Process() (out <-chan []byte)
}
