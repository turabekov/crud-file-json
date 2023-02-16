package config

type Config struct {
	Path string

	UserFileName string
	PostFileName string
}

func Load() Config {
	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/users.json"
	cfg.PostFileName = "/posts.json"

	return cfg
}
