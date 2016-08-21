package client

import (
	"github.com/go-kit/kit/log"
	"github.com/mosen/mdmappsvc/source"
)

func New(serviceAddr string, logger log.Logger) (source.Service, error) {
	return source.MakeClientEndpoints(serviceAddr)
}