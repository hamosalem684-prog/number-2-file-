package main 
import "fmt"







func main(){
	names := []string {"ahmed","omar","mohamed","salem"}
	sort.Strings(names)
	fmt.Println(names)
	i := sort.SearchStrings(names,"ahmed")
	fmt.Print("the position of this ",i)
}




//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$    //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$     //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$





func main(){
	f := 425
	b := 10
	j,p :=count(f,b)
	fmt.Printf("this is what j has = %d and p though %d ",j,p)
}

func count(a,h int)(j,p int){
	j = a % h   // modulo so a modulo h 
	p = a / h   // using slash or a divided by h 
	return  // reminder this word is important multiplied by this one *
}



//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$    //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$     //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$






func divmod(a *int , b *int)(g int , i int) {
	return *a % *b , *a / *b           // this by using return and multeble varaible or apporations two varaible becouse two appoartions 
}

func main(){
	f := 13
	h := 2
	g , i := divmod(&f,&h) // I'm calling return ok 
	fmt.Printf("this is the remander i think [%d] , but this is the slash [%d]",g,i)
}




//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$    //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$     //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$





func divmod(a,b int , div *int , mod *int ){
	*div , *mod = a / b , a % b  // modifay things through pointers 
}
func main(){
	f := 13
	m := 2
	var div int // its nesessary putting or definding these otherwise i wouldn't print anything 
	var mod int 
	divmod(f,m, &div , &mod)
	fmt.Printf("div = [%d] and mod [%d] ",div , mod )
}




//$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$    //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$     //$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$





func copy(a string) (s string){
	for g := len(a)-1 ; g >= 0 ; g-- { // this new thing len(a or anythings)-1 to get the last litter or number but i think no need to do it with numbers 
		s += string(a[g]) // s += to move on antill i reaches the last litter this is not like printing directorly
	}
	return
}

func main(){
	f := "mohamed"
	s := copy(f) // calling return and what it has 
	fmt.Print(s)
}