package config

type Config struct {
	LskyServer    string `yaml:"lsky-server"`
	LskyAuthToken string `yaml:"lsky-auth-token"`
}
