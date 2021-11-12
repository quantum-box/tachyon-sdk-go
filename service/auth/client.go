package tachyonauth

import (
	"context"
	"fmt"
	"log"
	"os"

	authpb "github.com/quantum-box/tachyon-sdk-go/service/auth/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon/config"
	"google.golang.org/grpc"
)

type TachyonAuthorityDriver interface {
	Verify(ctx context.Context, token string) error
}

// type checker
var _ TachyonAuthorityDriver = &Client{}

type Client struct {
	connection authpb.AuthorityApiClient
	config     *config.Config
}

func New(cfg *config.Config) (*Client, error) {
	cc := new(Client)
	//defer conn.Close()
	conn, err := newConnnection()
	if err != nil {
		return nil, err
	}
	cc.connection = authpb.NewAuthorityApiClient(conn)
	cc.config, err = config.New(cfg.ProjectID, cfg.AppID)
	if err != nil {
		return nil, err
	}
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
