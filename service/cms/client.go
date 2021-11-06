package tachyoncms

import (
	"context"
	"fmt"
	"log"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"google.golang.org/grpc"
)

type TachyonCmsDriver interface {
	// TODO: aggregationname, context
	GetById(ctx context.Context, id string) (*AggregateDto, error)
	FindAll(ctx context.Context, aggregationName string) ([]*AggregateDto, error)
	FindByPath(ctx context.Context, aggregationName string, paths []string, value string) ([]*AggregateDto, error)

	Create(ctx context.Context, in *AggregateDto) error
	Update(ctx context.Context, in *AggregateDto, aggregationName string) error
	Delete(ctx context.Context, id, aggregationName string) error
}

var _ TachyonCmsDriver = &Client{}

type Client struct {
	connection cmspb.CmsApiClient
}

func NewCmsClient() (*Client, error) {
	cc := new(Client)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, fmt.Errorf("ConnectionFailureErr:%v", err)
	}
	//defer conn.Close()
	cc.connection = cmspb.NewCmsApiClient(conn)
	return cc, nil
}
