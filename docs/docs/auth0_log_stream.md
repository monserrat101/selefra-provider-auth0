# Table: auth0_log_stream

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| filters | json | X | √ | Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered. | 
| id | string | X | √ |  | 
| name | string | X | √ | Name of the log stream. | 
| status | string | X | √ | The current status of the log stream. Options are "active", "paused", "suspended". | 
| type | string | X | √ | Type of the log stream, which indicates the sink provider. | 
| sink | json | X | √ | The sink configuration for the log stream. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


