package main

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	select {
	case <-ch1:
	case <-ch2:

	}
}

//  0x00d8 00216 (many_recv_case.go:6)      CALL    runtime.selectgo(SB)
