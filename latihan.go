package main
import "fmt"
const NMAX int = 100
type tabInt[NMAX]int
type tabStr[NMAX]string
func main(){
	var A tabInt
	var B tabStr
	var n, i, idx, j, temp int
	var s string
	fmt.Scan(&n)
	for i = 0; i < n;i ++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n;i++{
		idx = A[i]
		for j = i+1; j <n; j++{
			if idx < A[j]{
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
	for i = 0; i < n; i++{
		if A[i] == 1{
			B[i] = "Gold Medal"
		}else if A[i] == 2{
			B[i] = "Silver Medal"
		}else if A[i] == 3{
			B[i] = "Bronze Medal"
		}else{
			s = fmt.Sprintf("%d", i+1)
			B[i] = s
		}
	}
	for i = 0; i < n; i++{
		fmt.Print(B[i], " ")
	}
}