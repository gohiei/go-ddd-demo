package dddcore

type Input interface {
}

type Output interface {
}

type UseCase[I Input, O Output] interface {
	Execute(*I) (O, error)
}
