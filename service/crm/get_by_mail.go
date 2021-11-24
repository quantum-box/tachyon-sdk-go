package tachyoncrm

import (
	"context"
	"time"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
)

func (c *Client) GetByMail(ctx context.Context, aggregationName, mail string) (*CustomerDto, error) {
	ctx, err := c.withConfig(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.connection.GetByMail(ctx, &crmpb.GetRequest{
		AggregationName: aggregationName,
		Mail:            mail,
	})
	if err != nil {
		return nil, err
	}
	return fromGetResponse(res)
}

func fromGetResponse(in *crmpb.GetResponse) (*CustomerDto, error) {
	registeredAt, err := time.Parse(time.RFC3339, in.RawCustomer.RegisteredAt)
	if err != nil {
		return nil, err
	}
	lastSignedInAt, err := time.Parse(time.RFC3339, in.RawCustomer.LastSignedInAt)
	if err != nil {
		return nil, err
	}
	return &CustomerDto{
		ID:             in.RawCustomer.Id,
		RegisteredAt:   registeredAt,
		LastSignedInAt: lastSignedInAt,
		//ここoptionだけど大丈夫かな？？？
		Mail: in.RawCustomer.Mail,
	}, nil
}
