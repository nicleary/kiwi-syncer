package utils

import "github.com/google/uuid"

func IsDefaultUUID(idToCheck uuid.UUID) bool {
	return idToCheck == *new(uuid.UUID)
}
