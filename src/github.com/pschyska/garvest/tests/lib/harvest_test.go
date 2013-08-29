package lib

import (
	"fmt"
	"github.com/pschyska/garvest/app/lib"
	"testing"
)

func TestConnect(t *testing.T) {
	harvest := lib.Harvest{}
	resp, err := harvest.Connect()
	if err != nil {
		t.Fail()
	}
	fmt.Println("%v\n", resp)
}
