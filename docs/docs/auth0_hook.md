# Table: auth0_hook

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| script | string | X | √ | Code to be executed when this hook runs. | 
| secrets | json | X | √ | The secrets associated with the hook. | 
| trigger_id | string | X | √ | Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message. | 
| dependencies | json | X | √ | Dependencies of this hook used by the WebTask server. | 
| enabled | bool | X | √ | Whether the hook is enabled, or disabled. | 
| id | string | X | √ |  | 
| name | string | X | √ | Name of this hook. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


