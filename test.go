package main
import (
	"fmt"
	"strings"
)


// those are shortcuts

func main(){
	myname := "hamo salem"
	fmt.Println(strings.Contains(myname,"hamo")) // its clear from the word it makes sure is this thing exict or not and it gives you result eather true or fulse at the end
	fmt.Println(strings.Contains(myname,"omar"))
	fmt.Println(strings.ReplaceAll(myname , "salem", "abdelghany"))  // the same though but this use it to change a word or something and don't forget its a temporary thing
	fmt.Println(strings.ToUpper(myname)) // just makes everything C like this  (HAMO SALEM)
	fmt.Print("the main one ",myname)   // so here is the bottom line these shortcuts do not change forever just for a certin time in the moment
}



//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$


func names(b int , u int)(int,int){ // b and u are used to recieve a copy from the main function ok 
    b = 20  // doing this is not going to change anything in the main func so i does't affect anything its just a capy
	u = 10 // the same 
	return b + u , b - u  // this apporation takes either the copy or if i used pointer that's a diffrent thing 
}                         // using pointers could change what inside func main easly ok 

func main(){
	j := 10 // this is the main number not the copy one 
	k := 5
	n,g := names(j,k) // n and g those are related to return whearas names(j,k) are going to pass the first func and make a copy
	fmt.Printf("this is total %d  and this is the minse %d ",n , g) 
	fmt.Printf("j = %d and k = %d",j,k)
}

//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$

func str(n string){
}

func main(){

	fmt.Print(str("hello world")) // I'm taking this word to the first func as a copy 
}
//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$



func names(p,b int )(total int,mines int) {   // i can also write it like this names(p,b int ) not names(p int, b int )
	return p + b , p - b // two apporation so two (total int,mines int) got it ?
}

func main(){
	fmt.Print(names(10,5)) 
}
