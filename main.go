package main

type IncrementInput struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type IncrementResult struct {
	Value int `json:"value"`
}

var IncrementMap = make(map[string]int)

func IncrementingStructFunc(inp IncrementInput) IncrementResult {
	return 0
}

func main() {

}
