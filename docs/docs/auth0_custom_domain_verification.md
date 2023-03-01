# Table: auth0_custom_domain_verification

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| origin_domain_name | string | X | √ | The DNS name of the Auth0 origin server that handles traffic for the custom domain. | 
| timeouts | json | X | √ |  | 
| cname_api_key | string | X | √ | The value of the `cname-api-key` header to send when forwarding requests. Only present if the type of the custom domain is `self_managed_certs` and Terraform originally managed the domain's verification. | 
| custom_domain_id | string | X | √ | ID of the custom domain resource. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


