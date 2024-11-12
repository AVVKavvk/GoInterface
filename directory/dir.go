package directory

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Dir()  {
	
	dirpath:="/home/vipin/Go/1/by_code"

	// err:= os.MkdirAll(path.Join(dirpath,"1")   ,0777)
	// err= os.MkdirAll(path.Join(dirpath,"2")   ,0777)
	// // err:= os.MkdirAll(path,0777)
	// if err!=nil{
	// 	fmt.Println("Somethings went wrong")
	// }

	files,err:= ioutil.ReadDir(dirpath)
	if err!=nil{
		fmt.Println("Something went wrong")
	}else{
		for _, file := range files {
			if strings.Contains(file.Name(),"vip"){
				fmt.Println(file.Name())
			}
		}
	}

}