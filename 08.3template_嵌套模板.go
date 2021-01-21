package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	s1,_ := template.ParseFiles("08.header.tmpl","08.content.tmpl","08.footer.tmpl")
	s1.ExecuteTemplate(os.Stdout,"header",nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout,"content",nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout,"footer",nil)
	fmt.Println()
	s1.Execute(os.Stdout,nil)

}