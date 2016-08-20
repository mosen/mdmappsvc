package source

import (
	"github.com/satori/go.uuid"
)

type Source struct {
	UUID uuid.UUID `json:"uuid" db:"uuid"`
	typeUUID uuid.UUID `json:"type_uuid" db:"type_uuid"`

}
