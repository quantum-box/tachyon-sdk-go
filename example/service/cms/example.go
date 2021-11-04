package main

import (
	"context"
	"encoding/json"
	"time"

	tachyoncms "github.com/quantum-box/tachyon-sdk-go/service/cms"
)

//
// domain layer
//
// entity
type TestEntity struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	data      map[string]interface{}
}

func New(id string, createdAt, updatedAt time.Time, data map[string]interface{}) *TestEntity {
	return &TestEntity{id, createdAt, updatedAt, data}
}

// repository interface
type ContentRepository interface {
	GetById(ctx context.Context, id string) (*TestEntity, error)
	FindAll(ctx context.Context) ([]*TestEntity, error)
}

//
// interface adapter layer
//
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
	GetById(ctx context.Context, id string) (*tachyoncms.AggregateDto, error)
	FindAll(ctx context.Context) ([]*tachyoncms.AggregateDto, error)

	// Create(ctx context.Context) error
}

func (r *ContentRepositoryImpl) GetById(ctx context.Context, id string) (*TestEntity, error) {
	dto, err := r.cms.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.into(dto), nil

}

func (*ContentRepositoryImpl) FindAll(ctx context.Context) ([]*TestEntity, error) {
	panic("not implemented")
}

func (*ContentRepositoryImpl) into(in *tachyoncms.AggregateDto) *TestEntity {
	return New(in.ID, in.CreatedAt, in.UpdatedAt, in.Data)
}

//
// entrypoint
//
func main() {
	ctx := context.Background()
	cmsClient, err := tachyoncms.NewCmsClient()
	if err != nil {
		panic(err)
	}
	testRepo := NewContentRepositoryImpl(cmsClient, "test")
	entity, err := testRepo.cms.GetById(ctx, "01FKNB2BK5JA8M37586Z8673AG")
	if err != nil {
		panic(err)
	}
	ou, _ := json.Marshal(entity)
	println(string(ou))
}
