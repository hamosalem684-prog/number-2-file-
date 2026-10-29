
func main(){
	fmt.Println(number("123"))
	fmt.Println(number1("456"))
	fmt.Println(number2("007891"))
}

func number(s string)int{
	var i int 
	m := 1

	for a:= len(s)-1 ; a >= 0 ; a--{ // we either * 10 
		d := int(s[a] - '0')
		i += d * m
		m*= 10
	}
	return i
}


func number1(s string)int{
	var i int // the box that in fact holds everything 
	m := 1 // that will change from 1 to 3 to 2 and 1 again because it's the last number 

	for a := len(s)-1 ; a >= 0 ; a-- {
		// '3' - '0' 
		// to make 3 * 1 i gotta do that 
		d := int(s[a] - '0') // the original one d is just a varaible so to * any number from this make d * 1 and than 1 * 10 number 1 will change automateclly 
		i += d * m 
		m *= 10   
	}
	return i 
}

############################################################################

func number2(s string) int{
	var i int
	for a := 0 ; a <len(s) ; a++{
		i = (i * 10) + int(s[a] - '0')
	}
	return i
}
 // 7891 is number 9 is ones number and 7 is handsred ? 1 * 9 = 9 ,, 10 * 8 = ,, 7 * 100 = 
 // if we start from number 7 i must multiply it with 10 and every number as long as I'm going to * them with 10 
 // 7 * 10 = ,,  * 10 ,,  * 10 ,,  * 10



######################################################################################

func main(){
	number(321)
	// z01.PrintRune('\n')
	fmt.Print("\n")
	number(231)	
	fmt.Print("\n")
	number(0)
}

func number(s int){
	if s == 0{
		fmt.Print(0)
		return 
	}
	var i []int // [3,2,1]
	for s > 0 {
		i = append(i,s%10)
		s /= 10
	}
	
	for a := 0 ; a <len(i) ; a++{
		for j := a + 1 ; j <len(i) ; j++{
			if i[a] > i[j] {
				i[a] , i[j] = i[j] , i[a]
			}
		}
	}

	for _, r := range i {
		fmt.Print(r)
	} 
}



######################################################################################


func main(){
	number([]int {2,4,5,1,3,})
	fmt.Print("\n")
	number1(9876)
	fmt.Print("\n")
	number1(987654321)
}

func number(s []int){
	for a := 0 ; a <len(s) ; a++{
		for j := a + 1 ; j <len(s) ;j++{
			if s[a] > s[j] {
				s[a] , s[j] = s[j] , s[a]
			}
		}
	}
	for _, i := range s{
		fmt.Print(i)
	}
}

func number1(s int){
	if s == 0 {
		fmt.Print(0)
		return 
	}
	
	var i []int // 9810 
	for s > 0 {// 9810 is > indeed greater than 0 so it's vaild 
		i = append(i ,s % 10 )
		s /= 10
	}

	for a := 0 ; a < len(i) ; a++{
		for h := a + 1 ; h < len(i) ; h++{
			if i[a] > i[h] {
				i[a] , i[h] = i[h] , i[a]
			}
		}
		
	}
	for _, r := range i{
		fmt.Print(r)
	}
}

func main(){
	fmt.Println(number("hamo salem 2008"))
	fmt.Println(Capitalize("hello how are you"))
	fmt.Println(number2("hello how are you"))
	fmt.Println(number2("Hello! How are you? How+are+things+4you?"))
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func number(s string)string{
	var i string
	for _, a := range s{
		if a >= 'a' && a <= 'z' {
			i += string(a-32)
		}else {
			i += string(a)
		}
	}
	return i 

	// a - 32 
	// var i string this to box it 
}

// ("hello how are you"))

func Capitalize(s string)string{
	var i string
	if s[0] >= 'a' && s[0] <= 'z' {
		i += string(s[0]-32)
	}else {
		i += string(s[0])
	}

	for a := 1 ; a < len(s) ; a++{
		if s[a] >= '0' && s[a] <= '9' {
			i += string(s[a])
		}
		if s[a] >= 'a' && s[a] <= 'z' && s[a-1] == ' ' || s[a-1] == '+' {
			i += string(s[a]-32)
		}else{
			i += string(s[a])
		}
    }
	return i
}

func number2(s string) string{
	var i string
    for a := 0; a < len(s); a++ {
    // Look at this order: (a == 0 || s[a-1] == ' ')
        if (a == 0 || s[a-1] == ' ' ) && s[a] >= 'a' && s[a] <= 'z'  {
            i += string(s[a] - 32)
        } else {
            i += string(s[a])
        }
    }
	return i
}

func main(){
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string)string{
	var i string
	for a := 0 ; a <len(s) ; a++{
		if (a == 0 || s[a-1] == ' ' || s[a-1] == '+') && s[a] >= 'a' && s[a] <= 'z' {
			i += string(s[a]-32)
		}else{
			i += string(s[a])
		}
	}
	return i
}

func RUNE(s []rune)string{
	var i string
	for _, a := range s {
		i += string(a)
	}
	return i
}


func BasicJoin(elems []string) string {
	var i string
	for _, a:= range elems{
		i += string(a)
	}
	return i
}

func main(){
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
	for a := 'A' ; a <= 'Z' ; a++{
		fmt.Printf("the number [%d] and the letter is [%c]\n",a,a)
	}
	fmt.Println(RUNE([]rune {77 , 100}))
}

