package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm/dal"
	"gorm/utils"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "test1"
)

func init() {
	dal.DB = dal.ConnectDB().Debug()
}

func main() {
	users := utils.GetUserHaveGTOneOrder()
	for _, u := range users {
		fmt.Println(u)
	}
}
