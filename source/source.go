package source

import (
	"github.com/satori/go.uuid"
)

type Source struct {
	UUID      uuid.UUID `json:"uuid" db:"uuid"`
	TypeUUID  uuid.UUID `json:"type_uuid" db:"type_uuid"`
	PublicURI string    `json:"public_uri" db:"public_uri"`
}
