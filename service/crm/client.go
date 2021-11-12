package tachyoncrm

import (
	"context"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

type TachyonCrmDriver interface {
	Create(ctx context.Context, in *CustomerDto) error
}

var _ TachyonCrmDriver = &Client{}

type Client struct {
	connection crmpb.CrmApiClient
	config     *tachyon.Config
}
