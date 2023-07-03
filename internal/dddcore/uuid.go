package dddcore

import "github.com/google/uuid"

// UUID represents a universally unique identifier.
type UUID struct {
	uuid uuid.UUID
}

// NewUUID generates a new UUID.
func NewUUID() UUID {
	return UUID{uuid: uuid.New()}
}

// BuildUUID constructs a UUID from the given string.
// It returns an error if the string is not a valid UUID.
func BuildUUID(id string) (UUID, error) {
	if id == "" {
		return UUID{uuid: uuid.Nil}, nil
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return UUID{}, err
	}

	return UUID{uuid: uid}, nil
}

// String returns the string representation of the UUID.
// If the UUID is nil, an empty string is returned.
func (u UUID) String() string {
	if u.uuid == uuid.Nil {
		return ``
	}

	return u.uuid.String()
}

// IsNil checks if the UUID is nil.
// A UUID is considered nil if its underlying value is uuid.Nil.
func (u UUID) IsNil() bool {
	return u.uuid == uuid.Nil
}
