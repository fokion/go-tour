package main
import(
  "fmt"
)
func swap(x int ,y  string) (string,int){
  return y, x
}
func main()  {
a,b := swap(3,"test");
fmt.Println(a,b)  
fmt.Printf("Test %T type and %v value\n",a,a)
}