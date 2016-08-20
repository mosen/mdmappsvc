package source

import (
	"github.com/satori/go.uuid"
	"github.com/jmoiron/sqlx"
	sq "github.com/Masterminds/squirrel"
	kitlog "github.com/go-kit/kit/log"
)

type SourceRepository interface {
	Find(uuid uuid.UUID) (*Source, error)
	FindAll() ([]Source, error)
	Store(source *Source) error
	Delete(uuid uuid.UUID) error
	Update(source *Source) error
}

type sourceRepository struct {
	*sqlx.DB
	kitlog.Logger
}

func (r *sourceRepository) Find(uuid uuid.UUID) (*Source, error) {

}

func (r *sourceRepository) FindAll() ([]Source, error) {

}

func (r *sourceRepository) Store(source *Source) error {

}

func (r *sourceRepository) Delete(uuid uuid.UUID) error {

}

func (r *sourceRepository) Update(source *Source) error {

}
