package utils

import (
	_ "bytes"
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"path"
	"testing"
)

func TestNewXml(t *testing.T) {
	fmt.Println(t.Name())
	xml := NewXml("/tmp/430bf167-1e9c-42de-80fd-4fd8bafe4f50/META-INF/container.xml")
	fmt.Printf("%+v \n", xml)
}

func TestXml_Decode(t *testing.T) {
	fmt.Println(t.Name())
	xml := NewXml("/tmp/430bf167-1e9c-42de-80fd-4fd8bafe4f50/META-INF/container.xml")
	fmt.Printf("%+v \n", xml)
	if err := xml.Decode(); err == nil {
		fmt.Println("ok")
	} else {
		fmt.Println("error :", err)
	}
}

func TestXml_GetAttr(t *testing.T) {
	fmt.Println(t.Name())
	zip := NewZip("../tmp/scm.epub")
	defer zip.Close()
	fmt.Printf("%+v \n", zip)
	if err := zip.Decompress(); err == nil {
		xml := NewXml(path.Join(zip.TmpDir, "META-INF", "container.xml"))
		fmt.Printf("%+v \n", xml)
		fmt.Println(xml.GetAttr("//container/rootfiles/rootfile", "media-type"))
		fmt.Println(xml.GetAttr("//container/rootfiles/rootfile", "full-path"))
		fmt.Println(xml.GetAttr("//container/rootfiles/noexist", "full-path"))
		fmt.Println(xml.GetAttr("//container/rootfiles/rootfile", "noexist"))
	} else {
		fmt.Println("error :", err)
	}
}

func TestXml_GetValue(t *testing.T) {
	fmt.Println(t.Name())
	zip := NewZip("../tmp/scm.epub")
	defer zip.Close()
	fmt.Printf("%+v \n", zip)
	if err := zip.Decompress(); err == nil {
		xml := NewXml(path.Join(zip.TmpDir, "META-INF", "container.xml"))
		fmt.Printf("%+v \n", xml)
		if pathname, err := xml.GetAttr("//container/rootfiles/rootfile", "full-path"); err == nil {
			xml1 := NewXml(path.Join(zip.TmpDir, pathname))
			fmt.Println(xml1.GetValue("//package/metadata/dc:title"))
			fmt.Println(xml1.GetValue("//package/metadata/dc:noexist"))
		}
	} else {
		fmt.Println("error :", err)
	}
}
