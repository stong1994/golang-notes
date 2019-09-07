package main

func main() {
	ch1 := make(chan struct{})
	select {
	case <-ch1:
	default:

	}
}

//  0x004f 00079 (one_recv_case_with_default.go:6)  CALL    runtime.selectnbrecv(SB)
