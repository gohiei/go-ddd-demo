package dddcore

type Input interface {
}

type Output interface {
}

type UseCase interface {
	Execute(input *Input) (Output, error)
}
