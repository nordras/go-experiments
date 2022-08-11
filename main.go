
import "fmt"

func main()  {
	fmt.Print("Hello world \n")

	// variables
	const userName string = "Igor";
	var age int = 30
	occupation := "software engineer"
	list := []string

	var user string = "Stuart" //type is string
	var user2 = "Jane" //type is inferred
	x := 2 //type is inferred

	var a string
	var b int
	var c bool
  
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	var i,j string = "Hello","World"

	fmt.Print(i)
	fmt.Print(j)

	var arr1 = [...]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8}
  
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("Igor", age)
	fmt.Printf("%v, %t \n", occupation)
	fmt.Println("Set age:")
	fmt.Println("List:", list)
	fmt.Scan(&age)

	fmt.Println("Igor", age)
	//fmt.Scan()
}