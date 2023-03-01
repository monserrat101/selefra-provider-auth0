# Table: auth0_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| nickname | string | X | √ | Preferred nickname or alias of the user. This value can only be updated if the connection is a database connection (using the Auth0 store), a passwordless connection (email or sms) or has disabled 'Sync user profile attributes at each login'. For more information, see: [Configure Identity Provider Connection for User Profile Updates](https://auth0.com/docs/manage-users/user-accounts/user-profiles/configure-connection-sync-with-auth0). | 
| phone_verified | bool | X | √ | Indicates whether the phone number has been verified. | 
| picture | string | X | √ | Picture of the user. This value can only be updated if the connection is a database connection (using the Auth0 store), a passwordless connection (email or sms) or has disabled 'Sync user profile attributes at each login'. For more information, see: [Configure Identity Provider Connection for User Profile Updates](https://auth0.com/docs/manage-users/user-accounts/user-profiles/configure-connection-sync-with-auth0). | 
| app_metadata | string | X | √ | Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts. | 
| blocked | bool | X | √ | Indicates whether the user is blocked or not. | 
| given_name | string | X | √ | Given name of the user. This value can only be updated if the connection is a database connection (using the Auth0 store), a passwordless connection (email or sms) or has disabled 'Sync user profile attributes at each login'. For more information, see: [Configure Identity Provider Connection for User Profile Updates](https://auth0.com/docs/manage-users/user-accounts/user-profiles/configure-connection-sync-with-auth0). | 
| id | string | X | √ |  | 
| username | string | X | √ | Username of the user. Only valid if the connection requires a username. | 
| email | string | X | √ | Email address of the user. | 
| email_verified | bool | X | √ | Indicates whether the email address has been verified. | 
| roles | json | X | √ | Set of IDs of roles assigned to the user. | 
| phone_number | string | X | √ | Phone number for the user; follows the E.164 recommendation. Used for SMS connections.  | 
| user_id | string | X | √ | ID of the user. | 
| user_metadata | string | X | √ | Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences. | 
| verify_email | bool | X | √ | Indicates whether the user will receive a verification email after creation. Overrides behavior of `email_verified` parameter. | 
| connection_name | string | X | √ | Name of the connection from which the user information was sourced. | 
| family_name | string | X | √ | Family name of the user. This value can only be updated if the connection is a database connection (using the Auth0 store), a passwordless connection (email or sms) or has disabled 'Sync user profile attributes at each login'. For more information, see: [Configure Identity Provider Connection for User Profile Updates](https://auth0.com/docs/manage-users/user-accounts/user-profiles/configure-connection-sync-with-auth0). | 
| name | string | X | √ | Name of the user. This value can only be updated if the connection is a database connection (using the Auth0 store), a passwordless connection (email or sms) or has disabled 'Sync user profile attributes at each login'. For more information, see: [Configure Identity Provider Connection for User Profile Updates](https://auth0.com/docs/manage-users/user-accounts/user-profiles/configure-connection-sync-with-auth0). | 
| password | string | X | √ | Initial password for this user. Required for non-passwordless connections (SMS and email). | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


