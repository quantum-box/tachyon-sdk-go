package tachyoncms

import (
	"context"
	"encoding/json"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) Update(ctx context.Context, in *AggregateDto, aggregationName string) error {
	req, err := convUpdateRequest(in, aggregationName)
	if err != nil {
		return err
	}
	_, err = c.connection.Update(ctx, req)
	if err != nil {
		return err
	}
	return err
}

func convUpdateRequest(in *AggregateDto, aggregationName string) (*cmspb.UpdateRequest, error) {
	var deletedAt *string
	if in.DeletedAt != nil {
		d := in.DeletedAt.Format(time.RFC3339)
		deletedAt = &d
	}
	bytes, err := json.Marshal(in.Data)
	if err != nil {
		return nil, err
	}
	return &cmspb.UpdateRequest{
		AggregationName: aggregationName,
		RawContent: &cmspb.RawContent{
			Id:        in.ID,
			CreatedAt: in.CreatedAt.Format(time.RFC3339),
			UpdatedAt: in.UpdatedAt.Format(time.RFC3339),
			DeletedAt: deletedAt,
			Data:      bytes,
		},
	}, nil
}
