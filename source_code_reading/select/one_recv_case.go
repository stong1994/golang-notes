package main

func main() {
	ch1 := make(chan struct{})
	select {
	case <-ch1:

	}
}

// 0x0054 00084 (one_recv_case.go:6)       CALL    runtime.chanrecv1(SB)
