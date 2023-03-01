package provider

import (
	"errors"
	"log"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"

	"github.com/auth0/go-auth0/management"
)

type Config struct {
	Domain       string `json:"domain"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (c *Config) isVaild() bool {
	return c.Domain != "" && c.ClientID != "" && c.ClientSecret != ""
}

func (c *Config) isEnvVaild() bool {
	c.Domain = os.Getenv("AUTH0_DOMAIN")
	if c.Domain == "" {
		log.Println("cannot find your AUTH0_DOMAIN local environment variable.")
	}

	c.ClientID = os.Getenv("AUTH0_CLIENT_ID")
	if c.ClientID == "" {
		log.Println("cannot find your AUTH0_CLIENT_ID local environment variable.")
	}

	c.ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	if c.ClientSecret == "" {
		log.Println("cannot find your AUTH0_CLIENT_SECRET local environment variable.")
	}
	return c.isVaild()
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge

	Config    Config
	ApiClient *management.Management
}

func newClient(clientMeta *schema.ClientMeta, config Config) (*Client, error) {
	if !config.isVaild() {
		if !config.isEnvVaild() {
			ErrorF(clientMeta, "Config Error! Cannot find your Environment Variable.")
			return nil, errors.New("Config Error!")
		}
	}

	client, err := management.New(
		config.Domain,
		management.WithClientCredentials(config.ClientID, config.ClientSecret),
	)
	if err != nil {
		ErrorF(clientMeta, "failed to initialize the auth0 management API client: %w", err)
		return nil, errors.New("failed to initialize the auth0 management API client!")
	}

	return &Client{
		ApiClient: client,
		Config:    config,
	}, nil
}

func DebugF(clientMeta *schema.ClientMeta, msg string, args ...any) {
	if clientMeta != nil {
		clientMeta.DebugF(msg, args...)
	}
}

func ErrorF(clientMeta *schema.ClientMeta, msg string, args ...any) {
	if clientMeta != nil {
		clientMeta.ErrorF(msg, args...)
	}
}
