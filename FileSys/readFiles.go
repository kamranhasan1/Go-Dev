package main
import ("fmt"
		"os")

func main(){

	f,err := os.Open("example.txt")
	if err != nil {
		 panic(err)
	}
	defer f.Close()
	//basic operatons on file 
	
	fileInfo,err:= 	f.Stat()
	if err != nil{
		panic(err)
	}
	fmt.Println("filesize: ",fileInfo.Size())
	fmt.Println("file Name: ",fileInfo.Name())
	fmt.Println("file Or folder: ", fileInfo.IsDir())
	fmt.Println("file Permission: ",fileInfo.Mode())
	fmt.Println("last Modification ",fileInfo.ModTime())

	// Read the file content using Read + buffer
	buf := make([]byte,fileInfo.Size())

	d,err :=f.Read(buf)
	if err != nil{
		panic(err )
	}

	for i := 0; i<len(buf); i++{
 		fmt.Println(d,"content in file: ",string(buf[i]))
	}





	// Or simply using os.ReadFile
	data,err := os.ReadFile("example.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println("simply reading entire file", string(data))


	//read folder 
	dir,err := os.Open("../")
	if err != nil{
		panic(err)
	}
	defer dir.Close()
	fil,err:=dir.ReadDir(-1)
	for  _,fi:= range fil{
		fmt.Println("file in folder",fi.Name(),fi.IsDir())
	}
}
