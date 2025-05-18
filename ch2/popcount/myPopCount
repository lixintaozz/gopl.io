import (
	"fmt"
	"time"
)

var popTable [256]byte

func init() {
	for index, _ := range popTable {
		popTable[index] = byte(index&1) + popTable[index>>1]
	}
}

func main() {
	var num uint64
	fmt.Scanln(&num)

	//方法1使用8条独立的指令计算popcount，编译器不能对其进行优化
	var res int
	start := time.Now()
	res = int(popTable[byte(num)] +
		popTable[byte(num>>8)] +
		popTable[byte(num>>16)] +
		popTable[byte(num>>24)] +
		popTable[byte(num>>32)] +
		popTable[byte(num>>40)] +
		popTable[byte(num>>48)] +
		popTable[byte(num>>56)])

	fmt.Println(res)
	fmt.Println("Method1 time taken:", time.Since(start).Seconds())

	//方法2使用for循环操作更快的原因可能在于：编译器会对循环结构进行优化
	start = time.Now()
	res = func(num uint64) int {
		var count int
		for i := 0; i < 64; i = i + 8 {
			count += int(popTable[byte(num>>i)])
		}

		return count
	}(num)
	fmt.Println(res)
	fmt.Println("Method2 time taken:", time.Since(start).Seconds())

	//方法3用移位和与操作来计算popcount
	start = time.Now()
	res = func(num uint64) int {
		var count int
		for i := 0; i < 64; i++ {
			count += int(num & 1)
			num >>= 1
		}
		return count
	}(num)
	fmt.Println(res)
	fmt.Println("Method3 time taken:", time.Since(start).Seconds())

	//方法4使用x &= x - 1来计算popcount
	start = time.Now()
	res = func(num uint64) int {
		var count int
		for num != 0 {
			num &= num - 1
			count++
		}
		return count
	}(num)
	fmt.Println(res)
	fmt.Println("Method4 time taken:", time.Since(start).Seconds())
}
