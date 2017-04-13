package randomuser

import "testing"

func TestGenerate(t *testing.T) {
	var countResults = 5
	c := &QueryConfig{countResults, "", "", ""}
	users, err := Generate(c)
	if err != nil {
		t.Error(err)
		return
	}

	if len(users) != countResults {
		t.Error("Expected 5, got ", len(users))
	}
}
