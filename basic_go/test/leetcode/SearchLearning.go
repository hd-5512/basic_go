package leetcode

/*******************************************************************************************************/
//顺序查找 线性查找  返回第几个  O(n)  如果二维数组 就是 O(n*n)  根据数据规则 可以通过组合算法  O(nlogn)
func TestSequential_Search(arr []int,key int) int {
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == key {
			return i+1
		}
	}

	return 0
}

//哨兵优化 O(n)  最快 O(1) 适用小数据 或将概率出现大的放前面时候
//for循环过程中 每次都需要比对 i和n 做一次判断  加上 arr[i] == key 值 等于每次循环中比较有2次
//这里可以优化一次判断 大大提升性能

func TestBetterSeq_Search(arr []int,key int) int {

	key = 5
	if(arr[0] == key) {
		return 0
	}

	arr[0] = key  //通过设立一个0位置的哨兵  一定会退出循环 同时 返回 -1
	n := len(arr)
	for arr[n-1] != key {
		n--
	}

	return n-1
}

/**********************************************************************************************************************/

//二分查找 O(logn)  适用于很多情况
//必要的条件是一个数据有序排列   在开发环境中 获取数据的方式通常都是可以有意的去进行有序化  比如 sql 中的order
func TestBinary_Search(arr []int,key int) int {
	//arr = []int{1,2,3,4,5,6,7,8}
	//key = 6
	min := 0
	max := len(arr) - 1
	mid := 0

	for min <= max {
		mid = ( min + max ) / 2
		if key < arr[mid] {
			max = mid - 1
		} else if key > arr[mid] {
			min = mid + 1
		} else {
			return mid
		}
	}

	return 0
}


//插值查找
//比如 取值范围 0-10000 的100个数从小到大平均分布情况下 使用 插入查找更快
//基于二分查找 但是效率样本数量高平均分布情况的时候更高 理论时间复杂度还是 O(logn)
// mid = ( min + max ) / 2  = min + 1/2 * (max - min) = min + (key-arr[min]) / (arr[max]-arr[min] ) * (max - min)

func TestInsert_Search(arr []int,key int) int {
	min := 0
	max := len(arr) - 1
	mid := 0

	for min <= max {
		mid = min + (key-arr[min]) / (arr[max]-arr[min]) * (max - min)
		if key < arr[mid] {
			max = mid - 1
		} else if key > arr[mid] {
			min = mid + 1
		} else {
			return mid
		}
	}

	return 0
}


//斐波拉契查找 是利用了黄金分割的原理
//时间复杂度夜市 O(logn) 但是如果需要查找的数据在 列表的中间部分  优势会比较明显

func TestFibonacci_Search(arr []int,key int) int {
	min := 0
	max := len(arr) - 1
	mid := 0

	//Fib := Fibonacci1(max)

	//找到在 Fib数列中的位置  为了后面可以有根据的切割
	k := 0
	for max >= Fibonacci_digui(k) {
		k++
	}

	//相比上面的方式效率更高
	//k := 0
	//for max > Fibonacci_digui(k) - 1 {
	//	k++
	//}

	//Flen := Fibonacci_digui(k)
	//arr2 := make([]int,Flen)

	//将查找的数组 arr 补足数量到 斐波拉契的值  都使用自己最大的值去填充
	//for i:= 0; i<Flen; i++ {
	//	if i >= max {
	//		arr2[i] = arr[max]
	//	} else {
	//		arr2[i] = arr[i]
	//	}
	//}

	for i:=len(arr); i<Fibonacci_digui(k); i++ {
		arr = append(arr,arr[max])
	}


	for min <= max {
		mid = min + Fibonacci_digui(k-1)
		if key < arr[mid] {
			max = mid - 1
			k = k - 1
		} else if key > arr[mid] {
			min = mid + 1
			k = k - 2
		} else {
			if mid <= max {
				return mid  /*相等则说明 mid 即为查找到的位置*/
			}else{
				return max  /*说明是补全的数值  其实也就是 最大值*/
			}
		}
	}

	return 0
}


//递归生成斐波拉契数列
func Fibonacci1(n int) []int {
	if(n == 1){
		return []int{1}
	}

	if(n == 2){
		return []int{2}
	}

	f := make([]int,n)
	f[0] = 0
	f[1] = 1
	for i:=2; i<n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func Fibonacci2(n int) []int {
	F := make([]int,n)
	for i:=0; i<n; i++ {
		F[i] = Fibonacci_digui(i)
	}
	return F
}

func Fibonacci_digui(i int) int {
	if i < 2 {
		if i==0 {
			return 0
		} else {
			return 1
		}
	}else {
		f := Fibonacci_digui(i-1) + Fibonacci_digui(i-2)
		return f
	}
}



/**********************************************************************************************************************/


//索引查找 树形索引和多级索引

//线性索引查找 : 稠密 / 分块 / 倒排







