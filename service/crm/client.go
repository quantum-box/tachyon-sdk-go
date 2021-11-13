package tachyoncrm

import (
	"context"
	"fmt"
	"log"
	"os"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
	"google.golang.org/grpc"
)

type TachyonCrmDriver interface {
	Create(ctx context.Context, in *CustomerDto) error
}

var _ TachyonCrmDriver = &Client{}

type Client struct {
	connection crmpb.CrmApiClient
	config     *tachyon.Config
}

func NewCrmClient() (*Client, error) {
	cc := new(Client)
	conn, err := newConnnection()
	if err != nil {
		return nil, err
	}
	cc.connection = crmpb.NewCrmApiClient(conn)
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
