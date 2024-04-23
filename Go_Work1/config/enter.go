package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Email    Email    `yaml:"email"`
	Jwy      Jwy      `yaml:"jwy"`
	Node     Node     `yaml:"node"`
	Redis    Redis    `yaml:"redis"`
}
