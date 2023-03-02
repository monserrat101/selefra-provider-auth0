# Table: auth0_email

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_from_address | string | X | √ | Email address to use as the sender when no other "from" address is specified. | 
| enabled | bool | X | √ | Indicates whether the email provider is enabled. | 
| id | string | X | √ |  | 
| name | string | X | √ | Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`. | 
| credentials | json | X | √ | Configuration settings for the credentials for the email provider. | 
| settings | json | X | √ | Specific email provider settings. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


