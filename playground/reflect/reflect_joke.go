package main

import (
	"cheese/tools"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	StructNameReg = `[A-Z][a-z]*`
)

type AAndB struct {
	id      uint64
	name    string
	age     int
	address string
	// birthday time.Time
}

func main() {
	var a AAndB
	GenEntity(reflect.TypeOf(a))

	// _, err := tools.WriteStrToNewFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func GenEntity(t reflect.Type) error {
	tName := t.Name()
	ne := tools.FindAll(tName, StructNameReg)
	tools.StringTo(ne, strings.ToLower)
	tableName := strings.Join(ne, "_")
	// tableName := strings.Title

	println(tName)
	println(tableName)

	eles := make([]reflect.StructField, 0)
	eleNames := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name != "id" {
			eles = append(eles, f)
			eleNames = append(eleNames, f.Name)
		}
	}
	fmt.Printf("fields: %v\n", eleNames)

	var builder strings.Builder
	symbols := FieldsToSymbol(eles)
	ins := fmt.Sprintf("insert into %s(%s) values (%s)", tableName, strings.Join(eleNames, ","), strings.Join(symbols, ","))
	_, err := builder.WriteString(ins)
	if err != nil {
		log.Fatal(err)
	}
	println(builder.String())

	return nil
}

func FieldsToSymbol(fields []reflect.StructField) []string {
	res := make([]string, 0)
	for _, field := range fields {
		switch field.Type.Name() {
		case "int64", "int32", "int", "uint64", "uint32", "uint8", "byte":
			res = append(res, `"%d"`)
		case "string":
			res = append(res, `"%s"`)
		}
	}

	return res
}
