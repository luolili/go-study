package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	NAME string `xml:"name,attr"` //吧NAME变为person的节点属性, ,前面是属性名
	AGE  int
}

func main() {
	p := person{"de", 12}
	data, err := xml.MarshalIndent(p, "#", " ")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		//<person><NAME>de</NAME><AGE>12</AGE></person>
		/*
				indent:
				<person>
		 <NAME>de</NAME>
		 <AGE>12</AGE>
		</person>

				#<person>
		# <NAME>de</NAME>
		# <AGE>12</AGE>
		#</person>
		*/
		fmt.Println(string(data))

	}

	p2 := new(person)
	err01 := xml.Unmarshal(data, p2)
	if err01 != nil {
		fmt.Println(err01)
		return
	}
	fmt.Println(p2) //&{de 12} &指针 {}结构体
}
