package resources

import (
	"context"
	"os"

	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
	terraform_providers "github.com/selefra/selefra-provider-sdk/terraform/provider"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.3"

func GetSelefraTerraformProvider() *selefra_terraform_schema.SelefraTerraformProvider {
	return &selefra_terraform_schema.SelefraTerraformProvider{
		Name:		"auth0",
		Version:	Version,
		ResourceList:	getResources(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var conf Config
				if err := config.Unmarshal(&conf); err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s, You can view the documentation: https://www.selefra.io/docs/providers-connector/auth0", err.Error())
				}
				diagnostics := schema.NewDiagnostics()
				client, err := newClient(clientMeta, conf)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("Init client error: %s, You can view the documentation: https://www.selefra.io/docs/providers-connector/auth0", err.Error())
				}

				// run terraform providers
				if clientMeta.Runtime().Workspace != "" {
					providerSaveDirectory := clientMeta.Runtime().Workspace + "/" + clientMeta.Runtime().ProviderName + "/" + clientMeta.Runtime().ProviderVersion
					providerFileSlice := getTerraformProviderExecuteFileSlice()
					providerExecFilePath, err := terraform_providers.NewProviderDownloader(providerFileSlice).Download(providerSaveDirectory)
					if err != nil {
						return nil, diagnostics.AddError(err)
					}
					bridge := bridge.NewTerraformBridge(providerExecFilePath)

					// read terraform config from selefra provider's config file
					terraformProviderConfig := make(map[string]any, 0)
					if config != nil {
						err := config.Unmarshal(&terraformProviderConfig)
						if err != nil {
							return nil, schema.NewDiagnostics().AddError(err)
						}
					}

					err = bridge.StartBridge(context.Background(), terraformProviderConfig)
					if err != nil {
						return nil, diagnostics.AddError(err)
					}
					client.TerraformBridge = bridge
				}

				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# Config Your Auth0
# domain: <Your auth0 domain>
# client_id: <Your client id>
# client_secret: <Your client secret>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var conf Config
				if err := config.Unmarshal(&conf); err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if config.GetString("auth0_domain") == "" && config.GetString("client_id") == "" && config.GetString("client_secret") == "" {
					if os.Getenv("AUTH0_DOMAIN") == "" && os.Getenv("AUTH0_CLIENT_ID") == "" && os.Getenv("AUTH0_CLIENT_SECRET") == "" {
						return schema.NewDiagnostics().AddErrorMsg("No valid configuration was found in the configuration file or in the environment variables, You can view the documentation: https://www.selefra.io/docs/providers-connector/auth0")
					}
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList:	[]string{},
			DataSourcePullResultAutoExpand:		true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}

func getTerraformProviderExecuteFileSlice() []*terraform_providers.TerraformProviderFile {
	providerFileSlice := make([]*terraform_providers.TerraformProviderFile, 0)

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_darwin_amd64.zip",
		Arch:			"amd64",
		OS:			"darwin",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_darwin_arm64.zip",
		Arch:			"arm64",
		OS:			"darwin",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_freebsd_386.zip",
		Arch:			"386",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_freebsd_amd64.zip",
		Arch:			"amd64",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_freebsd_arm.zip",
		Arch:			"arm",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_freebsd_arm64.zip",
		Arch:			"arm64",
		OS:			"freebsd",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_linux_386.zip",
		Arch:			"386",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_linux_amd64.zip",
		Arch:			"amd64",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_linux_arm.zip",
		Arch:			"arm",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_linux_arm64.zip",
		Arch:			"arm64",
		OS:			"linux",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_windows_386.zip",
		Arch:			"386",
		OS:			"windows",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_windows_amd64.zip",
		Arch:			"amd64",
		OS:			"windows",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_windows_arm.zip",
		Arch:			"arm",
		OS:			"windows",
	})

	providerFileSlice = append(providerFileSlice, &terraform_providers.TerraformProviderFile{
		ProviderName:		"",
		ProviderVersion:	"v0.44.0",
		DownloadUrl:		"https://github.com/auth0/terraform-provider-auth0/releases/download/v0.44.0/terraform-provider-auth0_0.44.0_windows_arm64.zip",
		Arch:			"arm64",
		OS:			"windows",
	})

	return providerFileSlice
}

func getResources() []*selefra_terraform_schema.SelefraTerraformResource {
	return []*selefra_terraform_schema.SelefraTerraformResource{
		// GetResource_example(),
	}
}
