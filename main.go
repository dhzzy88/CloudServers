package main
import (
	"net/http"
	"io"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"time"
	"io/ioutil"
	"strconv"
)

const root ="1111"


func Getroot(w http.ResponseWriter,r *http.Request){

	var A Foodlist
	A.User = root
	var list= Select(A)
	liststring,err:=JsonDecode(list)
	if err!=nil{
		log.Fatal("err",err)
		return
	}
	io.WriteString(w,liststring)
}

func Postroot(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
 	s,_ := ioutil.ReadAll(r.Body)
 	json,_:= JsonEncode(string(s))
 	date:= time.Now()
 	_,m,d:=date.Date()
 	json.User = root
 	json.Data = m.String()+strconv.Itoa(d)
 	for i:=0;i<len(json.Food);i++{
 		Insert(json.Food[i],json)
	}

	project := Select(json)
	js ,_:=JsonDecode(project)
	io.WriteString(w,js)
}

func GetWeixin(w http.ResponseWriter,r *http.Request){
	var A Foodlist
	A.User = root
	list := Select(A)
	var Price Foodlist

	Price = list
	Price.User = "1111"
	fmt.Println(Price)
	JSON_STRING,_ :=JsonDecode(Price)
	io.WriteString(w,JSON_STRING)
}

func PostWeixin(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s,_ := ioutil.ReadAll(r.Body)
	fmt.Println(r.Form)

	fmt.Println(string(s))
	Foo,_:=JsonEncode(string(s))
	date := time.Now()
	_,m,d:=date.Date()
	mstr:=m.String()
	Foo.Data=mstr+strconv.Itoa(d)
	fmt.Println(Foo.Data)

	for i:=0;i<len(Foo.Food);i++ {

		Insert(Foo.Food[i],Foo)

	}
	fmt.Println(Foo.Data)
	foo:= Select(Foo)
	JsonString,_ := JsonDecode(foo)


	io.WriteString(w,JsonString)
}


func Updateroot(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s,_ := ioutil.ReadAll(r.Body)
	v ,err:= JsonEncode(string(s))
	CheckError(err)
	affect :=Upgrade(v)
	io.WriteString(w,string(affect))

}

func Deleteroot(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s,_ := ioutil.ReadAll(r.Body)
	v,err := JsonEncode(string(s))
	CheckError(err)
	affect := Remove(v)
	io.WriteString(w,string(affect))
}

func Updateuser(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s,_ := ioutil.ReadAll(r.Body)
	v ,err:= JsonEncode(string(s))
	CheckError(err)
	affect :=Upgrade(v)
	io.WriteString(w,string(affect))
}

func Deleteuser(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s,_ := ioutil.ReadAll(r.Body)
	v,err := JsonEncode(string(s))
	CheckError(err)
	affect := Remove(v)
	io.WriteString(w,string(affect))
}






func main() {
	var router= mux.NewRouter()
	router.HandleFunc("/", Getroot).Methods("GET")
	router.HandleFunc("/", Postroot).Methods("POST")
	router.HandleFunc("/root/update",Updateroot).Methods("POST")
	router.HandleFunc("/root/delete",Deleteroot).Methods("POST")
	router.HandleFunc("/weixin", GetWeixin).Methods("GET")
	router.HandleFunc("weixin/update",Updateuser).Methods("POST")
	router.HandleFunc("weixin/delete",Deleteuser).Methods("POST")
	router.HandleFunc("/weixin", PostWeixin).Methods("POST")
	err:= http.ListenAndServe(":8000",router)
	if err !=nil{
		log.Fatal("err",err)
	}
}