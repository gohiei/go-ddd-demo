package user

type OutsideRepository interface {
	GetEchoData() (string, error)
}
