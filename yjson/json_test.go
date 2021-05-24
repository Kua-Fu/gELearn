package yjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type employee struct {
	Name interface{} `json:"name"`
	Age  interface{} `json:"age"`
}

func TestJson(t *testing.T) {
	e1 := employee{
		Name: "yz",
		Age:  30,
	}
	j, err := json.Marshal(e1)
	if err != nil {
		fmt.Errorf("json dumps error")
		fmt.Println(err)

	}
	fmt.Println(string(j))
}

type reindex struct {
	Query string `json:"q1"`
	Res   string `json:"res"`
}

func TestStruct(t *testing.T) {

	// e1Json := `{"aggs":{},"query":{"bool":{"must":[{"bool":{"should":[{"term":{"source":{"value":"compose_dockerlog_1"}}}]}},{"range":{"date":{"gte":"1618984493320"}}},{"range":{"date":{"lte":"1618984493329"}}}]}},"size":1,"sort":[{"date":{"missing":"_first","order":"desc","unmapped_type":"string"}}],"track_total_hits":true}`

	// var e1 map[string]interface{}
	// err := json.Unmarshal([]byte(e1Json), &e1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(e1)
	// fmt.Println(e1["query"])

	// query := e1["query"]
	r1 := reindex{
		Query: "222",
	}

	fmt.Println("r1--", r1)

	res, err := json.Marshal(r1)

	if err != nil {
		fmt.Println("errr")
		fmt.Println(err)
	}

	fmt.Println("res--", string(res))
	strInfo := string(res)

	var ss reindex

	err = json.Unmarshal([]byte(strInfo), &ss)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	fmt.Println("ss ", ss)

}
