# Table: auth0_global_client

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| client_secret_rotation_trigger | json | X | √ | Custom metadata for the rotation. The contents of this map are arbitrary and are hashed by the provider. When the hash changes, a rotation is triggered. For example, the map could contain the user making the change, the date of the change, and a text reason for the change. For more info: [rotate-client-secret](https://auth0.com/docs/get-started/applications/rotate-client-secret). | 
| cross_origin_loc | string | X | √ | URL of the location in your site where the cross-origin verification takes place for the cross-origin auth flow when performing authentication in your own domain instead of Auth0 Universal Login page. | 
| form_template | string | X | √ | HTML form template to be used for WS-Federation. | 
| signing_keys | json | X | √ | List containing a map of the public cert of the signing key and the public cert of the signing key in PKCS7. | 
| sso_disabled | bool | X | √ | Indicates whether or not SSO is disabled. | 
| jwt_configuration | json | X | √ | Configuration settings for the JWTs issued for this client. | 
| mobile | json | X | √ | Additional configuration for native mobile apps. | 
| client_id | string | X | √ | The ID of the client. | 
| name | string | X | √ | Name of the client. | 
| id | string | X | √ |  | 
| client_secret | string | X | √ | Secret for the client. Keep this private. To access this attribute you need to add the `read:client_keys` scope to the Terraform client. Otherwise, the attribute will contain an empty string. | 
| custom_login_page | string | X | √ | The content (HTML, CSS, JS) of the custom login page. | 
| encryption_key | json | X | √ | Encryption used for WS-Fed responses with this client. | 
| is_token_endpoint_ip_header_trusted | bool | X | √ | Indicates whether the token endpoint IP header is trusted. | 
| logo_uri | string | X | √ | URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown. | 
| refresh_token | json | X | √ | Configuration settings for the refresh tokens issued for this client. | 
| client_aliases | json | X | √ | List of audiences/realms for SAML protocol. Used by the wsfed addon. | 
| allowed_logout_urls | json | X | √ | URLs that Auth0 may redirect to after logout. | 
| allowed_origins | json | X | √ | URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed. | 
| app_type | string | X | √ | Type of application the client represents. Possible values are: `native`, `spa`, `regular_web`, `non_interactive`, `sso_integration`. Specific SSO integrations types accepted as well are: `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`. | 
| allowed_clients | json | X | √ | List of applications ID's that will be allowed to make delegation request. By default, all applications will be allowed. | 
| initiate_login_uri | string | X | √ | Initiate login URI. Must be HTTPS or an empty string. | 
| sso | bool | X | √ | Applies only to SSO clients and determines whether Auth0 will handle Single Sign-On (true) or whether the identity provider will (false). | 
| web_origins | json | X | √ | URLs that represent valid web origins for use with web message response mode. | 
| addons | json | X | √ | Addons enabled for this client and their associated configurations. | 
| grant_types | json | X | √ | Types of grants that this client is authorized to use. | 
| custom_login_page_on | bool | X | √ | Indicates whether a custom login page is to be used. | 
| callbacks | json | X | √ | URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://. | 
| native_social_login | json | X | √ | Configuration settings to toggle native social login for mobile native applications. Once this is set it must stay set, with both resources set to `false` in order to change the `app_type`. | 
| organization_usage | string | X | √ | Defines how to proceed during an authentication transaction with regards to an organization. Can be `deny` (default), `allow` or `require`. | 
| cross_origin_auth | bool | X | √ | Whether this client can be used to make cross-origin authentication requests (`true`) or it is not allowed to make such requests (`false`). Requires the `coa_toggle_enabled` feature flag to be enabled on the tenant by the support team. | 
| description | string | X | √ | Description of the purpose of the client. | 
| is_first_party | bool | X | √ | Indicates whether this client is a first-party client. | 
| oidc_conformant | bool | X | √ | Indicates whether this client will conform to strict OIDC specifications. | 
| organization_require_behavior | string | X | √ | Defines how to proceed during an authentication transaction when `organization_usage = "require"`. Can be `no_prompt` (default) or `pre_login_prompt`. | 
| token_endpoint_auth_method | string | X | √ | Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters), `client_secret_basic` (client uses HTTP Basic). | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


