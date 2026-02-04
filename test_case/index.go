package main

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"
)

func main() {
	var m = map[any]any{
		"one":   "uno",
		"two":   "dos",
		"three": "tres",
	}
	var m_string, err = json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		return
	}
	fmt.Println(gconv.String(m_string))
	var m2 = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}
	json.Unmarshal(m_string, &m2)
	fmt.Println(m2)
}
