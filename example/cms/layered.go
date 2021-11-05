package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
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
	fmt.Println(id)
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

	Create(ctx context.Context, in *tachyoncms.AggregateDto) error
}

func (r *ContentRepositoryImpl) GetById(ctx context.Context, id string) (*TestEntity, error) {
	ctx = tachyoncms.WithAuth(ctx, "Bearer some-auth-token")
	dto, err := r.cms.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.into(dto), nil

}

func (*ContentRepositoryImpl) FindAll(ctx context.Context) ([]*TestEntity, error) {
	panic("not implemented")
}

func (r *ContentRepositoryImpl) Create(ctx context.Context, in *TestEntity) error {
	ctx = tachyoncms.WithAuth(ctx, "Bearer some-auth-token")
	if err := r.cms.Create(ctx, r.from(in)); err != nil {
		return err
	}
	return nil
}

func (*ContentRepositoryImpl) from(entity *TestEntity) *tachyoncms.AggregateDto {
	return &tachyoncms.AggregateDto{
		ID:        entity.id,
		CreatedAt: entity.createdAt,
		UpdatedAt: entity.updatedAt,
		DeletedAt: nil,
		Data:      entity.data,
	}
}
func (*ContentRepositoryImpl) into(in *tachyoncms.AggregateDto) *TestEntity {
	return New(in.ID, in.CreatedAt, in.UpdatedAt, in.Data)
}

//
// entrypoint
//
func main() {
	ctx := context.Background()

    // craete cms sdk client
	cmsClient, err := tachyoncms.NewCmsClient()
	if err != nil {
		panic(err)
	}

	testRepo := NewContentRepositoryImpl(cmsClient, "test")
	entity, err := testRepo.GetById(ctx, "01FKNB2BK5JA8M37586Z8673AG")
	if err != nil {
		panic(err)
	}
	entity.id = NewUlID()
	fmt.Println(entity.id)
	if err = testRepo.Create(ctx, entity); err != nil {
		panic(err)
	}
}

func NewUlID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
