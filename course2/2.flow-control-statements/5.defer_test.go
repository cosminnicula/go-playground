package main

// defers the execution of the current function until the surrounding function returns
// defer fmt.Println("world")
// fmt.Println("hello")

// deferred functions are stacked and executed LIFO
// fmt.Println("counting")
// for i := 0; i < 10; i++ {
// 	defer fmt.Println(i)
// }
// fmt.Println("done")
