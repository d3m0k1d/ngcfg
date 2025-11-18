# Table of directives

### Http block

| YAML    | Nginx directive  | Required | Type    |
| ----------- | --------------- | ----------- | ------ |
| client_max_body_size        | client_max_body_size     |            | string |
| keepalive_timeout   | keepalive_timeout        |            | int |
| send_timeout        | send_timeout             |            | int |
| gzip        | gzip on\|off             |            | bool |
| sendfile        | sendfile on\|off             |            | bool |
| worker_processes        | worker_processes             |            | int |
| tcp_nopush        | tcp_nopush on\|off             |            | bool |
| access_log        | access_log             |            | string |
| error_log        | error_log             |            | string |
| add_header        | add_header             |            | []string |
| server_tokens        | server_tokens             |            | string |
| limit_req        | limit_req             |            | string |
| limit_req_zone        | limit_req_zone             |            | string |
| limit_conn        | limit_conn             |            | string |
| limit_conn_zone        | limit_conn_zone             |            | string |


### Events block

| YAML    | Nginx directive  | Required | Type    |
| ----------- | --------------- | ----------- | ------ |
| worker_connections        | worker_connections             |            | int |
| multi_accept        | multi_accept             |            | bool |
| use        | use             |            | string |



#### Server block

| YAML    | Nginx directive  | Required | Type    |
| ----------- | --------------- | ----------- | ------ |
| name        | server_name     | ✓           | string |
| listen      | listen          | ✓           | int    |
| listen_v6   | listen [::1]    |             | int    |
| return      | return          |             | string |
| charset     | charset         |             | string |
| root_path_s | root            |             | string |
| ssl         | listen 443 ssl     |             | bool   |
| ssl_cert    | ssl_certificate |             | string |
| ssl_key     | ssl_certificate_key |             | string |
| ssl_proto   | ssl_protocols   |             | string |
| ssl_buffer_size  | ssl_buffer_size |             | string    |

#### Location block

| YAML    | Nginx directive | Required | Type    |
| ----------- | --------------- | ----------- | ------ |
| name        | location {name} | ✓           | string |
| root_path   | root            |             | string |
| alias_path  | alias           |             | string |
| proxy_pass  | proxy_pass      |             | string |
| proxy_buffer_size  | proxy_buffer_size |             | string |
| proxy_set_header  | proxy_set_header |             | []string |

