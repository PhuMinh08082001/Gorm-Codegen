package main

import (
	"fmt"
	"gorm/dal"
	"gorm/utils"
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
