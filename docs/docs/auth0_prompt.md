# Table: auth0_prompt

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| webauthn_platform_first_factor | bool | X | √ | Determines if the login screen uses identifier and biometrics first. | 
| id | string | X | √ |  | 
| identifier_first | bool | X | √ | Indicates whether the identifier first is used when using the new Universal Login experience. | 
| universal_login_experience | string | X | √ | Which login experience to use. Options include `classic` and `new`. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


