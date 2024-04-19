package config

type Node struct {
	Nid       string `json:"nid" yaml:"nid"`
	Nodename  string `json:"nodename" yaml:"nodename"`
	Country   string `json:"country" yaml:"country"`
	Bandwidth string `json:"bandwidth" yaml:"bandwidth"`
	Tel       string `json:"tel" yaml:"tel"`
	Addr      string `json:"addr" yaml:"addr"`
	Token     string `json:"token" yaml:"token"`
	Ip        string `json:"ip" yaml:"ip"`
	Link      string `json:"link" yaml:"link"`
	Nodetype  string `json:"nodetype" yaml:"nodetype"`
}
