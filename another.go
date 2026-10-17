package main 
import "fmt"



'2' = 50

'3' = 51

'4' = 52

'5' = 53

'6' = 54

'7' = 55

'8' = 56

'9' = 57

#####################################                      ########################################         #####################


func main(){
	fmt.Println(num("000000123"))
	fmt.Println(num("789"))
	fmt.Println(num("3356"))

	fmt.Println(revers("mohamed")) // i takes the first number which is 3 look 3 not '3' if you wonder because var i int i accapts int not anything else 
}// i (i * 10 ) + int(a[s]- '0')  the box (3 * 10 ) + int(51 - 48) = 

func num(a string )int {
	var i int 
	for s := 0 ; s <len(a) ; s++ {
		i = (i * 10) + int(a[s] - '0') // i is in fact 7 so 7 * 10 + 8 - 48  because 8 it turns out that '8' is 56 and '0' is 48 so 56 - 48 
 	} // 7 * 10 + 56 - 48 = 78
	return i // i here concederd a box that safes things or recive things 
}


#####################################                      ########################################         #####################



func main(){
	f := []int {1,2,4,3,5,6,7,9,8}
	f = append(f , 10)
	swap := f[3]
	f[3] = f[2] 
	f[2] = swap
	temp := f[7]
	f[7] = f[8]
	f[8] = temp
	for a := 0 ; a < 3 ; a++ {
		fmt.Println("ahmed")
	}

	fmt.Println(f)
	fmt.Println(swap)
}


#####################################                      ########################################         #####################


func main(){
	a := "mohamed"
	var h string
	for _, s := range a {
		h += string(s)  // h = h + s  0 + m = m , 1 + o = o 
		h = string(s) + h // what is the defrence between them 
 	}
	fmt.Print(h)
}


#####################################                      ########################################         #####################


func main(){
	a := "ahmed"
	var s string
	for f := 0 ; f <len(a) ; f++{
		s = string(a[f]) + s // 0 = "a" + 0 = a 
	} //                     "a" = "h" + "a" = ha
	fmt.Print(s)
}



#####################################                      ########################################         #####################


func main(){
	fmt.Print(reversed("gana"))
}

func reversed(a string)string{
	var t string
	for _, s := range a {
		t = string(s) + t  // 0 = "g" +  0 = "g"
		//                    "g" = "a" + "g" = ag and so on
	}
	return t
}



#####################################                      ########################################         ################


func main(){
	fmt.Print(recursion(3))
}

func recursion(n int) int {
	if n == 0 {
		return 1
	}

	return n * recursion(n - 1)
}


#####################################                      ########################################         ###################

func main(){
	fmt.Print(num(3))  
	fmt.Print("\n")// 5! = 5.4! this 4! referes to 4 * 3 , 3 * 2 and so on the result is 24 and 24 * 5 = 120
	fmt.Print(num(-1))  
} // every thing is clear but how do we applay that 

func num(a int)int{

	if a < 0 {
		fmt.Print("no ")
		return -1
	}
	if a == 0 {
		return 1 
	}
	return a * num(a-1) // 3 == 1 no 2 == 1 no 1== 1 yes so 1 says who is the last number that called me 2 replay saying me so 1 * 2 = 2 and then 
	// 2 * 3 = 6 finlly the result is 6
}



#####################################                      ########################################         ##############################



func main(){
	fmt.Print(num(4,3))  // 4 is the base number 3 is the power of 4 or the exponent 
	// and also 4 is the number that will be multiplied three times 
}

func num(a int,b int)int{
	if b < 0 {
		return 0  // here to avoid any iteration or an errer
	}

	i := 1 
	for s := 1 ; s <=b ; s++{
		i = i * a 
	}
	return i
}

##################################          ################################                 ####################################


func main(){
	fmt.Print(recursive(4,3))  // 1 * 4 = 4 so this is one step 1 from k which is 3 and 4 * 4 = 16 so two from 3 we have one left so 16 * 4
}

func recursive(a int , k int)int{
	if k == 0 {
		return 1 // 3 == 1 no,, so 2 == 1 no ,, so 1 == 1 yes so 1 * 2 , 2 * 3  == 1 * 4  = 4 and than 4 *  4 = 16 but if it ends up at 0 
		// it that will happen 4 * 4 and 3 == 0 no 2==0 no 1==0 no 0==0 return 1 so here where we start 0 look at the last number which is 1
		// so 1 * 1 that means 1 * 4 = 4 and 1 * 2 that means 4 * 4 = 16 and lastly 2 * 3 and this is the last one and that means 16 * 4 = 64 and thats it
	}

	return a * recursive(a , k - 1) 
}



