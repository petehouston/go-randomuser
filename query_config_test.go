package randomuser

import (
	"net/url"
	"strconv"
	"testing"
)

func TestQueryConfig(t *testing.T) {
	u, _ := url.Parse("")
	c := &QueryConfig{10, "", "", ""}
	c.encode(u)

	if cntResults, _ := strconv.ParseInt(u.Query().Get(KEY_RESULTS), 10, 64); cntResults != 10 {
		t.Error("Expected 10, got ", cntResults)
	}

	c.MaxResults = 99
	c.encode(u)
	if cntResults, _ := strconv.ParseInt(u.Query().Get(KEY_RESULTS), 10, 64); cntResults != 99 {
		t.Error("Expected 99, got ", cntResults)
	}
}
