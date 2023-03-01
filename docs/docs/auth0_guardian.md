# Table: auth0_guardian

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy | string | X | √ | Policy to use. Available options are `never`, `all-applications` and `confidence-score`. | 
| recovery_code | bool | X | √ | Indicates whether recovery code MFA is enabled. | 
| push | json | X | √ | Configuration settings for the Push MFA. If this block is present, Push MFA will be enabled, and disabled otherwise. | 
| webauthn_platform | json | X | √ | Configuration settings for the WebAuthn with FIDO Device Biometrics MFA. If this block is present, WebAuthn with FIDO Device Biometrics MFA will be enabled, and disabled otherwise. | 
| id | string | X | √ |  | 
| otp | bool | X | √ | Indicates whether one time password MFA is enabled. | 
| duo | json | X | √ | Configuration settings for the Duo MFA. If this block is present, Duo MFA will be enabled, and disabled otherwise. | 
| phone | json | X | √ | Configuration settings for the phone MFA. If this block is present, Phone MFA will be enabled, and disabled otherwise. | 
| webauthn_roaming | json | X | √ | Configuration settings for the WebAuthn with FIDO Security Keys MFA. If this block is present, WebAuthn with FIDO Security Keys MFA will be enabled, and disabled otherwise. | 
| email | bool | X | √ | Indicates whether email MFA is enabled. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


