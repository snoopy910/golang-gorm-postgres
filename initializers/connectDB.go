package initializers

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB(config *Config) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open("postgres", dsn)


	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
	return DB
}

// dsn := url.URL{
// 	User:     url.UserPassword(config.DBUserName, config.DBUserPassword),
// 	Scheme:   config.DBName,
// 	Host:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
// 	Path:     conf.DBName,
// 	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
// }
// db, err := gorm.Open("postgres", dsn.String())
