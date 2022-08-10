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

	// generate query using for db
	g.ApplyBasic(
		model.User{},
		model.Ordertab{},
		model.Category{},
		model.OrderDetail{},
		model.Product{},
	)

	g.Execute()
}
