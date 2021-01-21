package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

type Userinfo struct {
	Uid int `orm:"pk"`
	Username string
	Departname string
	Created time.Time
}
type User struct {
	Uid int `orm:"pk"`
	Name string
	Profile *Profile `orm:"rel(one)"`
	Post []*Post `orm:"reverse(many)"`
}
type Profile struct {
	Id int
	Age int16
	User *User `orm:"reverse(one)"`
}
type Post struct {
	Id int
	Title string
	User *User `orm:"rel(fk)"`
	Tags []*Tag `orm:"rel(m2m)"`
}
type Tag struct {
	Id int
	Name string
	Posts []*Post `orm:"reverse(many)"`
}
func init()  {
	orm.RegisterDataBase("default","mysql","root:123456@/hello?charset=utf8",30)
	orm.RegisterModel(new(Userinfo),new(User),new(Profile),new(Tag),new(Post))
	orm.RunSyncdb("default",false,true)
	orm.Debug = true
}

func main()  {
	o:= orm.NewOrm()
	//profile := new(Profile)
	//profile.Age = 20
	//
	//user := new(User)
	//user.Profile = profile
	//user.Name = "warcello"
	//
	//ins_pro, err := o.Insert(profile)
	//if err !=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(ins_pro)
	//ins_user, err := o.Insert(user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(ins_user)

	//查询
	profile := Profile{Id: 15}
	err := o.Read(&profile)
	if err == orm.ErrNoRows {
		fmt.Printf("查询不到: %v \n",err)
	} else if err == orm.ErrMissPK {
		fmt.Printf("找不到主键，%v \n",err)
	} else {
		fmt.Println(profile.Id,profile.Age)
	}
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(profile)

	//user := User{Name: "bohe"}
	////插入
	//id, err := o.Insert(&user)
	//fmt.Print("ID: %d , ERR: %v \n", id, err)
	//
	////更新
	//user.Name = "yinuo"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, Err: %v \n",num,err)

	//读取
	//u := User{Id: 1}
	//err := o.Read(&u)
	//fmt.Printf("Err: %v\n", err)



}