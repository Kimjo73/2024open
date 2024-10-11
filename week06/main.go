package main
 
import(
      "fmt"
	  "reflect"
)
func main(){
	i:=13
	var f float64 =12.9
	c1 := 'A'
	c2 := '2'

	fmt.printf("value i: %d,value f: %f\n",i,f)
	fmt.println("%d*%f = %d\n",i,f,i*int(f))
	fmt.println(c1,c2)
	fmt.println(reflect.typeOf(i),reflect,TypeOf(f),reflect.TypeOf(c2))
}