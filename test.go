package main 
import "fmt"



package main 
import "fmt"

// strings ,  int  , runes 

// range  and len  with for

// the next func or the copy one 

// return 

//                   the first 

func main(){
	name := []string {"ramadan ahemd"}
	fmt.Printf("this is your name %v",name)
}

//                   the socend 

func main(){
	names := []string {"ahmed salem", "omar mohamed","ali ahmed"}
	fmt.Printf("the first name is : [%v] \n",names[0])
	fmt.Printf("the last one is: [%v]\n",names[2])
	for a := 0 ; a <len(names) ; a++ {
		fmt.Print(a)
	}
}


//                   the third one and the best for me

func name()(n,m,h string){
	n = "ahmed salem"
	m = "ali emad"
	h = "motaaz"
	return 
}


func main (){
	n,m,h := name()
	fmt.Printf("the first name is: [%v] , the socend is: [%v] and the third one is: [%v]",n,m,h)
}





//                                  // the first and the best one for me 


func numbers()(i,a,o int){
	return 18 , 25 , 30 
}

func main(){
	i,a,o := numbers()
	fmt.Printf("number 1 is: [%d] , number 2 is: [%d] and number 3 is: [%d]",i,a,o)
}




//                                  using for either with len or range


//                                   the first let's say range 

func main(){
	a := []string {"jana","kamel","tamer"}
	for index,thing := range a {
		fmt.Printf("the index: [%d]  and  the thing: [%v]\n",index,thing)
	}
}




//                                   this time with len 



func main(){
	a := []string {"jana","kamel","tamer"}
	for i := 0 ; i <len(a) ; i++ {
		fmt.Print(a[i],"\n")
	}
}


//                                  another way with len 

//                   this will print these numbers many times depending on how many elements inside the slise 


func main(){
	a := []int {122,3,4,57,23}
	for i := 0 ; i <len(a) ; i++{
		fmt.Println(a)
	}
}


//                            with range how to know how many number inside the slise or letter



func main(){
	number := []int {1,3,2,5,4}
	temp := number[2]
	number[2] = number[1]
	number[1] = temp
	fmt.Print(number)

}

func num(a string)(i int){
	for u := 0 ; u <len(a) ; u++{  
	    i = (i * 10) + int(a[u] - '0')
	}
	return 
}

func main(){
	fmt.Println(num("589"))  // no need to create varaibles because i don't need to change something in func main or even make another condition 
	fmt.Println(num("00000001234")) // related to the apporation in return 
	fmt.Println(num("00000")) // if you want just to to something else add for example f := num(and the thing)
	
}


// another way to do that 


func main(){
	a := "789"
	var t int // defining t so its = 0 now 
	for s := 0 ; s <len(a) ; s++ {  // this appoartion to loop through a := "789"
		t = (t * 10) + int(a[s] - '0') // here is what's all about t hold the result and * 10 this must be with digits look down and see 
	} // + int to modiefy the string and make it int a[s] is '7' ,'8' and '9' so '7' - '0' means 55 - 48 it's hidding thing in the cumpeuter ok 
	fmt.Print(t) // '0' prabably works with everything i think that i'm not sure 
}



'0' = 48

'1' = 49

'2' = 50

'3' = 51

'4' = 52

'5' = 53

'6' = 54

'7' = 55

'8' = 56

'9' = 57
