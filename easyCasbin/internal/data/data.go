package data

import (
	"database/sql"
	"easyCasbin/internal/conf"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo, NewDbIniterRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

func createDataBase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// EnsureDB 确保数据库已存在，如果不存在就创建
func EnsureDB(c *conf.Data) error {
	// 如果配置文件中没有填写数据库名，不执行初始化动作，返回错误
	if c.Database.DbName == "" {
		return errors.New("need config system's database name")
	}
	initDbSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Database.DbName)

	dsn := strings.Split(c.Database.Source, "/")[0] + "/"
	err := createDataBase(dsn, c.Database.Driver, initDbSQL)
	if err != nil {
		return fmt.Errorf("init database failed: %w", err)
	}
	return nil
}

func NewDB(c *conf.Data) *gorm.DB {
	// 终端打印输入 sql 执行记录
	//newLonewLoggergger := logger.New(
	//	slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢查询 SQL 阈值
	//		Colorful:      true,        // 禁用彩色打印
	//		//IgnoreRecordNotFoundError: false,
	//		LogLevel: logger.Info, // Log lever
	//	},
	//)
	err := EnsureDB(c)
	if err != nil {
		log.Errorf("ensure database exist, failed to init database: %v", err)
		panic("failed to init database")
	}

	log.Info("failed opening connection to ")
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		//Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}
