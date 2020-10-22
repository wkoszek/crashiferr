package errhand

import (
    "strconv"
    "testing"
)

func TestDescr(t *testing.T) {
	_, err := strconv.Atoi("123a")
	expected := `strconv.Atoi: parsing "123a": invalid syntax`
	got := DescrIfErr(err)
	if got != expected {
		t.Errorf("DescrIfErr Expected: %v, got: %v", expected, got)
	}
}
