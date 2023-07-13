package main

import (
	"encoding/json"
	"fmt"
)

func arrayToNestedMap(array []string, resultMap map[string]interface{}, value interface{}) {
	currentMap := resultMap
	for i := 0; i < len(array)-1; i++ {
		val, ok := currentMap[array[i]]
		if ok {
			currentMap = val.(map[string]interface{})
		} else {
			currentMap[array[i]] = make(map[string]interface{})
			currentMap = currentMap[array[i]].(map[string]interface{})
		}

	}
	currentMap[array[len(array)-1]] = value
}

func main() {
	array := []string{"resource"}
	resultMap := make(map[string]interface{})
	arrayToNestedMap(array, resultMap, "a")
	jsonStr, err := json.Marshal(resultMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
