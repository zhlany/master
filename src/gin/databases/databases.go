package databases

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"zhl/src/gin/app/model"
	"zhl/src/gin/errors"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type MysqlConfig struct {
	Mysql ConfigDetail `yaml:"mysql"`
}

type ConfigDetail struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
	Url      string `yaml:"-"`
}

// InitDB 初始化数据库连接
func InitDB(configs *MysqlConfig) error {
	config := configs.Mysql
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
	config.Url = url
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return errors.NewError("DB err :", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//启用自动迁移模式 --可以保持mysql表更新到最新
	db.AutoMigrate(&model.Student{}, &model.Uptime{})

	DB = db

	return nil
}

// Load 解析yaml文件
func Load(path string) (*MysqlConfig, error) {
	conf := new(MysqlConfig)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("yamlFile.Get err:", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal("Unmarshal:", err)
		return nil, err
	}
	return conf, err
}
