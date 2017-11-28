package rdb

import (
	"fmt"

	"github.com/heriet/funicula/nifcloud"
	"github.com/heriet/funicula/nifcloud/client"
	// "github.com/heriet/funicula/nifcloud/signature/v2"
)

// RDB is service client for NIFCLOUD RDB
type RDB struct {
	Client *client.Client
}

const (
	// ServiceName is name of Service
	ServiceName = "rdb"
)

// New creates RDB service client
func New(cfg *nifcloud.Config) *RDB {
	svc := &RDB{
		Client: client.New(cfg),
	}

	if cfg.Region == "" {
		cfg.Region = "east-1"
	}

	if cfg.Endpoint == "" {
		cfg.Endpoint = svc.configEndpointFromRegion(cfg.Region)
	}

	return svc
}

func (svc *RDB) configEndpointFromRegion(region string) string {
	endpointRegion := ""

	switch region {
	case "east-1":
		endpointRegion = "jp-east-1"
	case "east-2":
		endpointRegion = "jp-east-2"
	case "east-3":
		endpointRegion = "jp-east-3"
	case "east-4":
		endpointRegion = "jp-east-4"
	case "west-4":
		endpointRegion = "jp-west-1"
	default:
		endpointRegion = "jp-east-1"
	}

	return fmt.Sprintf("https://rdb.%s.api.cloud.nifty.com", endpointRegion)
}
