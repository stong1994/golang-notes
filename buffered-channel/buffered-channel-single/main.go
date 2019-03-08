package main

func main() {

}

type Evaluator func(interface{}) (interface{}, error)

func DivideAndConquer(data interface{}, evaluators []Evaluator) ([]interface{}, []error) {
	gather := make(chan interface{}, len(evaluators))
	errors := make(chan error, len(evaluators))
	for _, v := range evaluators {
		go func(e Evaluator) {
			result, err := e(data)
			if err != nil {
				errors <- err
			} else {
				gather <- result
			}
		}(v)
	}
	out := make([]interface{}, 0, len(evaluators))
	errs := make([]error, 0, len(evaluators))
	// 循环len(evaluators)次以接收到所有的值。否则的话，需要增加一个计数器来监控（如sync.WaitGroup）
	for range evaluators {
		select {
		case r := <-gather:
			out = append(out, r)
		case e := <-errors:
			errs = append(errs, e)
		}
	}
	return out, errs
}
