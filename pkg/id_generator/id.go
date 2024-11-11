package id_generator

import "github.com/google/uuid"

type UUID = uuid.UUID

func NewUUID() (uuid.UUID, error) {
	uuid, err := uuid.NewV7()
	return uuid, err
}
