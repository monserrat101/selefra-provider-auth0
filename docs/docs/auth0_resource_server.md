# Table: auth0_resource_server

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enforce_policies | bool | X | √ | If this setting is enabled, RBAC authorization policies will be enforced for this API. Role and permission assignments will be evaluated during the login transaction. | 
| signing_secret | string | X | √ | Secret used to sign tokens when using symmetric algorithms (HS256). | 
| token_lifetime_for_web | float | X | √ | Number of seconds during which access tokens issued for this resource server via implicit or hybrid flows remain valid. Cannot be greater than the `token_lifetime` value. | 
| token_dialect | string | X | √ | Dialect of access tokens that should be issued for this resource server. Options include `access_token` or `access_token_authz`. If this setting is set to `access_token_authz`, the Permissions claim will be added to the access token. Only available if RBAC (`enforce_policies`) is enabled for this API. | 
| verification_location | string | X | √ | URL from which to retrieve JWKs for this resource server. Used for verifying the JWT sent to Auth0 for token introspection. | 
| allow_offline_access | bool | X | √ | Indicates whether refresh tokens can be issued for this resource server. | 
| id | string | X | √ |  | 
| skip_consent_for_verifiable_first_party_clients | bool | X | √ | Indicates whether to skip user consent for applications flagged as first party. | 
| scopes | json | X | √ | List of permissions (scopes) used by this resource server. | 
| identifier | string | X | √ | Unique identifier for the resource server. Used as the audience parameter for authorization calls. Cannot be changed once set. | 
| name | string | X | √ | Friendly name for the resource server. Cannot include `<` or `>` characters. | 
| options | json | X | √ | Used to store additional metadata. | 
| signing_alg | string | X | √ | Algorithm used to sign JWTs. Options include `HS256` and `RS256`. | 
| token_lifetime | float | X | √ | Number of seconds during which access tokens issued for this resource server from the token endpoint remain valid. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


