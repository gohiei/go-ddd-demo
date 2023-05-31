package dddcore

import "github.com/google/uuid"

// A wrapper of google/uuid
type UUID struct {
	uuid uuid.UUID
}

func NewUUID() UUID {
	return UUID{uuid: uuid.New()}
}

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

func (u UUID) String() string {
	if u.uuid == uuid.Nil {
		return ``
	}

	return u.uuid.String()
}

func (u UUID) IsNil() bool {
	return u.uuid == uuid.Nil
}
