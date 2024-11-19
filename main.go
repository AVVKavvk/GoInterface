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

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Interfaces()
//   // SomeInterface()
//   // SomeMap()
//   // file.Dummy()
//   // directory.Dir()
// }

// package main

// import "fmt"

// func worker(id int, jobs <-chan int, results chan<- int) {
//   for j := range jobs {
//       fmt.Println("Worker", id, "processing job", j)
//       results <- j * 2
//   }
// }

// func main() {
//   jobs := make(chan int, 5)
//   results := make(chan int, 5)

//   for w := 1; w <= 3; w++ {
//       go worker(w, jobs, results)
//   }

//   for j := 1; j <= 5; j++ {
//       jobs <- j
//   }
//   close(jobs)

//   for a := 1; a <= 5; a++ {
//       fmt.Println(<-results)
//   }
// }

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// func generator(nums ...int) <-chan int {
//   out := make(chan int)
//   go func() {
//       for _, n := range nums {
//           out <- n
//       }
//       close(out)
//   }()
//   return out
// }

// func square(in <-chan int) <-chan int {
//   out := make(chan int)
//   go func() {
//       for n := range in {
//           out <- n * n
//       }
//       close(out)
//   }()
//   return out
// }

// func main() {
//   in := generator(2, 3, 4)
//   out := square(in)

//   for n := range out {
//       fmt.Println(n)
//   }
// }
// func main() {

//   ch:=make(chan string)

//   wg:= sync.WaitGroup{}

//   wg.Add(2)

//   go func() {
//     fmt.Println("entering first")
//     time.Sleep(time.Second*5)
//     ch<- "vipin"
//     wg.Done()
//   }()

//   // val:=<-ch
//   // fmt.Println(val)

//   go func() {
//     fmt.Println("second")
//     fmt.Println(<-ch)
//     wg.Done()
//   }()

// wg.Wait()

// }


func  doing(done <-chan bool)  {
  for {
    select {
    case <-done:
      fmt.Println("byeeeeeeeeee")
      return
    default:
      fmt.Println("doing something")
      
    }
  }
}
// func main()  {

//   done:= make(chan bool)
//   go doing(done)

//   time.Sleep(time.Millisecond*2)
  
//   fmt.Println("middle")
//   close(done)

//   time.Sleep(time.Millisecond*2)
// }


// func main()  {
  
//   ch:=make(chan string)
//   wg:= sync.WaitGroup{}

//   wg.Add(2)

//   go func() {
//     fmt.Println("starting first....")
//     ch<- "vipin"
//     fmt.Println("added data to first....")
//     wg.Done()
//   }()

//   time.Sleep(time.Millisecond*1)

//   go func() {
//     fmt.Println("Starting second")
//     time.Sleep(time.Second*2)
//     fmt.Println("Reciver ready")
//     fmt.Println("getting val from frist", <-ch)
//     wg.Done()
//   }()

//   wg.Wait()
// }



// func main() {
// 	ch := make(chan interface{})
// 	stringCh := make(chan string)
// 	wg := sync.WaitGroup{}

// 	wg.Add(2)

// 	go func() {
// 		fmt.Println("Starting to add squared values to ch")
// 		defer wg.Done()
// 		defer close(ch)
// 		for i := 1; i < 8; i++ {
// 			ch <- i * i
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		defer close(stringCh)
// 		for val := range ch {
// 			if intVal, ok := val.(int); ok {
// 				stringCh <- strconv.Itoa(intVal)
// 			}
// 		}
// 	}()

// 	go func() {
// 		for stringVal := range stringCh {
// 			fmt.Printf(" %s ", stringVal)
// 		}
// 	}()

// 	wg.Wait()
// }

func sliceToChannel(nums []int) <-chan int{
  ch:= make(chan int)

  go func() {
    for _,val:= range nums{
      ch<-val
      fmt.Printf("Adding value to ch %d\n",val)
    }
    close(ch)
  }()
  return ch
}

func sq(ch <-chan int) <-chan int{
  data:= make(chan int)

  go func() {
    
    for val := range ch{
      data<-val*val
      fmt.Printf("conveting val to sq %d\n",val*val)
    }
    close(data)
  }()
  return data
}

// func main(){
//   nums:= [] int{1,2,3,4,5,6}

//   channel:= sliceToChannel(nums)

//   result:= sq(channel)

//   time.Sleep(time.Second*2)
//   for n := range result{
//     fmt.Printf("Printing value %d \n",n)
//   }
  
// }


func addData(ch chan int, wg *sync.WaitGroup)  chan int{
  
  fmt.Println("Started adding values")
  
 go func() {
  for i:=1;i<12;i++{
    time.Sleep(time.Second*1)
    fmt.Println("adding ",i)
    ch<-i
  }
  defer wg.Done()
  close(ch)
 }()

  return ch
}

func fecthData (ch chan int, wg *sync.WaitGroup){
  fmt.Println("Started fetching values")

  go func() {
    for val := range ch{
      fmt.Println("fetching ",val)
    }
    defer wg.Done()
  }()
}

func getStringToInt(s string) int{
  i,err:= strconv.Atoi(s)
  if err!=nil{
    return -1
  }
  return i 
}

func main(){  

  strs:=make([]string,0)
  strs = append(strs, "12","13","14","15","16","17","18","#","20","21","22","23","24","25","26","27","28","29","30")

  wg:=sync.WaitGroup{}
  mut:=sync.Mutex{}

  wg.Add(len(strs))

  result:=[]int{}

  for _,val:= range strs{
    go func() {
      defer wg.Done()
      i:=getStringToInt(val)
      mut.Lock()
      result = append(result, i)
      mut.Unlock()
    }()
  }

  wg.Wait()

  for _,val := range result{
    fmt.Println(val)

  }
}
// func main(){
  
//   ch:= make(chan int)

//   wg:= &sync.WaitGroup{}

//   wg.Add(2)

//   ch= addData(ch,wg)

//   fecthData(ch,wg)

//   wg.Wait()
  

// }









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