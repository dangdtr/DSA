package main
import (
	"github.com/dawnpanpan/go-dsa/algs"
	"github.com/dawnpanpan/go-dsa/algs/structure"
	"fmt"

)
func main() {
	pq_test()
}

func pq_test() {
	var key = algs.GenerateInterfaceSlice(10)
	// key := sort.StringSliceToInterface(k)
	fmt.Println(key...)
	
	maxPq := structure.NewMaxPQFrom(key)
	fmt.Println(maxPq.Show())

}
