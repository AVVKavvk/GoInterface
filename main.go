// // package main

// // import (
// // 	"fmt"
// // 	"net/http"
// // )

// // func main() {

// // 	data:= []string{}

// // 	urls:=[]string{
// // 		"https://google.com",
// // 		"https://fb.com",
// // 		"https://vipinnotes.onrender.com",
// // 		"https://instagram.com",
// // 	}

// // 	for _,url := range urls{
// // 		go getRespone(url)
// // 	}

// // 	for _,s:= range data
// // }

// // func getRespone(url string) string{
// // 	resp,err:=http.Get(url)

// // 	if err!=nil{
// // 		fmt.Printf("Somethings went wrong with %s\n",url)
// // 		return ""
// // 	}
// // 	return fmt.Sprintf("%d status code for %s",resp.StatusCode,url)
// // }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	urls := []string{
// 		"https://google.com",
// 		"https://fb.com",
// 		"https://vipinnotes.onrender.com",
// 		"https://instagram.com",
// 	}

// 	// Slice to store results, with an entry for each URL
// 	results := make([]string, len(urls))

// 	// Launch goroutines to fetch responses and store in the results slice
// 	for i, url := range urls {
// 		go func(i int, url string) {
// 			results[i] = getResponse(url)
// 		}(i, url)
// 		wg.Add(1)

// 	}

// 	for _, result := range results {
// 		fmt.Println(result)
// 	}
// 	wg.Wait()
// }

// func getResponse(url string) string {
// 	defer wg.Done()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return fmt.Sprintf("Something went wrong with %s: %v", url, err)
// 	}
// 	defer resp.Body.Close()

// 	return fmt.Sprintf("%d status code for %s", resp.StatusCode, url)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	fmt.Println("Starting.........")

// 	wg:= &sync.WaitGroup{}
// 	mut:= &sync.Mutex{}

// 	wg.Add(3)

// 	result:= []string{"0"}

// 	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
// 			fmt.Println("Go 1")
// 			mut.Lock()
// 			result = append(result, "1")
// 			mut.Unlock()
// 			wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
// 			fmt.Println("Go 2")
// 			mut.Lock()
// 			result = append(result, "2")
// 			mut.Unlock()
// 			wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
// 			fmt.Println("Go 3")
// 			mut.Lock()
// 			result = append(result, "3")
// 			mut.Unlock()
// 			wg.Done()
// 	}(wg, mut)

// 	wg.Wait()

// 	fmt.Println(result)

// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Starting.......")

// 	myCha:= make(chan string,2)

// 	myCha<-"vipin"

// 	fmt.Println(<-myCha)

// }

// package main

// import (
// 	"fmt"

// 	"time"
// )

// func Doing(done <-chan bool) {
// 	for {
// 		select{
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("Doing work...............................")
// 		}
// 	}
// }
// func main() {

// 	fmt.Println("Starting...............")
// 	done:= make(chan bool)

// 	go Doing(done)

// 	time.Sleep(time.Second*5)
// 	close(done)

// }

package main

import (
	"fmt"

	"github.com/AVVKavvk/1/directory"
)

func main() {
	// Interfaces()
  // SomeInterface()
  // SomeMap()
  // file.Dummy()
  directory.Dir()
}

func SomeMap(){

  data:= map[string]interface{}{
    "name": "vipin",
    "mis":112115174,
  }

  // fmt.Println(data["nam"])
  d:="nam"
  val,exits:=data[d];

  if !exits{
    fmt.Println("Doesn't exits")
  }else{
    fmt.Println(val)
  }
}




























func Interfaces()  {
  data:= []interface{}{"vipin","srijan","rythum","synder","sparsh"}

  // fmt.Println(data)
  // ret:= []string{}

  // for _,str:= range data{
  //   ret = append(ret, str.(string))
  // }
  ret:= InterfacesToList(data)
  fmt.Println(ret)
}

func InterfacesToList(inf interface{} ) []string{
  ret:= make([]string,0)

  for _,str:= range inf.([] interface{}){
    ret = append(ret, str.(string))
  }
  return ret
}

func SomeInterface(){

  data:=[]interface{}{

    []string{"vipin","srijan","synder"},
    map[string]string{
      "mis":"112115174",
      "sem":"7",
    },
    map[string]int{
      "pin":332021,
      "number":23,
    },
  }


  for _,event := range data{

    switch val:=event.(type) {
    case []string:
      fmt.Println("[] string")
        for _, name:= range val{
          fmt.Printf("%s ",name)
        }
      fmt.Println("")  
    case map[string]string:
      fmt.Println("map[string]string")
      
      for k,v:= range val{
        fmt.Printf("%s : %s  ",k,v)
      }
      fmt.Println("")
    case map[string]int:
      fmt.Println("map[string]int")
      
      for k,v:=range val{
        fmt.Printf("%s : %d  ",k,v)
      }
      fmt.Println("")
    }
  } 
}