package types

import (
	"github.com/ZaneWithSpoon/fathomBack/db"
	"time"
	"fmt"
)

type User struct {
	UUID  string				`gorm:"type:varchar(355);unique_index;not null"`
	Access_token string			`gorm:"type:varchar(200);unique_index;not null"`
	Created_on  time.Time
	Last_login time.Time
	Upload_count int			`gorm:"type:integer"`
}

func MigrateUsers() {
	fmt.Println("Migrating Users Table")
	err := db.GetDbService().Gorm.Debug().AutoMigrate(&User{}).Error;


	if err != nil {
		fmt.Println("there was an error in User migration")
		fmt.Println(err)
	}
}