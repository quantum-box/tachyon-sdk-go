package main

import (
	"encoding/json"
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
	cms             TachyonCmsDriver
	aggregationName string
}

func NewContentRepositoryImpl(in TachyonCmsDriver, aggregationName string) *ContentRepositoryImpl {
	return &ContentRepositoryImpl{in, aggregationName}
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

// entrypoint
func main() {
	cmsClient, err := tachyoncms.NewCmsClient()
	if err != nil {
		panic(err)
	}
	testRepo := NewContentRepositoryImpl(cmsClient, "test")
	entity, err := testRepo.cms.GetById("01FK5DK219TR117F5KE96TSPN3")
	if err != nil {
		panic(err)
	}
    ou, _ := json.Marshal(entity)
	println(string(ou))
}
