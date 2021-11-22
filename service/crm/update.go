package tachyoncrm

import (
	"context"
	"time"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
)

func (c *Client) Update(ctx context.Context, in *CustomerDto) error {
	ctx, err := c.withConfig(ctx)
	if err != nil {
		return err
	}
	req := fromCustomerUpdate(in)
	_, err = c.connection.Update(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func fromCustomerUpdate(in *CustomerDto) *crmpb.UpdateRequest {
	return &crmpb.UpdateRequest{
		AggregationName: "testttt",
		RawCustomer: &crmpb.RawCustomer{
			Id:             in.ID,
			RegisteredAt:   in.RegisteredAt.Format(time.RFC3339),
			LastSignedInAt: in.LastSignedInAt.Format(time.RFC3339),
		},
	}
}
