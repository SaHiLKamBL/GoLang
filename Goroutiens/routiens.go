package main

import (
	"fmt"
	// "time"
)

// func download(file string){
// 	  fmt.Println("Start downloading:", file)
//     time.Sleep(2 * time.Second)
//     fmt.Println("Finished downloading:", file)
// }

// func main() {
//     go download("File A")
//     go download("File B")

//     time.Sleep(3 * time.Second) // wait for goroutines
// }


//  Problem : in main func we explicitly mention the main func to wait but what when we dont know how much time routiens required to complete



import "sync"

func download(file string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Downloading:", file)
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2)

    go download("File A", &wg)
    go download("File B", &wg)

    wg.Wait() // waits until both are done
}


// wg.Add(2) → expecting 2 tasks
// Each goroutine → wg.Done()
// wg.Wait() → blocks until all done

// 👉 No guessing, perfect sync ✅