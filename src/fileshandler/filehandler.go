//Package fileshandler to handle files in cleaner
package fileshandler

import (
	"fmt"
	"reflect"
)

type item struct {
	status bool
}

//Discover recurently files in given JSON file
func Discover(m map[string]interface{}) {

	for k, v := range m {
		// mItem, ok := v.(interface{})
		// if ok {
		// 	fmt.Println(mItem)
		// }
		fmt.Println("type:", reflect.Value(v))
		switch vv := v.(type) {
		case bool:
			fmt.Println("its a file", vv)
		case map[string]interface{}:
			Discover(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
