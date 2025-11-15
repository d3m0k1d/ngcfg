package cmd

const FullTemplate = `
events {
    {{if .Events.Worker_connections}}worker_connections {{.Events.Worker_connections}};{{end}}
    {{if .Events.Multi_accept}}multi_accept on;{{end}}
    {{if .Events.Use}}use {{.Events.Use}};{{end}}
}

http {
    # Core settings
    {{if .ClientMaxBodySize}}client_max_body_size {{.ClientMaxBodySize}};{{end}}
    {{if .KeepaliveTimeout}}keepalive_timeout {{.KeepaliveTimeout}};{{end}}
    {{if .SendTimeout}}send_timeout {{.SendTimeout}};{{end}}
    
    # Compression
    {{if .Gzip}}gzip on;{{else}}gzip off;{{end}}
    
    # Performance tuning
    {{if .Sendfile}}sendfile on;{{end}}
    {{if .Worker_processes}}worker_processes {{.Worker_processes}};{{end}}
    {{if .Tcp_nopush}}tcp_nopush on;{{end}}
    
    # Logging
    {{if .Access_log}}access_log {{.Access_log}};{{end}}
    {{if .Error_log}}error_log {{.Error_log}};{{end}}
    
    # Security headers
    {{if .Add_header}}add_header {{range .Add_header}}{{.}};{{end}}{{end}}
    {{if .Server_tokens}}server_tokens {{.Server_tokens}};{{end}}

    {{range .Servers}}
    server {
        server_name {{.Name}};
        listen {{.Listen}};
        {{if .ListenV6}}listen [::]:{{.ListenV6}};{{end}}
        {{if .Charset}}charset {{.Charset}};{{end}}
        {{if .Root_path_s}}root {{.Root_path_s}};{{end}}
        {{if .Index}}index {{.Index}};{{end}}
        {{if .Return}}return {{.Return}};{{end}}
        
        # SSL block
        {{if .SSL}}listen 443 ssl;{{end}}
        {{if .SSL_Cert}}ssl_certificate {{.SSL_Cert}};{{end}}
        {{if .SSL_key}}ssl_certificate_key {{.SSL_key}};{{end}}
        {{if .SSL_proto}}ssl_protocols {{join .SSL_proto " "}};{{end}}
    
        {{range .Locations}}
        location /{{.Name}} {
            {{if .Alias_path}}alias {{.Alias_path}};{{else if .Root_path}}root {{.Root_path}};{{end}}
            {{if .Proxy_pass}}proxy_pass {{.Proxy_pass}};{{end}}
            {{if .Proxy_buffer_size}}proxy_buffer_size {{.Proxy_buffer_size}};{{end}}
            {{range .Proxy_set_header}}proxy_set_header {{.}};{{end}}
        }
        {{end}}
    }
    {{end}}
}
`
