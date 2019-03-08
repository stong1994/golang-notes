## 设计两个goroutine，一个打印a,c,e....，一个打印b,d,f...结果输出为a,b,c,d....z

* main方法为面试时写的，较笨重
* a-z-func2_test.go是回来后写的，简化后的。总的来说，用for循环来检测channel状态要比for-select-case 要简化很多