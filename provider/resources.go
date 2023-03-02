package provider

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"
)

// terraform resource: auth0_rule_config
func GetResource_auth0_rule_config() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_rule_config",
		TerraformResourceName: "auth0_rule_config",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			rulesConfigList, err := client.ApiClient.RuleConfig.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, ruleConfig := range rulesConfigList {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ArgumentMap: map[string]any{
						"key":   ruleConfig.Key,
						"value": ruleConfig.Value,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: auth0_attack_protection
// func GetResource_auth0_attack_protection() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_attack_protection",
// 		TerraformResourceName: "auth0_attack_protection",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			attackProtectionList, err := client.ApiClient.AttackProtection.GetBreachedPasswordDetection()
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}
// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{})
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: auth0_role
func GetResource_auth0_role() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_role",
		TerraformResourceName: "auth0_role",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			rolesList, err := client.ApiClient.Role.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, role := range rolesList.Roles {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *role.ID,
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_custom_domain_verification
func GetResource_auth0_custom_domain_verification() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_custom_domain_verification",
		TerraformResourceName: "auth0_custom_domain_verification",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			customDomainList, err := client.ApiClient.CustomDomain.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, customDomain := range customDomainList {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *customDomain.ID,
					ArgumentMap: map[string]any{
						"custom_domain_id": customDomain.ID,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_action
func GetResource_auth0_action() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_action",
		TerraformResourceName: "auth0_action",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			actionsList, err := client.ApiClient.Action.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, action := range actionsList.Actions {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *action.ID,
					ArgumentMap: map[string]any{
						"code":               action.Code,
						"name":               action.Name,
						"supported_triggers": action.SupportedTriggers,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: auth0_email_template
// func GetResource_auth0_email_template() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_email_template",
// 		TerraformResourceName: "auth0_email_template",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: auth0_email
func GetResource_auth0_email() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_email",
		TerraformResourceName: "auth0_email",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			email, err := client.ApiClient.Email.Read()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ArgumentMap: map[string]any{
					"credentials":          email.Credentials,
					"default_from_address": email.DefaultFromAddress,
					"name":                 email.Name,
				},
			})
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_guardian
func GetResource_auth0_guardian() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_guardian",
		TerraformResourceName: "auth0_guardian",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			policiesList, err := client.ApiClient.Guardian.MultiFactor.Policy()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, policy := range *policiesList {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ArgumentMap: map[string]any{
						"policy": policy,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_client_grant
func GetResource_auth0_client_grant() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_client_grant",
		TerraformResourceName: "auth0_client_grant",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			grantsList, err := client.ApiClient.Grant.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, grant := range grantsList.Grants {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *grant.ID,
					ArgumentMap: map[string]any{
						"audience":  grant.Audience,
						"client_id": grant.ClientID,
						"scope":     grant.Scope,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_hook
func GetResource_auth0_hook() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_hook",
		TerraformResourceName: "auth0_hook",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			hooksList, err := client.ApiClient.Hook.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, hook := range hooksList.Hooks {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *hook.ID,
					ArgumentMap: map[string]any{
						"name":      hook.Name,
						"script":    hook.Script,
						"tigger_id": hook.TriggerID,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_organization_connection
func GetResource_auth0_organization_connection() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_organization_connection",
		TerraformResourceName: "auth0_organization_connection",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			organizationsList, err := client.ApiClient.Organization.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, organization := range organizationsList.Organizations {
				orizationInvitations, err := client.ApiClient.Organization.Invitations(*organization.ID)
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, orizationInvitation := range orizationInvitations.OrganizationInvitations {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: *orizationInvitation.ID,
						ArgumentMap: map[string]any{
							"connection_id":   orizationInvitation.ConnectionID,
							"organization_id": orizationInvitation.OrganizationID,
						},
					})
				}
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_resource_server
func GetResource_auth0_resource_server() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_resource_server",
		TerraformResourceName: "auth0_resource_server",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			resourceServersList, err := client.ApiClient.ResourceServer.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, resourceServer := range resourceServersList.ResourceServers {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *resourceServer.ID,
					ArgumentMap: map[string]any{
						"identifier": resourceServer.Identifier,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// TODO
// // terraform resource: auth0_branding
// func GetResource_auth0_branding() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_branding",
// 		TerraformResourceName: "auth0_branding",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			branding, err := client.ApiClient.Branding.Read()
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}
// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			resourceRequestParamSlice = append(resourceRequestParamSlice,	&selefra_terraform_schema.ResourceRequestParam{
// 				ArgumentMap: map[string]any{
// 					"favicon_url": branding.FaviconURL,
// 				},
// 			})
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: auth0_user
func GetResource_auth0_user() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_user",
		TerraformResourceName: "auth0_user",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			usersList, err := client.ApiClient.User.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, user := range usersList.Users {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *user.ID,
					ArgumentMap: map[string]any{
						"connection_name": user.Connection,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// Not Support
// // terraform resource: auth0_tenant
// func GetResource_auth0_tenant() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_tenant",
// 		TerraformResourceName: "auth0_tenant",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			tenant, err := client.ApiClient.Tenant.Read()
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}
// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 				ArgumentMap: map[string]any{
// 					"identifier_first":               prompt.IdentifierFirst,
// 					"webauthn_platform_first_factor": prompt.WebAuthnPlatformFirstFactor,
// 					"universal_login_experience":     prompt.UniversalLoginExperience,
// 				},
// 			})
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: auth0_prompt
func GetResource_auth0_prompt() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_prompt",
		TerraformResourceName: "auth0_prompt",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			prompt, err := client.ApiClient.Prompt.Read()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ArgumentMap: map[string]any{
					"identifier_first":               prompt.IdentifierFirst,
					"webauthn_platform_first_factor": prompt.WebAuthnPlatformFirstFactor,
					"universal_login_experience":     prompt.UniversalLoginExperience,
				},
			})
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_organization
func GetResource_auth0_organization() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_organization",
		TerraformResourceName: "auth0_organization",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			organizationList, err := client.ApiClient.Organization.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, organization := range organizationList.Organizations {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *organization.ID,
					ArgumentMap: map[string]any{
						"name": organization.Name,
					},
				})
			}
			return nil, nil
		},
	}
}

// terraform resource: auth0_trigger_binding
func GetResource_auth0_trigger_binding() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_trigger_binding",
		TerraformResourceName: "auth0_trigger_binding",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			triggersList, err := client.ApiClient.Action.Triggers()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, trigger := range triggersList.Triggers {
				bindingsList, err := client.ApiClient.Action.Bindings(*trigger.ID)
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}

				for _, binding := range bindingsList.Bindings {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: *binding.ID,
						ArgumentMap: map[string]any{
							"actions": binding.Action,
							"trigger": binding.TriggerID,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_branding_theme
func GetResource_auth0_branding_theme() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_branding_theme",
		TerraformResourceName: "auth0_branding_theme",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			theme, err := client.ApiClient.BrandingTheme.Default()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: *theme.ID,
				ArgumentMap: map[string]any{
					"borders":         theme.Borders,
					"colors":          theme.Colors,
					"fonts":           theme.Fonts,
					"page_background": theme.PageBackground,
					"widget":          theme.Widget,
				},
			})

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_client
func GetResource_auth0_client() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_client",
		TerraformResourceName: "auth0_client",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			clientsList, err := client.ApiClient.Client.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, client := range clientsList.Clients {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *client.ClientID,
					ArgumentMap: map[string]any{
						"name": client.Name,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// TODO
// // find connection client API
// // terraform resource: auth0_connection_client
// func GetResource_auth0_connection_client() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_connection_client",
// 		TerraformResourceName: "auth0_connection_client",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			connectionsList, err := client.ApiClient.Connection.List()
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}
// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			for _, connection := range connectionsList.Connections {
// 				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 					ID: *connection.ID,
// 					ArgumentMap: map[string]any{
// 						"client_id": connection.DisplayName,
// 						"connection_id": connection.
// 					},
// 				})
// 			}
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: auth0_log_stream
func GetResource_auth0_log_stream() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_log_stream",
		TerraformResourceName: "auth0_log_stream",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			logStreamList, err := client.ApiClient.LogStream.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, logStream := range logStreamList {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *logStream.ID,
					ArgumentMap: map[string]any{
						"name": logStream.Name,
						"sink": logStream.Sink,
						"type": logStream.Type,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// // terraform resource: auth0_prompt_custom_text
// func GetResource_auth0_prompt_custom_text() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "auth0_prompt_custom_text",
// 		TerraformResourceName: "auth0_prompt_custom_text",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			client := taskClient.(*Client)
// 			prompt, err := client.ApiClient.Prompt.Read()
// 			if err != nil {
// 				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			}

// 			// promptCustomText := client.ApiClient.Prompt.CustomText(prompt, prompt.)

// 			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 				// ID:promot,
// 			})
// 			return resourceRequestParamSlice, nil
// 		},
// 	}
// }

// terraform resource: auth0_custom_domain
func GetResource_auth0_custom_domain() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_custom_domain",
		TerraformResourceName: "auth0_custom_domain",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			customDomainList, err := client.ApiClient.CustomDomain.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, customDomain := range customDomainList {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *customDomain.ID,
					ArgumentMap: map[string]any{
						"domain": customDomain.Domain,
						"type":   customDomain.Type,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_organization_member
func GetResource_auth0_organization_member() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_organization_member",
		TerraformResourceName: "auth0_organization_member",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			organizationList, err := client.ApiClient.Organization.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, organization := range organizationList.Organizations {
				organzationMemberList, err := client.ApiClient.Organization.Members(*organization.ID)
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, member := range organzationMemberList.Members {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ArgumentMap: map[string]any{
							"organization_id": organization.ID,
							"user_id":         member.UserID,
						},
					})
				}
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_connection
func GetResource_auth0_connection() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_connection",
		TerraformResourceName: "auth0_connection",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			connectionsList, err := client.ApiClient.Connection.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, connection := range connectionsList.Connections {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *connection.ID,
					ArgumentMap: map[string]any{
						"name":     connection.Name,
						"strategy": connection.Strategy,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_global_client
func GetResource_auth0_global_client() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_global_client",
		TerraformResourceName: "auth0_global_client",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			clientsList, err := client.ApiClient.Client.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, client := range clientsList.Clients {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *client.ClientID,
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: auth0_rule
func GetResource_auth0_rule() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "auth0_rule",
		TerraformResourceName: "auth0_rule",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)
			rulesList, err := client.ApiClient.Rule.List()
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, rule := range rulesList.Rules {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: *rule.ID,
					ArgumentMap: map[string]any{
						"name":   rule.Name,
						"script": rule.Script,
					},
				})
			}
			return resourceRequestParamSlice, nil
		},
	}
}
