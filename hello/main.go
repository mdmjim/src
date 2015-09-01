package main

import (
	"fmt"
	//"strings"
	"Rsystem"
)

func main() {
	 
	var str= Rsystem.Qstring{"helloRsystemhelloworld"}
	fmt.Println(str)
	fmt.Println(str ,"Len:",str.GetLen())
	fmt.Println("Has Prefix hell:",str.HasPrefix("hell")," Has Prefix well:",str.HasPrefix("well"))
    fmt.Println("Has suffix tem:",str.HasSuffix("tem")," Has suffix tom:",str.HasSuffix("tom"))
	fmt.Println("Index of lo:",str.Index("lo"),"Index of ao:",str.Index("ao"))
	fmt.Println("last index of lo:",str.LastIndex("lo")," last index of bo:",str.LastIndex("bo"))
	fmt.Println("split:",str.Split("h"))
	fmt.Println(str.Repeat(2))
	fmt.Println(str.Replace("he","we"))
	
}
