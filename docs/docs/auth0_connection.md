# Table: auth0_connection

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enabled_clients | json | X | √ | IDs of the clients for which the connection is enabled. | 
| metadata | json | X | √ | Metadata associated with the connection, in the form of a map of string values (max 255 chars). Maximum of 10 metadata properties allowed. | 
| name | string | X | √ | Name of the connection. | 
| show_as_button | bool | X | √ | Display connection as a button. Only available on enterprise connections. | 
| display_name | string | X | √ | Name used in login screen. | 
| id | string | X | √ |  | 
| is_domain_connection | bool | X | √ | Indicates whether the connection is domain level. | 
| realms | json | X | √ | Defines the realms for which the connection will be used (e.g., email domains). If not specified, the connection name is added as the realm. | 
| strategy | string | X | √ | Type of the connection, which indicates the identity provider. | 
| options | json | X | √ | Configuration settings for connection options. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


