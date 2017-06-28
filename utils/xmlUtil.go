package utils

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"path"
	//github.com/PuerkitoBio/goquery
	"errors"
)

type Xml struct {
	FullPathFilename string
	Filename         string
	Ext              string
}

func NewXml(fullPathFilename string) *Xml {
	xml := &Xml{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
	}
	return xml
}

func (x Xml) Decode() error {
	//if bts, err := ioutil.ReadFile(x.FullPathFilename); err == nil {
	//	inputReader := bytes.NewReader(bts)
	//	decoder := xml.NewDecoder(inputReader)
	//	var t xml.Token
	//	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
	//		switch token := t.(type) {
	//		// 处理元素开始（标签）
	//		case xml.StartElement:
	//			name := token.Name.Local
	//			fmt.Printf("Token name: %s\n", name)
	//			for _, attr := range token.Attr {
	//				attrName := attr.Name.Local
	//				attrValue := attr.Value
	//				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
	//			}
	//			// 处理元素结束（标签）
	//		case xml.EndElement:
	//			fmt.Printf("Token of '%s' end\n", token.Name.Local)
	//			// 处理字符数据（这里就是元素的文本）
	//		case xml.CharData:
	//			content := string([]byte(token))
	//			fmt.Printf("This is the content: %v\n", content)
	//		default:
	//			// ...
	//		}
	//	}
	//	if err.Error() != "EOF" {
	//		return err
	//	}
	//	return nil
	//} else {
	//	return err
	//}
	return nil
}

func (x Xml) GetValue(path string) (string, error) {
	if strs, err := x.GetValues(path); err == nil {
		if len(strs) != 0 {
			return strs[0], nil
		} else {
			return "", nil
		}
	} else {
		return "", err
	}
}

func (x Xml) GetValues(path string) ([]string, error) {
	bts, err := ioutil.ReadFile(x.FullPathFilename)
	doc := etree.NewDocument()
	if err = doc.ReadFromBytes(bts); err == nil {
		tmpElements := doc.FindElements(path)
		if len(tmpElements) != 0 {
			results := make([]string, 0)
			for _, value := range tmpElements {
				results = append(results, value.Text())
			}
			return results, nil
		} else {
			return make([]string, 0), errors.New("not found " + path)
		}

	} else {
		return make([]string, 0), err
	}
}

func (x Xml) GetAttr(path, key string) (string, error) {
	bts, err := ioutil.ReadFile(x.FullPathFilename)
	doc := etree.NewDocument()
	var element *etree.Element
	if err = doc.ReadFromBytes(bts); err == nil {
		tmpElements := doc.FindElements(path)
		if len(tmpElements) != 0 {
			element = tmpElements[0]
			return element.SelectAttrValue(key, ""), nil
		} else {
			return "", errors.New("not found " + path)
		}

	} else {
		return "read bytes error", err
	}
}
