package config

import "errors"

type Config struct {
	// project id is ulid
	ProjectID string

	// app id is globally unique ulid
	AppID string
}

func New(projectID, appID string) (*Config, error) {
	if projectID == "" || appID == "" {
		return nil, errors.New("projectID and appID must not be empty")
	}
	return &Config{
		ProjectID: projectID,
		AppID:     appID,
	}, nil
}
