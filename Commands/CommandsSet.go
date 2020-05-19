package Commands

var Commands map[string]ICommand //用于保存所有命令

func init()  {
	Commands=make(map[string]ICommand)
}
type CommandSet struct {} //用户操作Commands集合的 类

func NewCommandSet() *CommandSet  {
	return &CommandSet{}
}
func(this *CommandSet) Register (name string,cmd ICommand)  {
	 Commands[name]=cmd
	 //log.Println("命令注册:",name)
}
func(this *CommandSet) Each (fn func(item ICommand))  {
	 for _,v:=range Commands{
		 fn(v)
	 }
}
