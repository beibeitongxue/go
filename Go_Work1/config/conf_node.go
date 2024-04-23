package config

type Node struct {
	Nid       string `json:"nid" yaml:"nid"`
	NodeName  string `json:"node_name" yaml:"node_name"`
	Country   string `json:"country" yaml:"country"`
	Bandwidth string `json:"bandwidth" yaml:"bandwidth"`
	Tel       string `json:"tel" yaml:"tel"`
	Addr      string `json:"addr" yaml:"addr"`
	Token     string `json:"token" yaml:"token"`
	Ip        string `json:"ip" yaml:"ip"`
	Link      string `json:"link" yaml:"link"`
	NodeType  string `json:"node_type" yaml:"node_type"`
}
