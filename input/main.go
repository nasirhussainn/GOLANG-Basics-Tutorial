package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Nasir")
	reader :=  bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name:")

	// comma ok || err ok name, err - or _, err in loop etc. as there is no catch block
	name, _:= reader.ReadString('\n')
	fmt.Println("Your Name is: ", name)
	fmt.Printf("type of input %T", name)

}