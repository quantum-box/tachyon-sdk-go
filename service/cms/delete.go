package tachyoncms

import (
	"context"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) Delete(ctx context.Context, id, aggregationName string) error {
	if _, err := c.connection.Delete(ctx, &cmspb.DeleteRequest{
		Id:              id,
		AggregationName: aggregationName,
	}); err != nil {
		return err
	}
	return nil
}
