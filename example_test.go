// This example demonstrates how to use a cmap container.
package cmap_test

import (
	"fmt"
	"log"

	"github.com/kchristidis/cmap"
)

// This example creates a cmap, fills it to capacity, then adds one more
// key/value pair to demonstrate that the LRU key has been evicted.
func Example() {
	cm, err := cmap.New(2) // will hold up to 2 key/value pairs
	if err != nil {
		log.Fatal(err)
	}

	cm.Put("fooKey", "fooVal")
	v, ok := cm.Get("fooKey") // retrieve the value
	fmt.Println(ok)
	fmt.Println(v)

	cm.Put("barKey", "barVal") // retrieve value as above
	cm.Put("bazKey", "bazVal") // at this point "fooKey" is evicted

	_, ok = cm.Get("fooKey")
	fmt.Println(ok) // expected output: false
}
