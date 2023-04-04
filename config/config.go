package config

type Config struct {
	Database struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"db_name"`
		PGDriver string `yaml:"pg_driver"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`
}
