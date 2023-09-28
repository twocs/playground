package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MyTable struct {
	gorm.Model
}

func GetSchema(table any) *schema.Schema {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(table)
	return stmt.Schema
}

func main() {
	tableSchema := GetSchema(MyTable{})
	fmt.Println(tableSchema.Table)
}
