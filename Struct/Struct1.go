 
 //simple way
 package main
 import ("fmt")

 type Uni struct {
	Domain string
	A
	B
	D
}

type A struct {
	Uname string
}

type B struct {
	roll int
}
type D struct{
	regNo int
}

func main() {
	var myInfo Uni = Uni{
		Domain: "CSE",
		A:     A{Uname: "kamran"},
		B:     B{roll: 6},
		D:	   D{regNo: 12217803},

	}
 
	fmt.Println(myInfo.Domain,myInfo.Uname,myInfo.roll,myInfo.regNo) 
}
 