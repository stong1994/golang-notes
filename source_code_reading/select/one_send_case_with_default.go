package main

func main() {
	ch := make(chan struct{})
	select {
	case ch <- struct{}{}:
	default:

	}
}

// 0x0050 00080 (one_send_case.go:6)       CALL    runtime.selectnbsend(SB)
