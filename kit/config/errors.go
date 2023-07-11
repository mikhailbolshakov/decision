package config

import "github.com/mikhailbolshakov/decision/kit"

const (
	ErrCodeEnvRootPathNotSet             = "CFG-001"
	ErrCodeEnvFileOpening                = "CFG-002"
	ErrCodeEnvNewConfig                  = "CFG-003"
	ErrCodeEnvLoad                       = "CFG-004"
	ErrCodeServiceConfigNotSpecified     = "CFG-005"
	ErrCodeConfigPortEnvEmpty            = "CFG-006"
	ErrCodeConfigHostEnvEmpty            = "CFG-007"
	ErrCodeConfigTimeout                 = "CFG-008"
	ErrCodeEnvFileLoadingPath            = "CFG-009"
	ErrCodeConfigFileNotFound            = "CFG-010"
	ErrCodeConfigFileOpen                = "CFG-011"
	ErrCodeConfigPathEmpty               = "CFG-012"
	ErrCodeEnvFileNotFound               = "CFG-013"
	ErrCodeEnvFileOpen                   = "CFG-014"
	ErrCodeConfigInit                    = "CFG-015"
	ErrCodeConfigLoad                    = "CFG-016"
	ErrCodeConfigTargetObjectInvalidType = "CFG-017"
)

var (
	ErrEnvRootPathNotSet = func(v string) error {
		return kit.NewAppErrBuilder(ErrCodeEnvRootPathNotSet, "root path env variable %s isn't set", v).Err()
	}
	ErrEnvFileOpening = func(cause error, path string) error {
		return kit.NewAppErrBuilder(ErrCodeEnvFileOpening, "").Wrap(cause).F(kit.KV{"path": path}).Err()
	}
	ErrEnvNewConfig = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeEnvNewConfig, "").Wrap(cause).Err()
	}
	ErrEnvLoad = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeEnvLoad, "").Wrap(cause).Err()
	}
	ErrEnvFileLoadingPath = func(cause error, file string) error {
		return kit.NewAppErrBuilder(ErrCodeEnvFileLoadingPath, ".env loading").Wrap(cause).F(kit.KV{"path": file}).Err()
	}
	ErrServiceConfigNotSpecified = func(svc string) error {
		return kit.NewAppErrBuilder(ErrCodeServiceConfigNotSpecified, "config for service isn't specified").F(kit.KV{"svc": svc}).Err()
	}
	ErrConfigPortEnvEmpty = func() error {
		return kit.NewAppErrBuilder(ErrCodeConfigPortEnvEmpty, "env var CONFIG_CFG_GRPC_PORT is empty").Err()
	}
	ErrConfigHostEnvEmpty = func() error {
		return kit.NewAppErrBuilder(ErrCodeConfigHostEnvEmpty, "env var CONFIG_CFG_GRPC_HOST is empty").Err()
	}
	ErrConfigTimeout = func() error {
		return kit.NewAppErrBuilder(ErrCodeConfigTimeout, "not ready within timeout").Err()
	}
	ErrConfigFileNotFound = func(path string) error {
		return kit.NewAppErrBuilder(ErrCodeConfigFileNotFound, "config file %s not found", path).Err()
	}
	ErrEnvFileNotFound = func(path string) error {
		return kit.NewAppErrBuilder(ErrCodeEnvFileNotFound, "env file %s not found", path).Err()
	}
	ErrConfigFileOpen = func(cause error, path string) error {
		return kit.NewAppErrBuilder(ErrCodeConfigFileOpen, "open file %s", path).Wrap(cause).Err()
	}
	ErrEnvFileOpen = func(cause error, path string) error {
		return kit.NewAppErrBuilder(ErrCodeEnvFileOpen, "open file %s", path).Wrap(cause).Err()
	}
	ErrConfigPathEmpty = func() error {
		return kit.NewAppErrBuilder(ErrCodeConfigPathEmpty, "config path empty").Err()
	}
	ErrConfigTargetObjectInvalidType = func() error {
		return kit.NewAppErrBuilder(ErrCodeConfigTargetObjectInvalidType, "target object must be pointer on struct").Err()
	}
	ErrConfigInit = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeConfigInit, "").Wrap(cause).Err()
	}
	ErrConfigLoad = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeConfigLoad, "").Wrap(cause).Err()
	}
)