##################################          ################################                 #################################### 



func factorial(n int) int {
   // Base case: If n is 0, return 1.
   if n == 0 { 
	fmt.Print(n)
       return 1
   }
   // Recursive case: Multiply n by the result of factorial(n-1).
   return n * factorial(n-1) 
}


func main(){
	fmt.Printf("the number is [%d]", factorial(4))
}




##################################          ################################                 ####################################




another one 


func funcA(n int) {
    if n > 0 {
        fmt.Println("funcA:", n)
        funcB(n - 1)
    }
}

func funcB(n int) {
    if n > 0 {
        fmt.Println("funcB:", n)
        funcA(n - 1)
    }
}

func main() {
    funcA(4)
}


#####################################                      ########################################         ##########################

func main(){
		fmt.Printf("the factorul of this number is [%d]\n",recursion(7)) // only when i need the factorul of number
		fmt.Printf("the factorul of this number is [%d]",recursion(4))
}

func recursion(a int)int{
	if a < 0 {
		fmt.Print("not allowed",a)
		return 0
	}

	if a == 0 {
		return 1 // so as soon as it reachs 0 it returns 1 and than start multiplaing back up 1 * the last number is 1 so 1 * 1 = 1 and and than 1 * 2 = 2 and so on 
	}

	// here is the must important step 
	return a * recursion(a -1)
}




#####################################                      ########################################         ################################################



if i need to get something like this 4^3 and that means 4 to the power of 3 using recursion no_for loop 

func main(){
	fmt.Printf("5^2 = [%d]",recursion(5,2))
}

func recursion(a int , s int)int{
	if s == 0 {
		return 1 
	}
	if a < 0 {
		return 0
	}

	return a * recursion(a,s-1)
}



#####################################                      ########################################         ################################################


// and the for loop way though to iterate the number like you want 

func main(){
	fmt.Print(num(5,2))
}

func num(a int , s int)int{
	if s < 0 {
		return 0
	}

	result := 1
	for f := 1 ; f <= s ; f++{
		result = result * a 
	}
	return result
}


#####################################                      ########################################         ############################

func main() {
	fmt.Println(name("madam"))
	fmt.Println(name("ahemd"))

}

func name(a string)string{
	var result string
	for _, s := range a{
		result = string(s) + result 
	}
	return result
}


#####################################                      ########################################         ###########################


func main(){
	fmt.Println(num(8,2))
	fmt.Println(num(4,4))
}

func num(a int , s int)int{
	if s == 0 {
		return 1
		//         8 * 2 means 8 * the result of the first one not 2 
	} //           8 * 1 after the return = 8 
	return a * num(a , s - 1) // it's 8 at first and num (8 , s - 1)  2 == 0 , 1 == 0 , 0 == 0 so it returns 1 so (8 , 1) 8 * 1 = 8 and lastly 
	// (8 , 2) 8 * 8 = 64 so num = 64

	// with 4 ,, 4 * num(4 , s - 1) ,, 4 == 0 ,, 3 == 0 ,, 2 == 0 ,, 1 == 0 ,, 0 == 0 so it returns 1 
	// (4 , 1) , (4 , 2) , (4 , 3) , (4 ,4) so 4 * 1 = 4 so (4 , 2) 4 * 4 = 16 ,, (4 , 3) 4 * 16 = 64 ,, (4 * 4) 4 * 64 = 256
}

#####################################                      ########################################         ######################


func main(){
	fmt.Print(num(2,4))
}
func num(j int , l int)int {
	if l == 0 {
		return 1
	}

	if j < 0 {
		return 0
	}

	return j * num(j , l - 1) // 1 * 2 = 2 ,, 2 * 2 = 4 ,, 2 * 4 = 8,, 2 * 8 = 16 
}



#####################################                      ########################################         ##########################



func main(){
	fmt.Print(num(4))
}

func num(s int)int{
	f := 1 
	for a := 1 ; a <= s ; a++{
		f = f * a
	}
	return f
}