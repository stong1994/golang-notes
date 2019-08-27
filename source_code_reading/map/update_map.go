package main

func update() {
	var simpleMap = map[string]int{"a": 1}
	simpleMap["a"] = 2
}

/*
0x00a7 00167 (update_map.go:4)  CALL    runtime.mapassign_faststr(SB)
0x00e2 00226 (update_map.go:5)  CALL    runtime.mapassign_faststr(SB)
*/
