package dbservice

import (
	"testing"
	"github.com/spf13/viper"
    "github.com/sirupsen/logrus"
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
)

var db *Postgres

func initConfig() error {
    viper.AddConfigPath("../../config")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}

func setup() {
	if err := initConfig(); err != nil {
		logrus.Fatalf(err.Error())
	}

	config := PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname") + "_test",
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "",
	}

	var err error
	db, err = NewPostgres(config)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	if err := db.Migrate(); err != nil {
		logrus.Fatalf(err.Error())
	}
}

func teardown() {
	db.db.Where("id>=0").Delete(&m.User{})
	db.db.Where("id>=0").Delete(&m.UserLog{})
	db.db.Where("id>=0").Delete(&m.Admin{})
	db.db.Where("id>=0").Delete(&m.ErrorLog{})
}

func TestMain(m *testing.M) {
    setup()
    m.Run() 
	teardown()
}

func TestCreateAdmin(t *testing.T) {
    db.CreateAdmin(10, "default_state")
	exists, _ := db.CheckAdminExists(10)
	
	if !exists {
		t.Errorf("admin does not exists")
	}
}

func TestGetAdminState(t *testing.T) {
    db.CreateAdmin(20, "default_state")
	state, _ := db.GetAdminState(20)
	
	if state != "default_state" {
		t.Errorf("invalid state")
	}
}

func TestSetAdminState(t *testing.T) {
    db.CreateAdmin(30, "default_state")
	db.SetAdminState(30, "new_state")
	state, _ := db.GetAdminState(30)
	
	if state != "new_state" {
		t.Errorf("invalid state")
	}
}

func TestCreateUser(t *testing.T) {
    db.CreateUser(10, 0, "default_state")
	_, err := db.getUser(10)
	
	if err != nil {
		t.Errorf("admin does not exists")
	}
}

func TestGetUserState(t *testing.T) {
	db.CreateUser(20, 1, "default_state")
	state, _ := db.GetUserState(20)
	
	if state != "default_state" {
		t.Errorf("invalid state")
	}
}

func TestSetUserState(t *testing.T) {
	db.CreateUser(30, 2, "default_state")
	db.SetUserState(30, "new_state")
	state, _ := db.GetUserState(30)
	
	if state != "new_state" {
		t.Errorf("invalid state")
	}
}

func TestCreateUserLog(t *testing.T) {
	db.CreateUser(10, 0, "default_state")
	db.CreateUserLog(10, "ip", "info")
	_, err := db.getUniqueUserLogs(10)

	if err != nil {
		t.Errorf("log not found")
	}
}