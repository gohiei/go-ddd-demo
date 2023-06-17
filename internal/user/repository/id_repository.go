package user

type IDRepository interface {
	Incr(int) (int64, error)
}
