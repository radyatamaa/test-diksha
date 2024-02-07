package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sort"
)


func main() {
	req := []string{"gym","school","store"}

	blocks := map[int]interface{}{
		0 : map[string]interface{}{
			"gym":false,
			"school":true,
			"store":false,
		},
		1 : map[string]interface{}{
			"gym":true,
			"school":false,
			"store":false,
		},
		2 : map[string]interface{}{
			"gym":true,
			"school":true,
			"store":false,
		},
		3 : map[string]interface{}{
			"gym":false,
			"school":true,
			"store":false,
		},
		4 : map[string]interface{}{
			"gym":false,
			"school":true,
			"store":true,
		},
	}

	var result []int
	for _,e := range req {
		if blocks[0].(map[string]interface{})[e].(bool) {
			result = append(result,0)
			continue
		}

		if blocks[1].(map[string]interface{})[e].(bool) {
			result = append(result,1)
			continue
		}

		if blocks[2].(map[string]interface{})[e].(bool) {
			result = append(result,2)
			continue
		}

		if blocks[3].(map[string]interface{})[e].(bool) {
			result = append(result,3)
			continue
		}

		if blocks[4].(map[string]interface{})[e].(bool) {
			result = append(result,4)
			continue
		}
	}

	fmt.Println("result array jarak : ", result)

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	if len(result) > 0 {
		fmt.Println("result Jarak terjauh index : " , result[0])
	}
}

