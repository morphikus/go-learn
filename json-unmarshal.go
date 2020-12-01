package main

import (
	"encoding/json"
	"fmt"
)

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Толик","Last":"К.","age":39},{"First":"Галина","Last":"В.","Age":48}]`

	bs := []byte(s)

	fmt.Println(bs)

	var people []person2

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
	}

}
