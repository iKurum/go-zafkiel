package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	projectName := "config"
	log.Println("Init viper")
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName(projectName)

	viper.AddConfigPath("./src/zafkiel/")

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

//GetMysqlConnectingString func
func GetMysqlConnectingString() string {
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", user, pwd, host, db, charset)
}

//GetSMTPConfig func
func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

//GetServerURL func
func GetServerURL() (url string) {
	url = viper.GetString("server.url")
	return
}
