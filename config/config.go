package config

import (
	"sync"

	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
)

//Config - struct for holding ENVS
type Config struct {
	MysqlURL      string `envconfig:"MYSQL_URL" required:"true"`
	WorkerRunMode string `envconfig:"WORKER_RUN_MODE" default:"dev"`
}

var appConfig *Config
var once sync.Once
var mu sync.Mutex

//Initialize - the app config
func Initialize(cfg *Config) {
	if cfg != nil {
		runOnce := func() {
			appConfig = cfg
		}
		once.Do(runOnce)
	}
}

//Get - get the app config
func Get() *Config {
	mu.Lock()
	defer mu.Unlock()

	return appConfig
}

//Load - loads the ENVS to config struct
func Load(envFile string) (*Config, error) {
	cfg := &Config{}
	godotenv.Load(envFile)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
