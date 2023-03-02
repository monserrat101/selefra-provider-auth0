# Table: auth0_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enabled | bool | X | √ | Indicates whether the rule is enabled. | 
| id | string | X | √ |  | 
| name | string | X | √ | Name of the rule. May only contain alphanumeric characters, spaces, and hyphens. May neither start nor end with hyphens or spaces. | 
| order | float | X | √ | Order in which the rule executes relative to other rules. Lower-valued rules execute first. | 
| script | string | X | √ | Code to be executed when the rule runs. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


