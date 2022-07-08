package tests

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestDataBaseConn(t *testing.T) {
	dataSource := testConfig.Database.DSN
	_, err := gorm.Open("mysql", dataSource)
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
