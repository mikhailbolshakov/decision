package pg

import "github.com/mikhailbolshakov/decision/kit"

const (
	ErrCodeGooseMigrationUp     = "DB-001"
	ErrCodeGooseMigrationGetVer = "DB-002"
	ErrCodePostgresOpen         = "DB-003"
	ErrCodeGooseFolderNotFound  = "DB-004"
	ErrCodeGooseFolderOpen      = "DB-005"
	ErrCodeGooseMigrationLock   = "DB-006"
	ErrCodeGooseMigrationUnLock = "DB-007"
)

var (
	ErrGooseMigrationUp = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeGooseMigrationUp, "").Wrap(cause).Err()
	}
	ErrGooseMigrationGetVer = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeGooseMigrationGetVer, "").Wrap(cause).Err()
	}
	ErrPostgresOpen = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodePostgresOpen, "").Wrap(cause).Err()
	}
	ErrGooseFolderNotFound = func(path string) error {
		return kit.NewAppErrBuilder(ErrCodeGooseFolderNotFound, "folder not found %s", path).Err()
	}
	ErrGooseFolderOpen = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeGooseFolderOpen, "folder open").Wrap(cause).Err()
	}
	ErrGooseMigrationLock = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeGooseMigrationLock, "locking before migration").Wrap(cause).Err()
	}
	ErrGooseMigrationUnLock = func(cause error) error {
		return kit.NewAppErrBuilder(ErrCodeGooseMigrationUnLock, "unlocking after migration").Wrap(cause).Err()
	}
)
