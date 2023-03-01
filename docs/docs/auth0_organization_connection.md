# Table: auth0_organization_connection

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the enabled connection. | 
| organization_id | string | X | √ | The ID of the organization to enable the connection for. | 
| strategy | string | X | √ | The strategy of the enabled connection. | 
| assign_membership_on_login | bool | X | √ | When true, all users that log in with this connection will be automatically granted membership in the organization. When false, users must be granted membership in the organization before logging in with this connection. | 
| connection_id | string | X | √ | The ID of the connection to enable for the organization. | 
| id | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


