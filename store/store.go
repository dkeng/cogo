package store

import (
	"github.com/dkeng/cogo/entity"
	"github.com/dkeng/cogo/store/sqlite"
	"github.com/dkeng/pkg/logger"
	"github.com/spf13/viper"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Store 存储
type Store struct {
	DB *gorm.DB
}

// Open 打开存储
func (s *Store) Open() (err error) {
	// 初始化数据库
	db, err := gorm.Open("sqlite3", viper.GetString("sqlite.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 Sqlite 连接失败: %s \n",
			errors.Wrap(err, "打开 Sqlite 连接失败"),
		)
	}

	if db.DB().Ping() != nil {
		logger.Fatalf(
			"初始化 Sqlite 连接失败: %s \n",
			errors.Wrap(err, "Ping Sqlite 失败"),
		)
	}

	db.LogMode(viper.GetBool("sqlite.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("sqlite.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("sqlite.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(
		&entity.Application{},
		&entity.Config{},
	)
	s.DB = db
	return
}

// Close 关闭存储
func (s *Store) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

// AllStore mysql存储
type AllStore struct {
	AllSqliteStore *sqlite.AllSqliteStore
}

// Init 初始化
func (m *AllStore) Init(s *Store) {
	m.AllSqliteStore = new(sqlite.AllSqliteStore).Init(s.DB)
}
