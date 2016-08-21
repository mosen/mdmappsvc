package source

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

type Service interface {
	PostSource(ctx context.Context, s *Source) error
	GetSource(ctx context.Context, uuidStr string) (Source, error)
	PutSource(ctx context.Context, uuidStr string, s *Source) error
	PatchSource(ctx context.Context, uuidStr string, s *Source) error
	DeleteSource(ctx context.Context, uuidStr string) error
	GetSources(ctx context.Context) ([]Source, error)
}

var (
	ErrMalformedUUID = errors.New("malformed UUID")
	ErrNotFound      = errors.New("device not found")
	ErrQueryError    = errors.New("error performing query")
)

type service struct {
	store  SourceRepository
	logger log.Logger
}

func NewService(sr SourceRepository, logger log.Logger) Service {
	return &service{
		sr,
		logger,
	}
}

func (svc *service) PostSource(ctx context.Context, s *Source) error {
	if err := svc.store.Store(s); err != nil {
		return err
	}

	return nil
}

func (svc *service) GetSource(ctx context.Context, uuidStr string) (Source, error) {
	uuidObj, err := uuid.FromString(uuidStr)
	if err != nil {
		return Source{}, ErrMalformedUUID
	}

	source, err := svc.store.Find(uuidObj)
	if err != nil {
		return Source{}, ErrQueryError
	}

	if source == nil {
		return Source{}, ErrNotFound
	} else {
		return *source, nil
	}
}

func (svc *service) PutSource(ctx context.Context, uuidStr string, s *Source) error {
	return errors.New("Not Implemented")
}

func (svc *service) PatchSource(ctx context.Context, uuidStr string, s *Source) error {
	return errors.New("Not Implemented")
}

func (svc *service) DeleteSource(ctx context.Context, uuidStr string) error {
	uuidObj, err := uuid.FromString(uuidStr)
	if err != nil {
		return ErrMalformedUUID
	}

	if err := svc.store.Delete(uuidObj); err != nil {
		return err
	} else {
		return nil
	}
}

func (svc *service) GetSources(ctx context.Context) ([]Source, error) {
	return svc.store.FindAll()
}
