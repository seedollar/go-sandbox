package main
import "fmt"


func sum(nums []int, ch chan int) {
	var total = 0
	for _, v := range nums {
		total += v
	}
	ch <- total
}

func main() {
	var numbers = []int{1,2,3,4,5,6,7,8,9}
	ch := make(chan int)

	go sum(numbers[len(numbers)/3:], ch)
	go sum(numbers[:len(numbers)/2], ch)
	x,y := <-ch, <-ch

	fmt.Println(x,y)

}
