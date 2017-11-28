package nifcloud

import (
	"github.com/heriet/funicula/nifcloud/credential"
)

// Config for nifcloud client
type Config struct {
	Region string
	Endpoint string

	Credential *credential.Credential
}