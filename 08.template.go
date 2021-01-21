package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Emails []string
	Friends []*Friend
}
type Friend struct {
	Fname string
}


func main()  {
	//t := template.New("filedname example")
	//t, _ = t.Parse("Hello {{.UserName}}!")
	//p := Person{UserName: "warcello"}
	//t.Execute(os.Stdout, p)
	f1 := Friend{Fname: "xiaoming"}
	f2 := Friend{Fname: "huawei"}
	t := template.New("filedname example")
	t,_ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := Person{UserName: "warcello",Emails: []string{"war@qq.com","hello@qq.com"},Friends: []*Friend{&f1,&f2}}
	t.Execute(os.Stdout,p)
}
