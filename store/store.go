package store

import (
	"github.com/dkeng/cogo/entity"
	"github.com/dkeng/pkg/logger"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

// Store 存储
type Store struct {
	DB *gorm.DB
}

// Open 打开存储
func (s *Store) Open() (err error) {
	// 初始化数据库
	db, err := gorm.Open("mysql", viper.GetString("mysql.address"))
	if err != nil {
		logger.Fatalf(
			"初始化 MySQL 连接失败: %s \n",
			errors.Wrap(err, "打开 MySQL 连接失败"),
		)
		return err
	}
	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf(
			"初始化 MySQL 连接失败: %s \n",
			errors.Wrap(err, "Ping MySQL 失败"),
		)
		return err
	}

	db.LogMode(viper.GetBool("mysql.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
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

// AllStore 所有存储
type AllStore struct {
	ApplicationStore ApplicationStore
	ConfigStore      ConfigStore
}

// baseStore shared DB data
type baseStore struct {
	Db   *gorm.DB
	Name string
}

// Init 初始化
func (a *AllStore) Init(s *Store) *AllStore {
	a.ApplicationStore = new(ApplicationStore).Init(s.DB)
	a.ConfigStore = new(ConfigStore).Init(s.DB)
	return a
}
