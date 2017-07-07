package mobi

import (
	"testing"
	"fmt"
	_ "io/ioutil"
)

func TestMobiBook_Parse(t *testing.T) {
	var b = NewMobi("../tmp/zcdz.mobi")
	b = NewMobi("../tmp/st1.azw3")
	fmt.Println(b.Parse())
	fmt.Println(b.Title())
	fmt.Println(b.Author())
}
