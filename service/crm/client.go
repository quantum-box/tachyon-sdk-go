package tachyoncrm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	myContext "github.com/quantum-box/tachyon-sdk-go/internal/context"
	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TachyonCrmDriver interface {
	GetByMail(ctx context.Context, aggregationName, mail string) error
	Create(ctx context.Context, in *CustomerDto) error
	Update(ctx context.Context, in *CustomerDto) error
	Delete(ctx context.Context, aggregationName, id string) error
}

var _ TachyonCrmDriver = &Client{}

type Client struct {
	connection crmpb.CrmApiClient
	config     *tachyon.Config
}

func NewCrmClient(config *tachyon.Config) (*Client, error) {
	cc := new(Client)
	conn, err := newConnnection()
	if err != nil {
		return nil, err
	}
	cc.connection = crmpb.NewCrmApiClient(conn)
	if config.AppID == "" || config.ProjectID == "" {
		return nil, errors.New("appID and ProjectID cannot be empty")
	}
	cc.config = config
	return cc, nil
}

func newConnnection() (*grpc.ClientConn, error) {
	endpoint := os.Getenv("TACHYON_CRM_ENDPOINT")
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
