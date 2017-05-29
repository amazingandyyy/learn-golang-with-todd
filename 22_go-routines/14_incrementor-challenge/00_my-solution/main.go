package main

import (
	"fmt"
	"strconv"
)

func main() {
	var count int64
	ch := incrementor(2)
	for c := range ch {
		count++
		fmt.Println(c)
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(s int) chan string {
	chn := make(chan string)
	ok := make(chan bool)
	for i := 0; i < s; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				chn <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", j)
			}
			ok <- true
		}(i)
	}

	go func() {
		for i := 0; i < s; i++ {
			<-ok
		}
		close(chn)
	}()

	return chn

}

/*
CHALLENGE #1
-- Take the code from above and change it to use
channels instead of wait groups and atomicity
*/
