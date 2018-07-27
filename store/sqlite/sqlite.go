package sqlite

import (
	"github.com/jinzhu/gorm"
)

// baseSqliteStore shared DB data
type baseSqliteStore struct {
	Db   *gorm.DB
	Name string
}
