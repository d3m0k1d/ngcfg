package cmd

const ServerBlockTemplate = `
server {
    server_name {{.Name}};
    listen {{.Listen}};
    return {{.Return}};
    {{if .SSL}}ssl on;{{end}}

    {{range .Locations}}
    location /{{.Name}} {
        {{if and .Alias_path (not .Root_path)}}alias {{.Alias_path}};{{end}}
        {{if and .Root_path (not .Alias_path)}}root {{.Root_path}};{{end}}
        {{if and .Alias_path .Root_path}}alias {{.Alias_path}}{{end}}
    }
    {{end}}
}
`
