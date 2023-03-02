# Table: auth0_organization_member

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| organization_id | string | X | √ | The ID of the organization to assign the member to. | 
| roles | json | X | √ | The role ID(s) to assign to the organization member. | 
| user_id | string | X | √ | ID of the user to add as an organization member. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


