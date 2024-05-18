package config

type Result struct {
	LskyServer         string `yaml:"lsky-server"`
	LskyAuthToken      string `yaml:"lsky-auth-token"`
	InsecureSkipVerify bool   `yaml:"insecure-skip-verify"`
}
