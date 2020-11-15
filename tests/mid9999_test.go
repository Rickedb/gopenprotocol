package midtests

import (
	"testing"

	keepalive "github.com/rickedb/gopenprotocol/keepalive"
)

func TestRevision1(t *testing.T) {
	pkg := "00209999001       "
	mid := new(keepalive.Mid9999)
	err := mid.Parse(pkg)
	if err != nil {
		t.Errorf("Error ocurred while parsing: %s", err.Error())
		t.FailNow()
	}

	gg, err := mid.Pack()

	if err != nil {
		t.Errorf("Error ocurred while packing: %s", err.Error())
		t.FailNow()
	}

	if pkg != gg {
		t.Errorf("Package is not the same. Expected: %s \n Actual: %s", pkg, gg)
		t.FailNow()
	}

}
