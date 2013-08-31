package harvest

import (
	"log"
	"fmt"
	"github.com/pschyska/garvest/app/lib/harvest"
	"testing"
)

func TestConnect(t *testing.T) {
	h := harvest.New()
	resp, err := h.Connect()
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	fmt.Printf("%s\n", resp)
}

