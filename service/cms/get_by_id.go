package tachyoncms

import (
	"context"
	"encoding/json"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func (c *Client) GetById(ctx context.Context, id string) (*AggregateDto, error) {

	res, err := c.connection.GetById(ctx, &cmspb.GetRequest{
		Id: id, AggregationName: "test"})
	if err != nil {
		return nil, err
	}
	return fromGetResponse(res)
}

func fromGetResponse(in *cmspb.GetResponse) (*AggregateDto, error) {
	createdAt, err := time.Parse(time.RFC3339, in.RawContent.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, in.RawContent.UpdatedAt)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if in.RawContent.DeletedAt != nil {
		*deletedAt, err = time.Parse(time.RFC3339, *in.RawContent.DeletedAt)
		if err != nil {
			return nil, err
		}
	}
	var contentData map[string]interface{}
	if err = json.Unmarshal(in.RawContent.Data, &contentData); err != nil {
		return nil, err
	}
	return &AggregateDto{
		ID:        in.RawContent.Id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Data:      contentData,
	}, nil
}
