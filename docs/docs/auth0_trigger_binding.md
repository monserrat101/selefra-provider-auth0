# Table: auth0_trigger_binding

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| trigger | string | X | √ | The ID of the trigger to bind with. | 
| actions | json | X | √ | The actions bound to this trigger | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


