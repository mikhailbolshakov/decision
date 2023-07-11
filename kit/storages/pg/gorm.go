package pg

import (
	"gorm.io/gorm"
	"time"
)

// GormDto specifies base attrs for GORM dto
type GormDto struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt
}

func (g *GormDto) GetLastModified() *time.Time {
	if g == nil {
		return nil
	}
	if g.DeletedAt != nil && g.DeletedAt.Valid {
		return &g.DeletedAt.Time
	}
	if g.UpdatedAt != nil {
		return g.UpdatedAt
	}
	return g.CreatedAt
}

// StringToNull transforms empty string to nil string, so that gorm stores it as NULL
func StringToNull(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// NullToString transforms NULL to empty string
func NullToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
