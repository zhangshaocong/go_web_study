package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloNmae(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //解析参数，默认是不会解析
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v :=range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v,""))
	}
	fmt.Fprintf(w, "hello astaxie")
}
func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token不能为空")
		} else {
			fmt.Println("token 不存在")
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		fmt.Println("crutime: ", crutime)
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handle,err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handle.Header)
		f, err := os.OpenFile("./test/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main()  {
	http.HandleFunc("/", sayhelloNmae)
	http.HandleFunc("/login",login) //设置login访问的路由
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}