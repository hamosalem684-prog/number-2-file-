~package main
import (
	"fmt"
)
// "github.com/01-edu/z01"

// ctrl + r to change the file directly from hamo-go to anything else


func main(){
	fmt.Println(number("12345"))
	fmt.Println(number("str123ing45"))
	fmt.Println(number("012 345"))
	fmt.Println(number("Hello World!"))
	fmt.Println(number("sd+x1fa2W3s4"))
	fmt.Println(number("sd-x1fa2W3s4"))
	fmt.Println(number("sdx1-fa2W3s4"))
	fmt.Println(number("sdx1+fa2W3s4"))
	fmt.Println(name("ahmed omar mohamed salem"))
}

func number(s string)int{
	var i int // the box
	o := 1
	Nonagtive := false // this used to check things but here I'm not doing that I'm multiplaing 1 and -1 with the fanill result 
	for _, a:= range s {
		if a >= '0' && a <='9'{
			i = i * 10 + int(a - '0') // this gives me every number whether positiov or nagtive 
			Nonagtive = true // what this does ? 
		}else if a == '-' && !Nonagtive { // and this !Nonagitve  what it does 
			o = -1
		}
	}
	return i * o 
}


func name(s string)string{
	var i string
	for a := 0 ; a <len(s) ; a++{
		if (a == 0 || s[a-1] == ' ') && (s[a] >= 'a' && s[a] <= 'z') {
			i += string(s[a] - 32)
		}else {
			i += string(s[a])
		}
	}
	return i
}


#################################################################


func main(){
	f := []int {1,2,3,4,5,6}
	even := false
	for _, a := range f{
		if a % 2 == 0 {
			even = true 
			fmt.Println(a)
		}
	}
	if even {
		fmt.Println("this is even number")
	}else{
		fmt.Println("this is odd number")
	}
}


################################################################

finding wheter the thing is true or false 

func main(){
	str := name([]string {"ahmed", "omar", "abdo"})
	if str {
		fmt.Println("found")
	}else{
		fmt.Println("not found")
	}
}
func name(s []string)bool{
	for _, a := range s{
		if a == "ahmed" {
			return true
		}
	}
	return false
}

########################################################################

an unxcpeted error like this 

func main(){
	numbers, err := num(20) // now it says num returns two values why ?
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numbers) // what this would call ? if i have it up here 
}

func num(s int)(string,error){
	if s < 18 {
		return "",errors.New("we accept at least someone who's 18 ") // tell me what is the binfets of putting this "" tell me agein with good example i wanna git it 
	}
	return "you can have it for sure ", nil // what nil means here what it does 
}

#####################################################################

with nagetive numbers 

func main(){
	thenumber,err := funC(-1)
	if err != nil { // this means if this is wrong if the value is not what i need or isn't vaild send me this error
		fmt.Println(err) // and this error massage 
	}else{
		fmt.Println(thenumber)// this is the succes massage ok 
	}
}

func funC(s int)(string,error){
	if s < 0 {
		return "",errors.New("no nagetive numbers")
	}
	return "ok" ,nil
}

###########################################################################  yo

func main(){
	num(6791198465345)
}
func num(s int){
	var i []int
	for s > 0 {
		i = append(i , s%10)
		s /= 10
	} 

	for h := 0 ; h < len(i) ; h++{
		for t := h + 1 ; t < len(i) ; t++{
			if i[h] > i[t] {
				i[h] , i[t] = i[t] , i[h]
			}
		}
	}
	for _, o := range i {
		fmt.Print(o)
	}
}

