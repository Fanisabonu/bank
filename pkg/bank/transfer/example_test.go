package transfer
import ("fmt")
import "bank/pkg/bank/types"
func ExampleTotal()  {
fmt.Println(Total(0))
fmt.Println(Total(500_000))
fmt.Println(Total(1_000_000))

 //Output:
	 // 0
	// 502500
	//1005000

}

