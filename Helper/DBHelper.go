package Helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)
//把数据结构 映射成 map切片
func DBMap(columns []string,rows *sql.Rows)  ([]interface{},error){
	allRows:=make([]interface{},0) //所有行  大切片
	for rows.Next(){
		oneRow:=make([]interface{},len(columns))  //定义一行切片
		scanRow:=make([]interface{},len(columns))
		fieldMap:=make(map[string]interface{})
		for i,_:=range oneRow{
			scanRow[i]=&oneRow[i]
		}
		err:=rows.Scan(scanRow...)
		if err!=nil{
			return nil,err
		}
		for i,val:=range oneRow{
			v,ok:=val.([]byte)  //断言
			if(ok){
				fieldMap[columns[i]]=string(v)
			}
		}
		allRows=append(allRows,fieldMap)
	}
	return allRows,nil
}
type DBModel struct {
	TableName string
	Data []interface{} //map切片
}
type MyDB struct{
	*gorm.DB
}
//获取表结构
func(this *MyDB) DescTable(tableName string) ( *DBModel,error){
	rows,err:= this.Raw("desc "+tableName).Rows()
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	columns,err:=rows.Columns()
	if err!=nil{
		return nil,err
	}
	data,err:=DBMap(columns,rows)
	if err!=nil{
		return nil,err
	}
	return &DBModel{Data:data,TableName:tableName},nil

}

func (this *MyDB) GetTable() ( []interface{},error) {
	rows,err := this.Raw("show tables").Rows()
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	columns,err:=rows.Columns()
	if err!=nil{
		return nil,err
	}
	data,err:=DBMap(columns,rows)
	if err!=nil{
		return nil,err
	}
	return data,nil
}

func NewDB(driver string, dsn string) *MyDB {
	db,err:=gorm.Open(driver,dsn)
	if err!=nil{
		log.Fatal(err)
	}
	return &MyDB{db}
}
