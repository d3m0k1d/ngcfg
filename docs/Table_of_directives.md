
#### Server block

| YAML    | Nginx directive  | Required | Type    |
| ----------- | --------------- | ----------- | ------ |
| name        | server_name     | ✓           | string |
| listen      | listen          | ✓           | int    |
| listen_v6   | listen [::1]    |             | int    |
| charset     | charset         |             | string |
| root_path_s | root            |             | string |
| ssl         | ssl on\|off     |             | bool   |
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
