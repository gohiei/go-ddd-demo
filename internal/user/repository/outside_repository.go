package repository

type OutsideRepository interface {
	GetEchoData() (string, error)
}
