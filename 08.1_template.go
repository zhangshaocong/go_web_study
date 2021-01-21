package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails []string
	Friends []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " at " + substrs[1])
}

func main()  {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("Filedname example")
	t.Funcs(template.FuncMap{"emailDeal":EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an emails {{.|emailDeal}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my frined name is {{.Fname}}
		{{end}}
		{{end}}
	`)
	p := Person{UserName: "Astaxie",
				Emails: []string{"astaxie@beego.me","aasdfa@gmail.com"},
				Friends: []*Friend{&f1,&f2}}
	t.Execute(os.Stdout,p)
}