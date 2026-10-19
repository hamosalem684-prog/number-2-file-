package main 
import "fmt"

func main(){
	fmt.Printf("this is the factorul [%d]\n",recursion(5))
	fmt.Printf("and this is 10 to the power of 2:[%d]\n",recursion2(10,2))
	fmt.Printf("this is the reversed name [%v]\n",reversed("ahmed"))
	fmt.Printf("this is the converting number [%d]\n",number("789"))
}

func recursion(t int)int{ // this is how we git the factorul by using fot loop 
	i := 1 
	for a := 1 ; a <= t ; a++{
		i = i * a 
	}
	return i
}


func recursion2(s int , o int)int{ // this way without for and not only that this way is much bitter 
	if o == 0 {
		return 1
	}

	if s < 0 {
		return 0
	}

	return s * recursion2(s , o-1)
}



func reversed(p string)string{ // reversion 
	var h string
	for _, n := range p {
		h = string(n) + h // this is what's all about
	}
	return h
}




func number(k string)int{ // making text numbers int
	// "789"
	var m int
	for t := 0 ; t <len(k) ; t++{
		m = (m * 10) + int(k[t] - '0') // 7 * 10 + '8' - '0' = 78 and so on untill it becomes 789 
	}
	return m

}

//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                        #########################################



// 1- what is array and what is slise 

func main(){
	// this is what is called array with fixed numbers
	num := [3]string {"ahemd","omat","mohamed"}	
	fmt.Print(num)
}



func main(){
	// this is what called slise with no fixed numbers
	num := []string {"hema","hanan","nour"}
 fmt.Print(num)
}



// the socened way of doing this fibnacci

func main(){
	fmt.Print(fibnacci(4))
}

func fibnacci(h int)int{
	if h <= 1 {
		return h
	}

	f , b := (0,1)

	for a := 2 ; a <= n ; a++{
		f,b = b , f + b // here is a litter bit diffrent is starts from 2 because for did you see that so 
		// 2 has 1 and 0 and here this part start f,b = b , f + b ,, as long as 2 has 1 and 0 so f = 0 and b = 1 
		// b , f + b = 1 this 1 b takes it and f takes 1 that was with b
 		// 1 , 0 + 1 = 1 
		// and than 3 starts 
		// 3 = b ,f + b = 1 + 1 = 2 
		// 4 = b ,f + b = 2 + 1 = 3
		// the bottem line is f takes what b has and b receives the sum of f + b
	}	
	return b 
}


//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                        #########################################


func sayGoodby(s []string ,f func(string)){ // this is a new way I'v just discaverd and usefull though passing s in another function 
	for _, v := range s{
		f(v)
	}
}

func cycle(s string){
	fmt.Printf("goodbay [%v]\n",s)
}


func sayBay(s string){
	
	fmt.Printf("just bay %v\n",s)
	
}
func main(){
	sayGoodby([]string {"ahmed","omar","eman"},cycle)
	sayGoodby([]string {"ahmed","omar","eman"},sayBay)
	
}


//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                        #########################################


func numbers(s []int){ // a new way it turns out that you can call slises this way 
	fmt.Println(s)
}

func call(a []int){
	fmt.Println(a)
}

func main(){
	numbers([]int {1,2,3,4,5})
	call([]int {6,7,8,9,10})
}


//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                        #########################################

 


func main(){
	name("hamo",name2)
}

func name(s string , f func(string)){
	f(s)
}

func name2(s string){
	fmt.Print("hi ",s)
}


//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                        #########################################


type things struct {
	name,city string
	age int
}

func main(){
	a := things{name:"ahmed" ,city:"usa" ,age: 18} 
	fmt.Printf("hi my name is %v I'm from %v and I'm %d years old by the way\n",a.name,a.city,a.age)
	var t func(int)int
	t = func(b int)int{
		if b < 0 {
			return -1
		}

	    i := 1
		for a := 1 ; a<=b ; a++{
			i = i * a
		}
		return i
	}
	fmt.Println(t(4))
	fmt.Println(f_or(6)) // 0 1 1 2 3 5 8 13
	slises([]string {"ahmed","omar","mohamed"}) //!!
}

func f_or(i int )int{
	if i < 0 {
		return -1
	}

	g , h := 0,1

	for a := 2 ; a <= i ; a++{
		g , h = h , g + h 
	}
	return h
}


func f_or2(a int)int{
	if a == 0 {
		return 0
	}

	if a == 1 {
		return 1
	}

	return f_or2(a-1) + f_or2(a-2) 
}

func slises(a []string){
	fmt.Print(a)
} 