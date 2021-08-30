package main
import(
	"fmt"
	"os"
	"io"
)

func main(){
	tty, err :=os.OpenFile("./helloReflect",os.O_RDWR, 0)

	if err != nil {
		fmt.Println("openfile error",err)
		return 
	}
	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("hello!"))
}