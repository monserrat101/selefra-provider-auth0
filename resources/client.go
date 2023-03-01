package resources

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
	if c.Domain == "" || c.ClientID == "" || c.ClientSecret == "" {
		return false
	}
	return true
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge

	Config Config
	// TODO You can continue to refine your client
	ApiClient *management.Management
}

func newClient(clientMeta *schema.ClientMeta, config Config) (*Client, error) {
	if !config.isVaild() {
		config.Domain = os.Getenv("AUTH0_DOMAIN")
		config.ClientID = os.Getenv("AUTH0_CLIENT_ID")
		config.ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	}

	if !config.isVaild() {
		ErrorF(clientMeta, "Config Error!")
		return nil, errors.New("Get Config Error!")

	}

	client, err := management.New(
		config.Domain,
		management.WithClientCredentials(config.ClientID, config.ClientSecret),
	)
	if err != nil {
		log.Fatal("failed to initialize the auth0 management API client: %w", err)
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
