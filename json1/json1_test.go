package json1

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestJson2(t *testing.T) {
	type ColorGroup struct {
		Sources []string `json:"_source,omitempty"`
		Name    string
		Colors  []string
	}
	group := ColorGroup{
		Sources: []string{"user_name"},
		Name:    "Reds",
		Colors:  []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("--res--", string(b))
}

func TestDeleteDupli(t *testing.T) {
	ns := []string{}
	m := map[string]string{}
	s := []string{"1", "2", "1"}
	for _, i := range s {
		m[i] = "1"
	}
	for k := range m {
		ns = append(ns, k)
	}
	fmt.Println("---", ns)

}

func TestStringList(t *testing.T) {

	// s := []string{"1", "2", "3"}
	// strS, _ := json.Marshal(s)
	strS := "[\"111\",\"222\"]"
	res := []string{}
	fmt.Println("--byte--", []byte(strS))
	err := json.Unmarshal([]byte(strS), &res)
	if err != nil {
		fmt.Println("error---", err)
	}
	fmt.Println("res--", res)
}

func TestMap(t *testing.T) {
	m := map[string]interface{}{
		"1": "111",
		"2": "222",
	}
	res, ok := m["3"].(string)
	if !ok {
		fmt.Println("-----", res)
	}

}

func TestSlice(t *testing.T) {

	s := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	slen := len(s)
	kkk := 3

	s1 := slen / kkk
	s2 := slen % kkk
	fmt.Println("s1, s2--", s1, s2)
	for i := 0; i < s1; i++ {
		k1 := i * kkk
		k2 := (i + 1) * kkk
		fmt.Println("ki, k2--", k1, k2)
	}
	if s2 > 0 {
		fmt.Println("--last--", s1*kkk, slen)
	}

}

func TestPrintES(t *testing.T) {
	names := []string{"n1", "n2"}
	nameStr := strings.Join(names, ",")
	s := fmt.Sprintf(`{"size":1000,"query":{"bool":{"filter":[{"term":{"class":"label"}},{"terms":{"object_name": %s}}]}}}`, names)
	s1 := fmt.Sprintf(`{"size":1000,"query":{"bool":{"filter":[{"term":{"class":"label"}},{"terms":{"object_name": %s}}]}}}`, nameStr)
	fmt.Println(s)
	fmt.Println(s1)
	xx := ""
	sMap := map[string]string{"1": "111"}
	xx, y := sMap["1"]
	fmt.Println(len(sMap), xx, y)
	var z map[string][]string
	zz := []string{"1"}
	zzz := "111"
	z[zzz] = zz
	fmt.Println(z)
}

type Serie struct {
	Name    string            `json:"name"`
	Tags    map[string]string `json:"name"`
	Columns []string          `json:columns`
	Values  []int64           `json:values`
}

type Row struct {
	Name    string            `json:"name,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
	Columns []string          `json:"columns,omitempty"`
	Values  [][]interface{}   `json:"values,omitempty"`
	Partial bool              `json:"partial,omitempty"`
}

type Series struct {
	Series []Serie `json:"series"`
}

type Content struct {
	Content []Series `json:"content"`
}

func TestUnmarsha1(t *testing.T) {

	TestRes := `
{
	"content": [{
		"series": [{
			"name": "df_metering",
			"tags": {
				"workspaceUUID": "wksp_23a08717bb394816891ea5f493f053b5"
			},
			"columns": [
				"time",
				"max(count)"
			],
			"values": [
				[
					1622044801000,
					5
				]
			]
		}]
	}]
}
`
	contents := Content{}
	err := json.Unmarshal([]byte(TestRes), &contents)
	if err != nil {
		fmt.Println("--error---", err)
	}

	fmt.Println(contents)

}

func TestStringSplit(t *testing.T) {
	s := " status != 'ok' "
	s1 := strings.ReplaceAll(s, " ", "")
	s11 := strings.ReplaceAll(s1, "'", "")
	s111 := strings.ReplaceAll(s11, "`", "")
	fmt.Println("s--", s)
	fmt.Println("s1--", s1)
	fmt.Println("s11--", s11)
	fmt.Println("s111--", s111)
	s2 := strings.Split(s111, "!=")
	fmt.Println("s2--", s2)
}
