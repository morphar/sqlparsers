package generator

import (
	"fmt"
	"go/format"
	"log"
	"reflect"
	"strconv"

	parser "github.com/morphar/sqlparsers/mysql"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func Parse(stmt parser.Statement) (src string) {
	srcCode := recurse(stmt)

	byteSrc, err := format.Source([]byte(srcCode))
	if err != nil {
		fmt.Println(srcCode)
		log.Fatal(err)
	}

	return string(byteSrc)
}

func recurse(cur interface{}) (src string) {
	curVal := reflect.ValueOf(cur)
	curTyp := reflect.TypeOf(cur)

	ptr := ""
	if curTyp.Kind() == reflect.Ptr {
		ptr = "&"
		curVal = curVal.Elem()
		curTyp = curTyp.Elem()
	}

	src = ptr + curTyp.String()

	if curTyp.Kind() == reflect.Struct && curTyp.String() == "parser.ColIdent" {
		src = `parser.NewColIdent`

	} else if curTyp.Kind() == reflect.Struct && curTyp.String() == "parser.TableIdent" {
		src = `parser.NewTableIdent`

	} else if curTyp.Kind() == reflect.Bool {
		if curVal.Bool() {
			src += "(true)"
		} else {
			src += "(false)"
		}
		return
	}

	if curTyp.Kind() == reflect.Array || curTyp.Kind() == reflect.Slice {
		src += "{"
		for i := 0; i < curVal.Len(); i++ {
			src += "\n" + recurse(curVal.Index(i).Interface()) + ",\n"
		}
		src += "}"

	} else if curTyp.Kind() == reflect.String {
		src += `("` + curVal.String() + `")`

	} else if curTyp.Kind() == reflect.Int64 {
		// Leave it alone

	} else if curTyp.Kind() == reflect.Struct && curTyp.String() == "parser.ColIdent" {
		src += `("` + curVal.Field(1).String() + `")`

	} else if curTyp.Kind() == reflect.Struct && curTyp.String() == "parser.TableIdent" {
		src += `("` + curVal.Field(0).String() + `")`

	} else if curTyp.Kind() == reflect.Struct {
		src += "{\n"

		for i := 0; i < curVal.NumField(); i++ {
			valField := curVal.Field(i)
			typField := curTyp.Field(i)

			// Check if this is an unexported field
			if !valField.CanInterface() {
				continue
			}

			src += typField.Name + ": "

			// Handle basic types
			// Probably missing an int and/or float type right now...
			switch val := valField.Interface().(type) {
			case []byte:
				src += `[]byte("` + string(valField.Bytes()) + `")`

			case bool:
				if val {
					src += "true"
				} else {
					src += "false"
				}
			case string:
				src += `"` + valField.String() + `"`
			case int:
				src += strconv.Itoa(int(valField.Int()))
			case int64:
				src += strconv.Itoa(int(valField.Int()))
			default:
				if valField.Kind() == reflect.String {
					src += reflect.TypeOf(valField.Interface()).String() + "(" + `"` + valField.String() + `"` + ")"

				} else if valField.Kind() == reflect.Int {
					src += strconv.Itoa(int(valField.Int()))

				} else if valField.Kind() == reflect.Struct {
					src += recurse(valField.Interface())

				} else if valField.IsNil() {
					// if valField.Kind() == reflect.Ptr {
					src += "nil"
					// } else {
					// 	src += valField.Type().String() + "{}"
					// }

				} else {
					src += recurse(valField.Interface())
				}
			}
			src += ",\n"
		}

		src += "}"

	} else {
		log.Println("UNCAUGHT TYPE:")
		log.Printf("%#v\n", curVal)
		log.Printf("%#v\n", curTyp.Kind().String())
		log.Printf("%#v\n", curTyp.String())
	}

	return
}
