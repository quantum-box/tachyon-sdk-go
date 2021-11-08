package tachyoncms

import (
	"context"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) FindByPath(ctx context.Context, aggregationName string, paths []string, value string) ([]*AggregateDto, error) {
	ctx, err := c.withConfig(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.connection.FindByPath(ctx, &cmspb.FindByPathRequest{
		AggregationName: aggregationName,
		Paths:           paths,
		Value:           value,
	})
	if err != nil {
		return nil, err
	}
	raws, err := convAggregateRaws(res)
	if err != nil {
		return nil, err
	}
	return raws, err
}
