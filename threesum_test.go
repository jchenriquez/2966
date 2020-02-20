package threesum

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"encoding/json"
	"testing"
)

type Test struct {
	Nums []int `json:nums`
	Result [][]int `json:"result"`
}
 
func TestThreeSum (t *testing.T) {
	f, err := os.Open("./tests.json")
	
	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				t.Run(name, func(st *testing.T) {
					res := ThreeSum(test.Nums)

					if fmt.Sprintf("%v", test.Result) != fmt.Sprintf("%v", res) {
						st.Errorf("returned %v", res)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}
	}
}