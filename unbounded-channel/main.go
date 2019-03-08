package main

func MakeInfinite() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})

	output := []interface{}{}
	outCh := func() chan<- interface{} {
		if len(output) == 0 {
			return nil
		}
		return out
	}
	curVal := func() interface{} {
		if len(output) == 0 {
			return nil
		}
		return output[0]
	}
	go func() {
		for len(output) > 0 || in != nil {
			select {
			case v, ok := <-in:
				if !ok {
					in = nil // in通道关闭后，将in设置为nil，select语句就不会再次尝试这个分支
				} else {
					output = append(output, v)
				}
			case outCh() <- curVal():
				if len(output) > 0 {
					output = output[1:]
				}
			}
		}
		close(out)
	}()
	return in, out
}
