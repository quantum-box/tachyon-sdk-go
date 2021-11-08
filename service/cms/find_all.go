package tachyoncms

import (
	"context"
	"encoding/json"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) FindAll(ctx context.Context, aggregationName string) ([]*AggregateDto, error) {
	ctx, err := c.withConfig(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.connection.FindAll(ctx, &cmspb.FindRequest{AggregationName: aggregationName})
	if err != nil {
		return nil, err
	}
	raws, err := convAggregateRaws(res)
	if err != nil {
		return nil, err
	}
	return raws, nil
}

func convAggregateRaws(in *cmspb.FindResponse) ([]*AggregateDto, error) {
	raws := make([]*AggregateDto, 0, len(in.RawContents))
	for _, v := range in.RawContents {
		raw, err := convAggregateRaw(v)
		if err != nil {
			return nil, err
		}
		raws = append(raws, raw)
	}
	return raws, nil
}

func convAggregateRaw(in *cmspb.RawContent) (*AggregateDto, error) {
	createdAt, err := time.Parse(time.RFC3339, in.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, in.UpdatedAt)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if in.DeletedAt != nil {
		*deletedAt, err = time.Parse(time.RFC3339, *in.DeletedAt)
		if err != nil {
			return nil, err
		}
	}
	var contentData map[string]interface{}
	if err = json.Unmarshal(in.Data, &contentData); err != nil {
		return nil, err
	}
	return &AggregateDto{
		ID:        in.Id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Data:      contentData,
	}, nil
}
