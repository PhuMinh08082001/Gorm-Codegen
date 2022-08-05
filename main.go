package main

import (
	"fmt"
	"gorm/dal"
	"gorm/dal/query"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	LastName  string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Ordertabs []Ordertab
}

type Ordertab struct {
	OrderId uint `gorm:"primaryKey"`
	UserId  uint
}

// var user User
//
//	func GetAll(db *gorm.DB) ([]User, error) {
//		var users []User
//		err := db.Model(&User{}).Preload("Ordertabs").Find(&users).Error
//		return users, err
//	}
func init() {
	dal.DB = dal.ConnectDB().Debug()
}

func main() {
	u := query.Use(dal.DB).User
	user, err := u.Where(u.ID.Eq(2)).First()

	if err != nil {
		panic(fmt.Errorf("query fail #{err}"))
	}

	fmt.Println(user)
}
