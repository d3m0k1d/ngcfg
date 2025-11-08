package cmd

const ngconfTemplate = `server {
	server_name {{.Name}};
	listen {{.Listen}};
	{{if .SSL}}ssl on;{{end}}
	
	{{range .Locations}}
	location /{{.Name}} {
		alias {{.Path}};
	}
	{{end}}
}`
