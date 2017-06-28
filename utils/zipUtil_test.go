package utils

import (
	"fmt"
	"testing"
)

func TestNewZip(t *testing.T) {
	fmt.Println(t.Name())
	zip := NewZip("../tmp/scm.epub")
	defer zip.Close()
	fmt.Printf("%+v \n", zip)
}

func TestZip_Decompress(t *testing.T) {
	fmt.Println(t.Name())
	zip := NewZip("../tmp/scm.epub")
	defer zip.Close()
	fmt.Printf("%+v \n", zip)
	if err := zip.Decompress(); err == nil {
		fmt.Println("ok")
	} else {
		fmt.Println("error :", err)
	}
}
