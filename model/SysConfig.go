package model

type SysConfig struct {
	Port     int    `json:"port"`
	InitRes  string "json:initRes"
	InitPath string "json:initPath"
}
