# tachyon-sdk-go

## use cms

use this interface and injection client

```go
type TachyonCmsDriver interface {
	GetById(ctx context.Context, aggregationName, id string) (*AggregateDto, error)
	FindAll(ctx context.Context, aggregationName string) ([]*AggregateDto, error)
	FindByPath(ctx context.Context, aggregationName string, paths []string, value string) ([]*AggregateDto, error)

	Create(ctx context.Context, aggregationName string, in *AggregateDto) error
	Update(ctx context.Context, aggregationName string, in *AggregateDto) error
	Delete(ctx context.Context, aggregationName, id string) error
}
```

```go
type AggregateDto struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Data      map[string]interface{}
}
```
