package cmd

const FullTemplate = `
http {
    {{if .ClientMaxBodySize}}client_max_body_size {{.ClientMaxBodySize}};{{end}}
    {{if .KeepaliveTimeout}}keepalive_timeout {{.KeepaliveTimeout}};{{end}}
    {{if .SendTimeout}}send_timeout {{.SendTimeout}};{{end}}

    {{range .Servers}}
    server {
        server_name {{.Name}};
        listen {{.Listen}};
        {{if .ListenV6}}listen [::]:{{.ListenV6}};{{end}}
        {{if .Charset}}charset {{.Charset}};{{end}}
        {{if .Root_path_s}}root {{.Root_path_s}};{{end}}
        {{if .Index}}index {{.Index}};{{end}}
        {{if .Return}}return {{.Return}};{{end}}
        #SSL block
        {{if .SSL}}ssl on;{{end}}
        {{if .SSL_Cert}}ssl_certificate {{.SSL_Cert}};{{end}}
        {{if .SSL_key}}ssl_certificate_key {{.SSL_key}};{{end}}
        {{if .SSL_proto}}ssl_protocols {{join .SSL_proto " "}};{{end}}
        
        {{range .Locations}}
        location /{{.Name}} {
            {{if and .Alias_path (not .Root_path)}}alias {{.Alias_path}};{{end}}
            {{if and .Root_path (not .Alias_path)}}root {{.Root_path}};{{end}}
            {{if and .Alias_path .Root_path}}alias {{.Alias_path}};{{end}}
        }
        {{end}}
    }
    {{end}}
}
`
