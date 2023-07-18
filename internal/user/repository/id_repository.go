package repository

// IDRepository represents a repository for generating and managing unique identifiers.
type IDRepository interface {
	// Incr increments the identifier by the specified value and returns the updated value.
	Incr(int) (int64, error)
}
