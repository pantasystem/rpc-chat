package main

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"com.github/Kinoshita0623/rpc-chat/app/src/repository/impl"
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"

)

func main() {
	fmt.Printf("Hello go")
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "test", "password", "db:3306", "database")

	db, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Post{})

	/*rf := impl.Repository{
		DB: db,
	}
	rf.UserRepository().Find(int64(1))
	*/
}