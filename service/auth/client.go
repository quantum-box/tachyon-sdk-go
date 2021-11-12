package tachyonauth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	authpb "github.com/quantum-box/tachyon-sdk-go/service/auth/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
	"google.golang.org/grpc"
)

type TachyonAuthorityDriver interface {
	Verify(ctx context.Context, token string) error
}

// type checker
var _ TachyonAuthorityDriver = &Client{}

type Client struct {
	connection authpb.AuthorityApiClient
	config     *tachyon.Config
}

func New(config *tachyon.Config) (*Client, error) {
	cc := new(Client)
	//defer conn.Close()
	conn, err := newConnnection()
	if err != nil {
		return nil, err
	}
	cc.connection = authpb.NewAuthorityApiClient(conn)
	if config.AppID == "" || config.ProjectID == "" {
		return nil, errors.New("appID and ProjectID cannot be empty")
	}
	cc.config = config
	return cc, nil
}

func newConnnection() (*grpc.ClientConn, error) {
	endpoint := os.Getenv("TACHYON_AUTHORITY_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:50052"
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, fmt.Errorf("ConnectionFailureErr:%v", err)
	}
	return conn, nil
}
