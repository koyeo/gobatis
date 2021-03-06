package main

import (
	"encoding/xml"
	"fmt"
	"github.com/koyeo/mybatis.go/schema"
	"io/ioutil"
)

func main() {
	//doc := etree.NewDocument()
	//if err := doc.ReadFromFile("xml/mybatis-config.xml"); err != nil {
	//	panic(err)
	//}
	////doc.ReadFromString(``)
	//d, _ := json.MarshalIndent(doc, "", "\t")
	//fmt.Println(string(d))
	////fmt.Println(doc.Tag)
	////fmt.Println(doc.Root().Tag)
	////fmt.Println(doc.Root().Attr)

	data, err := ioutil.ReadFile("./mybatis-config.xml")
	if err != nil {
		panic(err)
		return
	}
	config := new(schema.Configuration)
	err = xml.Unmarshal(data, config)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(config.String())
}
