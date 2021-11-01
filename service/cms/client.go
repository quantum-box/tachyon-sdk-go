package tachyoncms

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TachyonCmsDriver interface {
	// TODO: aggregationname, context
	GetById(id string) (*AggregateDto, error)
	FindAll() ([]*AggregateDto, error)
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

func (c *Client) GetById(id string) (*AggregateDto, error) {
	// TODO: token
	md := metadata.New(map[string]string{"authorization": "Bearer some-auth-token"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.connection.GetById(ctx, &cmspb.GetRequest{
		Id: id, AggregationName: "test"})
	if err != nil {
		return nil, err
	}
	return c.from(res)
}

func (c *Client) FindAll() ([]*AggregateDto, error) {
	panic("not implemented")
}

func (*Client) from(in *cmspb.GetResponse) (*AggregateDto, error) {
	createdAt, err := time.Parse(time.RFC3339, in.Entity.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, in.Entity.UpdatedAt)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if in.Entity.DeletedAt != nil {
		*deletedAt, err = time.Parse(time.RFC3339, *in.Entity.DeletedAt)
		if err != nil {
			return nil, err
		}
	}
	var contentData map[string]interface{}
	if err = json.Unmarshal(in.Entity.Data, &contentData); err != nil {
		return nil, err
	}
	return &AggregateDto{
		ID:        in.Entity.Id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Data:      contentData,
	}, nil
}
