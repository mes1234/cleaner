//Package fileshandler to handle files in cleaner
package fileshandler

import "fmt"

//Discover recurently files in given JSON file
func Discover(m map[string]interface{}) {

	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			fmt.Println(k, "is bool", vv)
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case map[string]interface{}:
			Discover(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
