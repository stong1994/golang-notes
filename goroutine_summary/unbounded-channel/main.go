package main

func MakeInfinite() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})

	inQueue := []interface{}{}
	outCh := func() chan<- interface{} {
		if len(inQueue) == 0 {
			return nil
		}
		return out
	}
	curVal := func() interface{} {
		if len(inQueue) == 0 {
			return nil
		}
		return inQueue[0]
	}
	go func() {
		for len(inQueue) > 0 || in != nil {
			select {
			case v, ok := <-in:
				if !ok {
					in = nil // in通道关闭后，将in设置为nil，select语句就不会再次尝试这个分支
				} else {
					inQueue = append(inQueue, v)
				}
			case outCh() <- curVal():
				inQueue = inQueue[1:]
			}
		}
		close(out)
	}()
	return in, out
}
