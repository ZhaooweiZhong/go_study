package main
import(
	"fmt"
	"time"
)
func newTask(){
	i:=0
	for{
		i++
		fmt.Printlf("new Goroutine :i = %d\n",i)
		time.Sleep(1 * time.Second)
	}
	
}

func main(){
	go newTask()
}