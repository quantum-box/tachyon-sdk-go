package tachyoncms

import (
	"context"
	"encoding/json"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) Create(ctx context.Context, in *AggregateDto) error {

	req, err := fromAggregateDto(in)
	if err != nil {
		return err
	}
	_, err = c.connection.Create(ctx, req)
	if err != nil {
		return err
	}
	return err
}

func fromAggregateDto(in *AggregateDto) (*cmspb.CreateRequest, error) {
	var deletedAt *string
	if in.DeletedAt != nil {
		d := in.DeletedAt.Format(time.RFC3339)
		deletedAt = &d
	}
	bytes, err := json.Marshal(in.Data)
	if err != nil {
		return nil, err
	}
	return &cmspb.CreateRequest{
        AggregationName: "test",
		Entity: &cmspb.RawContent{
			Id:        in.ID,
			CreatedAt: in.CreatedAt.Format(time.RFC3339),
			UpdatedAt: in.UpdatedAt.Format(time.RFC3339),
			DeletedAt: deletedAt,
			Data:      bytes,
		},
	}, nil
}
