package main

import (
	"bytes"
	"fmt"
)

type Buffer bytes.Buffer

var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		// Grab a buffer if available; allocate if not.
		select {
		case b = <-freeList:
			fmt.Println("client get one")
			// Got one; nothing more to do.
		default:
			// None free, so allocate a new one.
			fmt.Println("client create new one")
			b = new(Buffer)
		}
		load(b)         // Read next message from the net.
		serverChan <- b // Send to server.
	}
}

func load(buffer *Buffer) {
	fmt.Println("client load a buffer")
}

func server() {
	for {
		b := <-serverChan // Wait for work.
		process(b)
		// Reuse buffer if there's room.
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default: // TODO for + select default 需小心使用，避免CPU频繁调用
			// Free list full, just carry on.
		}
	}
}

func process(buffer *Buffer) {
	fmt.Println("server process a buffer")
}
