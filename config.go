package decision

import (
	"github.com/mikhailbolshakov/decision/kit"
	kitConfig "github.com/mikhailbolshakov/decision/kit/config"
	kitHttp "github.com/mikhailbolshakov/decision/kit/http"
	"github.com/mikhailbolshakov/decision/kit/storages/pg"
	"os"
	"path/filepath"
)

var (
	// Logger service logger
	Logger = kit.InitLogger(&kit.LogConfig{Level: kit.TraceLevel, Format: kit.FormatterJson})
)

func LF() kit.CLoggerFunc {
	return func() kit.CLogger {
		return kit.L(Logger)
	}
}

func L() kit.CLogger {
	return LF()()
}

type CfgStorages struct {
	Database *pg.DbClusterConfig
}

type Config struct {
	Storages *CfgStorages
	Log      *kit.LogConfig
	Http     *kitHttp.Config
}

func LoadConfig() (*Config, error) {

	// get root folder from env
	rootPath := os.Getenv("DECISIONROOT")
	if rootPath == "" {
		return nil, kitConfig.ErrEnvRootPathNotSet("DECISIONROOT")
	}

	// config path
	configPath := filepath.Join(rootPath, "config.yml")

	// .env path
	envPath := filepath.Join(rootPath, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = ""
	}

	// load config
	config := &Config{}
	err := kitConfig.NewConfigLoader(LF()).
		WithConfigPath(configPath).
		WithEnvPath(envPath).
		Load(config)

	if err != nil {
		return nil, err
	}
	return config, nil
}
