package config

type SiteInfo struct {
	Webside  string `yaml:"web_side" json:"web_side"`
	Email    string `yaml:"email" json:"email"`
	Telegram string `yaml:"telegram" json:"telegram"`
	Version  string `yaml:"version" json:"version"`
	Platform string `yaml:"platform" json:"platform"`
	Newest   string `yaml:"newest" json:"newest"`
	Desc     string `yaml:"desc" json:"desc"`
	Minimum  string `yaml:"minimum" json:"minimum"`
}
