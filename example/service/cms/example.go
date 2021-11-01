package main

import (
	"time"

	tachyoncms "github.com/quantum-box/tachyon-sdk-go/service/cms"
)

// domain layer
// entity
type ContentEntity struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	data      map[string]interface{}
}

func New(id string, createdAt, updatedAt time.Time, data map[string]interface{}) *ContentEntity {
	return &ContentEntity{id, createdAt, updatedAt, data}
}

// repository interface
type ContentRepository interface {
	GetById(id string) (*ContentEntity, error)
	FindAll() ([]*ContentEntity, error)
}

// interface adapter layer
// repository implementation
var _ ContentRepository = &ContentRepositoryImpl{}

type ContentRepositoryImpl struct {
	cms TachyonCmsDriver
}

type TachyonCmsDriver interface {
	GetById(id string) (*tachyoncms.AggregateDto, error)
	FindAll() ([]*tachyoncms.AggregateDto, error)
}

func (r *ContentRepositoryImpl) GetById(id string) (*ContentEntity, error) {
	dto, err := r.cms.GetById(id)
	if err != nil {
		return nil, err
	}
	return r.into(dto), nil

}

func (*ContentRepositoryImpl) FindAll() ([]*ContentEntity, error) {
	panic("not implemented")
}

func (*ContentRepositoryImpl) into(in *tachyoncms.AggregateDto) *ContentEntity {
	return New(in.ID, in.CreatedAt, in.UpdatedAt, in.Data)
}
