package tachyoncms

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	myContext "github.com/quantum-box/tachyon-sdk-go/internal/context"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TachyonCmsDriver interface {
	GetById(ctx context.Context, aggregationName, id string) (*AggregateDto, error)
	FindAll(ctx context.Context, aggregationName string) ([]*AggregateDto, error)
	FindByPath(ctx context.Context, aggregationName string, paths []string, value string) ([]*AggregateDto, error)

	Create(ctx context.Context, aggregationName string, in *AggregateDto) error
	Update(ctx context.Context, aggregationName string, in *AggregateDto) error
	Delete(ctx context.Context, aggregationName, id string) error
}

var _ TachyonCmsDriver = &Client{}

type Client struct {
	connection cmspb.CmsApiClient
	config     *tachyon.Config
}

func NewCmsClient(config *tachyon.Config) (*Client, error) {
	cc := new(Client)
	//defer conn.Close()
	conn, err := newConnnection()
	if err != nil {
		return nil, err
	}
	cc.connection = cmspb.NewCmsApiClient(conn)
	if config.AppID == "" || config.ProjectID == "" {
		return nil, errors.New("appID and ProjectID cannot be empty")
	}
	cc.config = config
	return cc, nil
}

func newConnnection() (*grpc.ClientConn, error) {
	endpoint := os.Getenv("TACHYON_CMS_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:50051"
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, fmt.Errorf("ConnectionFailureErr:%v", err)
	}
	return conn, nil
}

func (c *Client) withConfig(ctx context.Context) (context.Context, error) {
	token, err := myContext.CurrentAuthzToken(ctx)
	if err != nil {
		return nil, err
	}
	md := metadata.New(map[string]string{
		"authorization": token,
		"project_id":    c.config.ProjectID,
		"app_id":        c.config.AppID,
	})
	return metadata.NewOutgoingContext(ctx, md), nil
}
