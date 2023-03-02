# Table: auth0_custom_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tls_policy | string | X | √ | TLS policy for the custom domain. Available options are: `compatible` or `recommended`. Compatible includes TLS 1.0, 1.1, 1.2, and recommended only includes TLS 1.2. Cannot be set on self_managed domains. | 
| type | string | X | √ | Provisioning type for the custom domain. Options include `auth0_managed_certs` and `self_managed_certs`. | 
| verification | json | X | √ | Configuration settings for verification. | 
| custom_client_ip_header | string | X | √ | The HTTP header to fetch the client's IP address. Cannot be set on auth0_managed domains. | 
| domain | string | X | √ | Name of the custom domain. | 
| id | string | X | √ |  | 
| origin_domain_name | string | X | √ | Once the configuration status is `ready`, the DNS name of the Auth0 origin server that handles traffic for the custom domain. | 
| primary | bool | X | √ | Indicates whether this is a primary domain. | 
| status | string | X | √ | Configuration status for the custom domain. Options include `disabled`, `pending`, `pending_verification`, and `ready`. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


