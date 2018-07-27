package store

import (
	"github.com/dkeng/cogo/entity"
	"github.com/dkeng/pkg/logger"
	"github.com/spf13/viper"

	csqlite "github.com/dkeng/cogo/store/sqlite"
	"github.com/jinzhu/gorm"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

// Store 存储
type Store struct {
	DB *gorm.DB
}

// Open 打开存储
func (s *Store) Open() (err error) {
	switch viper.GetString("db.dialect") {
	case "sqlite":
		return s.initSqlite3()
	}
	return nil
}

func (s *Store) initSqlite3() error {
	// 初始化数据库
	db, err := gorm.Open("sqlite3", viper.GetString("db.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 Sqlite 连接失败: %s \n",
			errors.Wrap(err, "打开 Sqlite 连接失败"),
		)
		return err
	}
	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf(
			"初始化 Sqlite 连接失败: %s \n",
			errors.Wrap(err, "Ping Sqlite 失败"),
		)
		return err
	}

	db.LogMode(viper.GetBool("db.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("db.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("db.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(
		&entity.Application{},
		&entity.Config{},
	)
	s.DB = db
	return nil
}

// Close 关闭存储
func (s *Store) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

// AllStore mysql存储
type AllStore struct {
	ApplicationStore ApplicationStore
	ConfigStore      ConfigStore
}

func newSqlite(s *Store) *AllStore {
	return &AllStore{
		ApplicationStore: new(csqlite.ApplicationStore).Init(s.DB),
		ConfigStore:      new(csqlite.ConfigStore).Init(s.DB),
	}
}

// Init 初始化
func (m *AllStore) Init(s *Store) *AllStore {
	switch viper.GetString("db.dialect") {
	case "sqlite":
		return newSqlite(s)
	}
	return nil
}
