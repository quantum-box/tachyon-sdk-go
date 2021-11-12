package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	tachyoncms "github.com/quantum-box/tachyon-sdk-go/service/cms"
	"github.com/quantum-box/tachyon-sdk-go/tachyon/config"
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
	GetById(ctx context.Context, aggregationName, id string) (*tachyoncms.AggregateDto, error)
	FindAll(ctx context.Context, aggregationName string) ([]*tachyoncms.AggregateDto, error)

	Create(ctx context.Context, aggregationName string, in *tachyoncms.AggregateDto) error
}

func (r *ContentRepositoryImpl) GetById(ctx context.Context, id string) (*TestEntity, error) {
	ctx = testhelper.NewContextWithToken()
	dto, err := r.cms.GetById(ctx, "test", id)
	if err != nil {
		return nil, err
	}
	return r.into(dto), nil

}

func (*ContentRepositoryImpl) FindAll(ctx context.Context) ([]*TestEntity, error) {
	panic("not implemented")
}

func (r *ContentRepositoryImpl) Create(ctx context.Context, in *TestEntity) error {
	ctx = testhelper.NewContextWithToken()
	if err := r.cms.Create(ctx, "test", r.from(in)); err != nil {
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

	cfg := &config.Config{
		ProjectID: os.Getenv("TACHYON_PROJECT_ID"),
		AppID:     os.Getenv("TACHYON_APP_ID"),
	}

	// craete cms sdk client
	cmsClient, err := tachyoncms.NewCmsClient(cfg)
	if err != nil {
		panic(err)
	}

	testRepo := NewContentRepositoryImpl(cmsClient, "test")
	entity, err := testRepo.GetById(ctx, "01FKNYSFQ4P7F8ETGR0ZGE9N6B")
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
