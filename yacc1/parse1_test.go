package yacc1

import (
	"log"
	"testing"
)

func TestParse1(t *testing.T) {

	line := "['select f1, f2 from t1']"
	sql, err := Parse(line)
	if err != nil {
		log.Fatalf("Parse error: %s", err)
	}
	log.Println("get sql:", sql.String())
}
