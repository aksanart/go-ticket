package config

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}, str ...interface{}) {
	switch v := data.(type) {
	case []byte:
		var obj map[string]interface{}
		json.Unmarshal(v, &obj)
		data = obj
	}
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Printf("--------------------------------%s--------------------------------\n", str)
	jr, _ := json.MarshalIndent(data, "| ", "   ")
	fmt.Println("|", string(jr))
	fmt.Printf("--------------------------------%s--------------------------------\n", str)
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
}
