func main(){
	fmt.Printf("the first way recursion [%d]\n",recursion(5))
	fmt.Printf("5 to the power of 2 [%d]\n",recursion2(5,2))
	fmt.Printf("the position of number 4 in fibnacci [%d]\n",fibnacci(4))
	for a := 0 ; a < 10 ; a++{
		fmt.Printf("fibnacci seaqense [%d]",fibnacci2(a))
	}
	fmt.Printf("\nthe sq root 36 [%d]\n",sqr(36))
	fmt.Printf("the sq root 5 [%d]\n",sqr(5))
	fmt.Printf("Prime [%v]\n",Prime(36))
	usingfor(5,2)
}

func recursion(t int)int{
	if t < 0 {
		return 0
	}
	if t == 0 {
		return 1
	}

	return t * recursion(t-1) // 
}


func recursion2(s int, t int)int{
	if t == 0 {
		return 1
	}
	
	return s * recursion2(s, t -1) // 5 * 5 * 5 * 5
}

func fibnacci(s int)int{
	if s == 0 {
		return 0
	}
	
	if s == 1 {
		return 1
	}
	
	return fibnacci(s-1) + fibnacci(s-2)   
}

func fibnacci2(s int) int{
	if s == 0 {
		return 0
	}
	
	if s == 1 {
		return 1
	}

	return fibnacci2(s-1) + fibnacci2(s-2)
}

func Prime(s int)bool{
	for a := 2 ; a * a <= s ;a++{
		if s % a == 0 {
			return false
		}
	}
	return true
}

func sqr(s int)int{
	for a := 2 ; a * a <=s ; a++{
		if a * a == s{
			return a
		}
	}
	return 0
}

func usingfor(s int,t int){
	i := 1
	for a := 1 ; a <= t ; a++{
		i = i * s
	}
	fmt.Printf("5 to the power of 2 using for [%d]",i)
}