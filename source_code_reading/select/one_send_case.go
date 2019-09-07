package main

func main() {
	ch := make(chan struct{})
	select {
	case ch <- struct{}{}:
	}
}

// 0x0055 00085 (one_send_case.go:6)       CALL    runtime.chansend1(SB)
