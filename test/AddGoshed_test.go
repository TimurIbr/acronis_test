package test

import (
	"../AddGoshed"
	"github.com/stretchr/testify/require"
	"testing"
)
/*The test of 1 for-loop*/
func TestAddGoschedSimpleFile(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
"runtime"
)
func main() {
	for i := 1; i < 10; i+=1 {
		println("Hello, World!") 
	}
}
`
	required_answer := `package main

import (
	"runtime"
)

func main() {
	for i := 1; i < 10; i += 1 {
		println("Hello, World!")
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	require.True(t, required_answer == answer)
}

/*The test of for-loop in for-loop*/
func TestAddGosched2level(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
	"runtime"
)
func main() {
	for i := 1; i < 10; i+=1 {
		for i := 1; i < 10; i+=1 {
			println("level2!") 
		}
		println("level1") 
		for i := 1; i < 10; i+=1 {
			println("level2") 
		}
	}
}
`
	required_answer := `package main

import (
	"runtime"
)

func main() {
	for i := 1; i < 10; i += 1 {
		for i := 1; i < 10; i += 1 {
			println("level2!")
			runtime.Gosched()
		}
		println("level1")
		for i := 1; i < 10; i += 1 {
			println("level2")
			runtime.Gosched()
		}
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}

/*The test of for-loop in declared function*/
func TestAddGoschedFuncDecl(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
	"runtime"
)
func foo(){
	for i := 1; i < 10; i+=1 {
			println("foo:level1") 
		}
}

func main() {
	for i := 1; i < 10; i+=1 {
		foo()
		println("level1") 
		foo()
	}
}
`
	required_answer := `package main

import (
	"runtime"
)

func foo() {
	for i := 1; i < 10; i += 1 {
		println("foo:level1")
		runtime.Gosched()
	}
}

func main() {
	for i := 1; i < 10; i += 1 {
		foo()
		println("level1")
		foo()
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}

/*The test of for-loop in function-object*/
func TestAddGoschedFuncObj(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
	"runtime"
)

func main() {
foo := func(i int) int{
		for i := 1;; 1=1 {
			i += 1
			println("foo:level1") 
		}
		return i + 1
	}
	for i := 1; i < 10; i+=1 {
		foo(i)
		println("level1") 
		foo(i)
	}
}
`
	required_answer := `package main

import (
	"runtime"
)

func main() {
	foo := func(i int) int {
		for i := 1; ; 1 = 1 {
			i += 1
			println("foo:level1")
			runtime.Gosched()
		}
		return i + 1
	}
	for i := 1; i < 10; i += 1 {
		foo(i)
		println("level1")
		foo(i)
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}

/*The test of for-loop with range insteaod of condition*/
func TestAddGoschedRange(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
	"runtime"
)

func main() {
	for i:= range []int{0,1,2}{
		println("level1")
	}
}
`
	required_answer := `package main

import (
	"runtime"
)

func main() {
	for i := range []int{0, 1, 2} {
		println("level1")
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}

/*The test of for-loop in func and func-jbj*/
func TestAddGoschedComplexFile(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main
import (
	"runtime"
)
func tem(){
	for i := 1; i < 10; i+=1 {
		println("Hello, innerfuncWorld!")
		for j := 1; i < 10; i+=1 {
			for j := 1; i < 10; i+=1 {
				println("Hello,SUPER DUPER MEGA innerfuncWorld!")
			}
			println("Hello,SUPER innerfuncWorld!")
			for j := 1; i < 10; i+=1 {
				println("Hello,SUPER DUPER innerfuncWorld!")
			}
		}
	}
}
func main() {
	for i := 1; i < 10; i+=1 {
		println("Hello, World!") // reaully hello
	}
	a := func(i int) int{
		for i := 1;; 1=1 {
			i += 1
			tem()
		}
		return i + 1
	}
	for i := 1;; 1=1 {
		i += 1
		a(i)
	}

}

`

	required_answer := `package main

import (
	"runtime"
)

func tem() {
	for i := 1; i < 10; i += 1 {
		println("Hello, innerfuncWorld!")
		for j := 1; i < 10; i += 1 {
			for j := 1; i < 10; i += 1 {
				println("Hello,SUPER DUPER MEGA innerfuncWorld!")
				runtime.Gosched()
			}
			println("Hello,SUPER innerfuncWorld!")
			for j := 1; i < 10; i += 1 {
				println("Hello,SUPER DUPER innerfuncWorld!")
				runtime.Gosched()
			}
			runtime.Gosched()
		}
		runtime.Gosched()
	}
}
func main() {
	for i := 1; i < 10; i += 1 {
		println("Hello, World!")
		runtime.Gosched()
	}
	a := func(i int) int {
		for i := 1; ; 1 = 1 {
			i += 1
			tem()
			runtime.Gosched()
		}
		return i + 1
	}
	for i := 1; ; 1 = 1 {
		i += 1
		a(i)
		runtime.Gosched()
	}

}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}

/*The test of adding runtime import*/
func TestAddGoschedRuntimeInsert(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i+=1 {
		fmt.Print("Hello, World!") 
	}
}
`
	required_answer := `package main

import "runtime"

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i += 1 {
		fmt.Print("Hello, World!")
		runtime.Gosched()
	}
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	print(answer)
	require.True(t, required_answer == answer)
}

/*The test of for-loop in multi tread programm*/
func TestAddGoschedComplex2File(t *testing.T) {
	// t.Skip()
	t.Parallel()
	src := `
package main

import "fmt"
import "time"

func main() {
    
    numcpu := runtime.NumCPU()
    fmt.Println("NumCPU", numcpu)
    //runtime.GOMAXPROCS(numcpu)
    runtime.GOMAXPROCS(1)
    
	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < 1000000; i++ {
			ch1 <- i
		}
		ch1 <- -1
		ch2 <- 0.0
	}()
	go func() {
          total := 0.0
		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < 100000; i++ {
				m := <-ch1
				if m == -1 {
					ch2 <- total
				}
			}
			t2 := time.Now().UnixNano()
			dt := float64(t2 - t1) / 1000000.0
			total += dt
			fmt.Println(dt)
		}
	}()
	
	fmt.Println("Total:", <-ch2, <-ch2)
}`
	required_answer := `package main

import "runtime"

import "fmt"
import "time"

func main() {

	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)

	runtime.GOMAXPROCS(1)

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < 1000000; i++ {
			ch1 <- i
			runtime.Gosched()
		}
		ch1 <- -1
		ch2 <- 0.0
	}()
	go func() {
		total := 0.0
		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < 100000; i++ {
				m := <-ch1
				if m == -1 {
					ch2 <- total
				}
				runtime.Gosched()
			}
			t2 := time.Now().UnixNano()
			dt := float64(t2-t1) / 1000000.0
			total += dt
			fmt.Println(dt)
			runtime.Gosched()
		}
	}()

	fmt.Println("Total:", <-ch2, <-ch2)
}
`
	answer := AddGoshed.AddGoschedToFile("SimpleFile", src)
	//print(answer)
	require.True(t, required_answer == answer)
}