package internal

type HttpConfig struct {
	Http Http `yaml:"http"`
}

type Http struct {
	Events  Events   `yaml:"events"`
	Servers []Server `yaml:"servers"`

	ClientMaxBodySize string `yaml:"client_max_body_size"`
	KeepaliveTimeout  int    `yaml:"keepalive_timeout"`
	SendTimeout       int    `yaml:"send_timeout"`
	Gzip              bool   `yaml:"gzip"`

	Sendfile         bool   `yaml:"sendfile"`
	Worker_processes string `yaml:"worker_processes"`
	Tcp_nopush       bool   `yaml:"tcp_nopush"`

	Access_log string `yaml:"access_log"`
	Error_log  string `yaml:"error_log"`

	Add_header      []string `yaml:"add_header"`
	Server_tokens   bool     `yaml:"server_tokens"`
	Limit_req       string   `yaml:"limit_req"`
	Limit_req_zone  string   `yaml:"limit_req_zone"`
	Limit_conn_zone string   `yaml:"limit_conn_zone"`
	Limit_conn      string   `yaml:"limit_conn"`
}

type Events struct {
	Worker_connections int    `yaml:"worker_connections"`
	Multi_accept       bool   `yaml:"multi_accept"`
	Use                string `yaml:"use"`
}

type Server struct {
	Name        string `yaml:"name" validate:"required"`
	Listen      int    `yaml:"listen" validate:"required,min=1,max=65535"`
	ListenV6    int    `yaml:"listen_v6" validate:"min=1,max=65535"`
	Return      string `yaml:"return"`
	Root_path_s string `yaml:"root_path_s"`
	Charset     string `yaml:"charset"`
	Index       string `yaml:"index"`

	SSL             bool   `yaml:"ssl"`
	SSL_buffer_size string `yaml:"ssl_buffer_size"`

	SSL_Cert  string   `yaml:"ssl_cert"`
	SSL_key   string   `yaml:"ssl_key"`
	SSL_proto []string `yaml:"ssl_protocols"`

	Locations []Location `yaml:"locations"`
}

type Location struct {
	Name              string   `yaml:"name"`
	Root_path         string   `yaml:"root_path"`
	Alias_path        string   `yaml:"alias_path"`
	Proxy_pass        string   `yaml:"proxy_pass"`
	Proxy_buffer_size string   `yaml:"proxy_buffer_size"`
	Proxy_set_header  []string `yaml:"proxy_set_header"`
}
