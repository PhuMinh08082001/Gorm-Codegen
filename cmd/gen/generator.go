package main

import (
	"gorm.io/gen"
	"gorm/dal"
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

	g.ApplyBasic(g.GenerateModel("users"))

	g.Execute()
}
