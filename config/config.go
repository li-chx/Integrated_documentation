package config

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

// 连接数据库的相关信息
type dbConfig struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Hostname  string `yaml:"hostname"`
	Port      string `yaml:"port"`
	Dbname    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Local     string `yaml:"local"`
}

type mailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	FromName string `yaml:"fromName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type config struct {
	DbConfig   dbConfig   `yaml:"dbConfig"`
	MailConfig mailConfig `yaml:"mailConfig"`
}

var Config config

func init() {
	yamlFile, err := os.ReadFile("./conf/Config.yaml")
	if err != nil {
		log.Errorf("read file error: %+v", errors.WithStack(err))
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Errorf("unmarshal error: %+v", errors.WithStack(err))
	}
	if dbUsername := os.Getenv("db_username"); dbUsername != "" {
		Config.DbConfig.Username = dbUsername
	}
	if dbPas := os.Getenv("db_password"); dbPas != "" {
		Config.DbConfig.Password = dbPas
	}
	if dbHostname := os.Getenv("db_hostname"); dbHostname != "" {
		Config.DbConfig.Hostname = dbHostname
	}
	if dbPort := os.Getenv("db_port"); dbPort != "" {
		Config.DbConfig.Port = dbPort
	}
}

func GetDbConfig() dbConfig {
	return Config.DbConfig
}

func GetMailConfig() mailConfig {
	return Config.MailConfig
}
