package config

type Config struct {
	Version               string    `yaml:"version"`
	StatusHost            string    `yaml:"status_host"`
	ProxyStatusListenPort string    `yaml:"proxy_status_listen_port"`
	HttpListenPort        string    `yaml:"http_listen_port"`
	HttpsListenPort       string    `yaml:"https_listen_port"`
	Logfile               string    `yaml:"log_file"`
	Services              []Service `yaml:"services"`
	SSLEnabled            bool      `yaml:"ssl_enabled"` // SSLEnabled is by default falue
	SSLCrtPem             string    `yaml:"ssl_crt_pem"` // SSLCrtPem is the path to the SSL certificate
	SSLKeyPem             string    `yaml:"ssl_key_pem"` // SSLKeyPem is the path to the SSL key
}

type Service struct {
	Name            string `yaml:"name"`
	Type            string `yaml:"type"` // Type is one of ['web', 'proxy', 'static']
	IngressUrl      string `yaml:"ingress_url"`
	EgressUrl       string `yaml:"egress_url"`
	XFrameOptions   string `yaml:"x_frame_options"`   // XFrameOptions is one of ['DENY', 'SAMEORIGIN', 'ALLOW-FROM']
	HSTSMaxAge      int    `yaml:"hsts_max_age"`      // HSTSMaxAge is the max age in seconds
	BlackListIPFile string `yaml:"blacklist_ip_file"` // BlackListIPFile is the CIDR newline separated txt file, see example
}
