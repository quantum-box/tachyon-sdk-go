package tachyoncrm

import (
	"context"
	"time"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
)

func (c *Client) Create(ctx context.Context, in *CustomerDto) error {
	req := fromCustomerDto(in)
	_, err := c.connection.Create(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func fromCustomerDto(in *CustomerDto) *crmpb.CreateRequest {
	return &crmpb.CreateRequest{
		AggregationName: "aaaagrename",
		RawCustomer: &crmpb.RawCustomer{
			Id:             in.ID,
			RegisteredAt:   in.RegisteredAt.Format(time.RFC3339),
			LastSignedInAt: in.LastSignedInAt.Format(time.RFC3339),
		},
	}
}
