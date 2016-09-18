package source

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
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
	log.Logger
}

func (r *sourceRepository) Find(uuid uuid.UUID) (*Source, error) {
	device := sq.Select("*").From("sources").Where(sq.Eq{"uuid": uuid.String()})
	sql, args, err := device.ToSql()
	if err != nil {
		return nil, err
	}

	var result Source
	if err := r.Get(&result, sql, args...); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *sourceRepository) FindAll() ([]Source, error) {
	stmt := sq.Select("*").From("sources")
	sql, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	var result []Source
	if err := r.Select(&result, sql, args...); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *sourceRepository) Store(source *Source) error {
	query, args, err := sq.Insert("sources").
		Columns(
			"type_uuid",
		).
		Values(
			source.TypeUUID.String(),
		).
		Suffix("RETURNING \"uuid\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	if err := r.QueryRow(query, args...).Scan(&source.UUID); err != nil {
		return err
	}

	return nil
}

func (r *sourceRepository) Delete(uuid uuid.UUID) error {
	stmt := sq.Delete("sources").Where(sq.Eq{"uuid": uuid.String()})
	sql, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.Exec(sql, args...); err != nil {
		return err
	}

	return nil
}

func (r *sourceRepository) Update(source *Source) error {
	stmt := sq.Update("sources").SetMap(
		sq.Eq{
			"type_uuid": source.TypeUUID.String(),
		},
	).Where(sq.Eq{"uuid": source.UUID.String()})

	sql, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.Exec(sql, args...); err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sqlx.DB, logger log.Logger) SourceRepository {
	return &sourceRepository{
		db,
		logger,
	}
}
