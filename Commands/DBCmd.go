package Commands

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/phpwei/go-Gentity/Helper"
	"github.com/phpwei/go-Gentity/TplParser"
	"log"
	"os"
	"strings"
)

func init() {
	vc := NewDBCmd()
	vc.Register("DBCommand", vc) //注册命令
}

type DBCommand struct {
	//用来处理静态资源的生成
	DSN       *string
	TableName *string
	IsCreate  *bool
	CommandSet
	ServiceCommandSet *flag.FlagSet
}

func NewDBCmd() *DBCommand {
	var dsn, tablename = "", ""
	return &DBCommand{DSN: &dsn, TableName: &tablename,
		ServiceCommandSet: flag.NewFlagSet("db args", flag.ExitOnError)}
}
func (this *DBCommand) Init() {
	//go run main.go resource
	if len(os.Args) > 2 && os.Args[1] == "db" {
		this.TableName = this.ServiceCommandSet.String("m", "", "create model")
		this.IsCreate = this.ServiceCommandSet.Bool("y", false, "create ")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}
func (this *DBCommand) Run() {
	config, err := Helper.LoadConfig()
	if err != nil {
		log.Fatal("missing config-file:application.yaml")
	}
	var input string
	if *this.TableName == "all" {
		db := Helper.NewDB(config.DB.Driver, config.DB.DSN)
		data, err := db.GetTable()
		if err != nil {
			log.Fatal("model error:", err.Error())
		}
		for _, k := range data {
			if v, ok := k.(map[string]interface{}); ok {
				ret, err := db.DescTable(fmt.Sprintf("%s", v["Tables_in_test"])) //DBModel 包含了 tablename 以及 字段名等
				if err != nil {
					log.Fatal("model error:", err.Error())
				}
				if !*this.IsCreate {
					fmt.Println("Entity ", Helper.RemoveTablePrefix(fmt.Sprintf("%s", v["Tables_in_test"])))
					fmt.Println("Please confirm (yes|no)[default:no]")
					inputReader := bufio.NewReader(os.Stdin)
					input, err = inputReader.ReadString('\n')
					if err == nil {

						if strings.Trim(input, "\r\n") == "yes" || strings.Trim(input, "\r\n") == "y" {
							this.createModel(ret, Helper.RemoveTablePrefix(fmt.Sprintf("%s", v["Tables_in_test"])))
						}
					}
				}else {
					this.createModel(ret, Helper.RemoveTablePrefix(fmt.Sprintf("%s", v["Tables_in_test"])))
				}
			}
		}

	} else if *this.TableName != "" { // 生成模型
		db := Helper.NewDB(config.DB.Driver, config.DB.DSN)

		data, err := db.DescTable(fmt.Sprintf("%s%s", config.DB.Prefix, *this.TableName)) //DBModel 包含了 tablename 以及 字段名等
		if err != nil {
			log.Fatal("model error:", err.Error())
		}
		if !*this.IsCreate {
			fmt.Println("Entity ", *this.TableName)
			fmt.Println("Please confirm (yes|no)[default:no]")
			inputReader := bufio.NewReader(os.Stdin)
			input, err = inputReader.ReadString('\n')
			if err == nil {

				if strings.Trim(input, "\r\n") == "yes" || strings.Trim(input, "\r\n") == "y" {
					this.createModel(data, *this.TableName)
				}
			}
		} else {
			this.createModel(data, *this.TableName)
		}

	}
}

func (this *DBCommand) createModel(data *Helper.DBModel, tableName string) {
	dm := TplParser.NewDBModelParser()
	dm.Parse(data, tableName)
	fmt.Println("模型生成完成")
}
