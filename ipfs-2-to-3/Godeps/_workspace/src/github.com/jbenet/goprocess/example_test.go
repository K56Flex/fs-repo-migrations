package goprocess_test

import (
	"fmt"
	"time"

	"github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/goprocess"
)

func ExampleGo() {
	p := goprocess.Go(func(p goprocess.Process) {
		ticker := time.Tick(200 * time.Millisecond)
		for {
			select {
			case <-ticker:
				fmt.Println("tick")
			case <-p.Closing():
				fmt.Println("closing")
				return
			}
		}
	})

	<-time.After(1100 * time.Millisecond)
	p.Close()
	fmt.Println("closed")
	<-time.After(100 * time.Millisecond)

	// Output:
	// tick
	// tick
	// tick
	// tick
	// tick
	// closing
	// closed
}
