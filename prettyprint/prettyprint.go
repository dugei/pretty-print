package prettyprint

import (
	"reflect"
	"strings"
	"time"
	"fmt"
)

const prettyPrintIndentNum  = 2
const prettyPrintIndentChar = " "

func P(values... interface{}){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	strPlus := strings.Repeat("+", 40)
	fmt.Printf("\n%c[0;40;32m%s%s %s %c[0m\n\n", 0x1B, strPlus, currentTime, strPlus, 0x1B)
	for k,s := range values {
		printOne(1, s, false)
		if k < (len(values)-1) {
			fmt.Println(strings.Repeat("-", 80))
		}
	}
	fmt.Printf("\n%c[0;40;32m%s%s %s %c[0m\n\n", 0x1B, strPlus, currentTime, strPlus, 0x1B)
}

func printOne(deep int, s interface{}, needIndent bool){

	currPrefixSpaces := strings.Repeat(prettyPrintIndentChar, deep * prettyPrintIndentNum)
	parentPrefixSpaces := ""
	if len(currPrefixSpaces) != 0 {
		parentPrefixSpaces = strings.Repeat(prettyPrintIndentChar, (deep-1)*prettyPrintIndentNum)
	}

	if s == nil {
		if needIndent == true {
			fmt.Print(currPrefixSpaces)
		}
		fmt.Println(s)
		return
	}
	switch reflect.TypeOf(s).Kind() {
	case reflect.Ptr:
		printOne(deep, reflect.ValueOf(s).Elem().Interface(), false)
	case reflect.Struct:
		value := reflect.ValueOf(s)
		valueType := reflect.TypeOf(s)
		fmt.Println("{")
		for i := 0; i < value.NumField(); i++ {
			fieldName :=  currPrefixSpaces + valueType.Field(i).Name +": "
			fmt.Printf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			if value.Field(i).CanInterface() {
				printOne(deep+1, value.Field(i).Interface(), false)
			}else{
				fmt.Printf("%+v\n", value.Field(i))
			}
		}
		fmt.Println(parentPrefixSpaces + "}")
	case reflect.Array, reflect.Slice:
		value := reflect.ValueOf(s)
		fmt.Println("[")
		for i:=0; i<value.Len();i++{
			printOne(deep + 1, value.Index(i).Interface(), true)
		}
		fmt.Println(parentPrefixSpaces + "]")
	case reflect.Map:
		value := reflect.ValueOf(s)
		if needIndent == true {
			fmt.Print(parentPrefixSpaces)
		}
		fmt.Println("{")
		for _,v := range value.MapKeys() {
			fieldName := currPrefixSpaces + v.String() +": "
			fmt.Printf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			if value.MapIndex(v).CanInterface() {
				printOne(deep + 1, value.MapIndex(v).Interface(), false)
			}else {
				fmt.Println(value.MapIndex(v))
			}
		}
		fmt.Println(parentPrefixSpaces + "}")

	default:
		if needIndent == true {
			fmt.Print(currPrefixSpaces)
		}
		fmt.Println(s)
	}
}
