package Commands

import (
	"flag"
)

func init()  {

	 cs:=NewCommandSet()
	 cs.Each(func(item ICommand) {
		 item.Init()
	 })
	flag.Parse()

	 cs.Each(func(item ICommand) {
		item.Run()
	})



}
func Parse()  {

}