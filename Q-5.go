
package main
import "fmt"
func divide(num1 float64,num2 float64){
	defer func(){
		//recover() is a built-in function in Go that returns
		//nil if current go routine is not panicking
		if r:=recover();r!=nil{
			fmt.Println("Recovered from panic:",r)
		}
	}()
	if num2==0{
		panic("Division by zero is UNDEFINED!!")
	}
	res:=num1/num2
	fmt.Println("num1/num2 = ",res)

}
func main(){
	var num1 float64
	var num2 float64
	fmt.Println("Enter dividend:")
	fmt.Scan(&num1)
	fmt.Println("Enter divisor:")
	fmt.Scan(&num2)
	divide(num1,num2)
	



}
