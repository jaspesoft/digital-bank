package adapter

import "github.com/google/uuid"

type (
	UUIDEntityID struct {
	}
)

func NewUUIDEntityID() *UUIDEntityID {
	return &UUIDEntityID{}
}

func (e UUIDEntityID) GetID() string {
	return uuid.New().String()
}
