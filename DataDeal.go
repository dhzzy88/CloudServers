package main


import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"log"

	"fmt"
)

func CheckError(err error) error{
	if err !=nil{
		log.Fatalf("err%s",err)
		return err
	}else{
		return nil
	}
}


func OpenDb() (*sql.DB,error){
	db,err:=sql.Open("mysql","root:8865439ZZY@tcp(localhost:3306)/bbq?charset=utf8")
	CheckError(err)
	return db,nil
	}

func Insert(Foo Food_info,foo Foodlist) (int64){
	db, err := OpenDb()
	defer db.Close()
	CheckError(err)
	//插入数据
	stmt, err := db.Prepare("INSERT PRICE SET Id=?,Name_Foo=?,Desc_Foo=?,Price=?,Number_Foo=?,Users=?,Date_list=?,DeskNumber=?,Status=?")
	CheckError(err)
	//res, err := stmt.Exec("1111","腌肉","使用养殖的","55",0)
	res, err := stmt.Exec(Foo.Id,Foo.Name,Foo.Desc,Foo.Price,Foo.Number,foo.User,foo.Data,foo.DeskNumber,foo.Status)
	CheckError(err)
	id, err := res.LastInsertId()
	CheckError(err)
	fmt.Println(id)
	return id
}


func Upgrade(foodlist Foodlist)(int64) {
	//更新数据
	db,err := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE PRICE SET Status=? WHERE Id =?")
	CheckError(err)
	res, err := stmt.Exec(foodlist.Status,foodlist.DeskNumber)
	CheckError(err)
	affect, err := res.RowsAffected()
    return affect
}
func Remove(foodlist Foodlist)(int64) {
	//删除数据
	db,err := OpenDb()
	CheckError(err)
	defer db.Close()
	var affect int64
	for i :=0;i<len(foodlist.Food);i++ {
		stmt, err := db.Prepare("DELETE from PRICE WHERE Users=? AND Id=?")
		CheckError(err)
		res, err := stmt.Exec(foodlist.User, foodlist.Food[i].Id)
		CheckError(err)
		affect, err = res.RowsAffected()
		CheckError(err)
	}
		return affect

}

func Select (foodlist Foodlist)(Foodlist){
	//查询数据

	db,err :=OpenDb()
	defer db.Close()
	CheckError(err)

	var User = foodlist.User
	var Data = foodlist.Data
	var Desk = foodlist.DeskNumber
	var Status = foodlist.Status

	if User !=""{
		User = ` Users="`+User+`"`
	}
	if Data !=""{
		Data = ` AND Date_list="`+Data+`"`
	}

	if Desk !=""{
		Desk = ` AND DeskNumber=`+Desk+`"`
	}
	if Status !=""{
		Status = ` AND Status=`+Status+`"`
	}



	query:="SELECT * FROM PRICE "+" WHERE"+User+Data+Desk+Status+";"

	rows,err := db.Query(query)

	CheckError(err)

	var foo =new(Foodlist)
	foo.Food=make([] Food_info,1,1)
	var count  =0


	for rows.Next(){


		var id string
        var name string
        var desc string
        var number string
        var price  string
        var user string
        var data string
        var desk string
        var status string
		err =rows.Scan(&id,&name,&desc,&price,&number,&user,&data,&desk,&status)
		CheckError(err)
		foo.Food = append(foo.Food, foo.Food[count])
		foo.Food[count].Id =id
		foo.Food[count].Price=price
		foo.Food[count].Number= number
		foo.Food[count].Desc =desc
		foo.Food[count].Name = name
		foo.DeskNumber = desk
		foo.Status = status
		foo.Data = data
		foo.User = user
		count++
		//var info = []string{tmpid,name,desc,price,tmpnumber,user}

	}
   return *foo


}

/*
func main5() {
         var fooo Foodlist
         var Food =[] Food_info{{"2","leo","people","15","25",},{"2","leo","people","15","25"}}
         fooo=Foodlist{Food,"2222","0809"}

	     var Food1 =[] Food_info{{"2","12","people","15","25"},{"2","leo","people","15","25"}}
         var fooo1 Foodlist
         fooo1 = Foodlist{Food1,"3333","0809"}
	     Insert(fooo.Food[0],fooo.User,fooo.Data)
	     Insert(fooo1.Food[0],fooo1.User,fooo.Data)

		 foo :=Select("3333","0809")
         fmt.Println(foo)
}
*/