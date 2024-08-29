package cookies

import (
	"log"

	"github.com/gofrs/uuid"
)

func newUUID() uuid.UUID{
	uudi, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return uudi
}