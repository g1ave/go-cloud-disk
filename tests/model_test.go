package tests

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/rest"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

type Config struct {
	rest.RestConf `yaml:",inline"`

	Database struct {
		DSN string `yaml:"dsn"`
	}
}

var configFile = flag.String("f", "../core/etc/core-api.yaml", "the config file")

func TestDataBaseConn(t *testing.T) {
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		t.Errorf("read file error: %v", err)
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		t.Errorf("yaml unmarshal error: %v", err)
		return
	}
	dataSource := config.Database.DSN
	_, err = gorm.Open("mysql", dataSource)
	if err != nil {
		t.Errorf("connect DB error: %v, %v", err, dataSource)
		return
	}
}

func TestReadYamlFile(t *testing.T) {
	type Database struct {
		DSN string
	}
	db := &Database{}
	err := yaml.Unmarshal([]byte("dsn: test"), db)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	fmt.Println(db.DSN)
}
