package main

import (
	"gorm.io/gen"
	"gorm/dal"
	"gorm/dal/model"
)

func init() {
	dal.DB = dal.ConnectDB().Debug()
}

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})

	g.UseDB(dal.DB)

	// generate all table from database
	//g.ApplyBasic(g.GenerateAllTable()...)
	//g.GenerateModel("users")
	//g.GenerateModel("ordertabs")
	//g.ApplyBasic(g.GenerateModel("users"))
	g.ApplyBasic(model.User{}, model.Ordertab{})

	g.Execute()
}
