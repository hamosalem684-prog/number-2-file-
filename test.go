package main 
import  "fmt"






// using (len)

// three ways though 
// 1 = the number one way is that printint the varaible directly that will print the thing sevrel times
// 2 = printting the varaible that belongs to the for loop like a down here that will print the index or the position of every litter in the word
// 3 = printting the two varaible the one that belongs to j and the for loop like j[a] that will print the runes 


func main(){
	j := "mohamed salem"   // here is the name or the thing you need
	for a := 0 ; a <len(j);a++ {  // the process using for to pass thing or git into them 
		fmt.Println(j) // here is the bottom line the result either you choose j or a or j[a]
	}
}



func main(){
	j:= "hamo"
	for a := 0 ; a <len(j) ;a++{
		fmt.Println(a)
	}
}



func main(){
	j:= "hamo"
	for a := 0 ; a <len(j) ;a++{
		fmt.Println(j[a])
	}
}



//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$	//$$$$$$$$$$$$$$$$$$$ //$$$$$$$$$$$$$$$$$$$




// here using range the keyword 

// these are the ways we use for and range  three ways and the fact is that's what i know now ok 



for () , () := range {  // the first one is using two i'll call them things 
	fmt.Print(...)    // this first is prabably to give you the index or the position of any thing inside the varaible
}                // the socend one for runes ok that means i passes the loop and take each charactar and of course as numbers 



for ():= range .. {  // this is the first so that will print runes not the position of the word or whatever 
	fmt.Print(...)  
}



for range ..{ // this is like using len but at the end printing the varaible and that makes it print the thing sevrel time 
	fmt.Print(...)
}