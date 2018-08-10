package main

import (
	"encoding/json"

)

//author is :@leo

type Food_info struct{
	Id string     `json:"Id"`
	Name string   `json:"Name"`
	Desc string    `json:"Desc"`
	Price string    `json:"Price"`
	Number string   `json:"Number"`

}

type Foodlist struct{
	Food []Food_info
	User  string
	Data  string
	Status string
	DeskNumber string
}


func JsonEncode(s string) (Foodlist,error) {
	/*
	输入一个字符串例如：
	s := `{"FOod":
       [{"Id":"2","Name":"D","Desc":"Gkf","Number":"2"}]}`
	返回一个json结构
	 */
	var b =Foodlist{}
	err := json.Unmarshal([] byte(s), &b)
	if err!=nil {
		return b,err
	}

	return b,nil
}

func JsonDecode(Json Foodlist) (string,error) {
	/*
	 输入一个jsond对象，返回一个字符串
	{"Food":[{"Id":"2","Name":"D","Desc":"Gkf","Number":"2"}]}
	 */
	v,err:=json.Marshal(Json)
	if err!=nil{
		return "",err
	}

	return string(v),nil
}


/*
type updatejson struct{
	Id string        `json:"id"`
	Dataname string  `json:"dataname"`
	Data   string    `json:"data"`
}



func jsoncod(jsonname string)(updatejson){
	var j updatejson

	err := json.Unmarshal([] byte(jsonname),&j)
	if err!=nil{
		return updatejson{}
	}
	return j
}

*/
/*
func main(){
	/
	s := `{"Food":
       [{"Id":"2","Name":"D","Desc":"Gkf","Number":"2"},{"Id":"2","Name":"D","Desc":"Gkf","Number":"2"}]}`
	/
	s := `{"Food":["Desc","Number","Id","Name","Desc","Number","Id","Name"]}`
	l,_:= JsonEncode(s)
	fmt.Println(l)
}
*/