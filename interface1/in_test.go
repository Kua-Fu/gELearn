package interface1

func TestNULL(t *testing.T) {
	ss := nil
	for k,v := range ss{
		fmt.Println(k,v)
	}
}
