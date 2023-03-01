# Table: auth0_action

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the action. | 
| version_id | string | X | √ | Version ID of the action. This value is available if `deploy` is set to true. | 
| dependencies | json | X | √ | List of third party npm modules, and their versions, that this action depends on. | 
| secrets | json | X | √ | List of secrets that are included in an action or a version of an action. | 
| supported_triggers | json | X | √ | List of triggers that this action supports. At this time, an action can only target a single trigger at a time. Read [Retrieving the set of triggers available within actions](https://registry.terraform.io/providers/auth0/auth0/latest/docs/guides/action_triggers) to retrieve the latest trigger versions supported. | 
| deploy | bool | X | √ | Deploying an action will create a new immutable version of the action. If the action is currently bound to a trigger, then the system will begin executing the newly deployed version of the action immediately. | 
| id | string | X | √ |  | 
| code | string | X | √ | The source code of the action. | 
| runtime | string | X | √ | The Node runtime. Defaults to `node12`. Possible values are: `node12`, `node16` or `node18`. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


