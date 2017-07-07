package mobi

import (
	"fmt"
	_ "io/ioutil"
	"testing"
)

func TestMobiBook_Parse(t *testing.T) {
	var b = NewMobi("../tmp/zcdz.mobi")
	b = NewMobi("../tmp/st1.azw3")
	b = NewMobi("../tmp/wms.mobi")
	fmt.Println(b.Parse())
	fmt.Println(b.Title())
	fmt.Println(b.Author())
}
