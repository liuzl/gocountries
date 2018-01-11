package gocountries

import (
	"testing"
)

func TestFindCountryByAlpha(t *testing.T) {
	ret := FindCountryByAlpha("CN")
	ret = FindCountryByAlpha("US")

	t.Log(ret)
}
