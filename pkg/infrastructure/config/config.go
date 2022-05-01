package config

import "time"

type Config struct {
	EndpointUrl string
	Region      string
	Timeout     time.Duration
	ID          string
	SecretKey   string
	AccessKey   string
}
