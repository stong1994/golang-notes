package main

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	select {
	case <-ch1:
	case <-ch2:
	default:

	}
}

//  0x00f4 00244 (many_send_case.go:6)      CALL    runtime.selectgo(SB)
