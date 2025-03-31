 
 //simple way
 package main
 import ("fmt")

 type Uni struct {
	Domain string
	A
	B
}

type A struct {
	Uname string
}

type B struct {
	roll int
}

func main() {
	var myInfo Uni = Uni{
		Domain: "CSE",
		A:     A{Uname: "kamran"},
		B:     B{roll: 6},

	}
 
	fmt.Println(myInfo.Domain,myInfo.Uname,myInfo.roll) 
}
 