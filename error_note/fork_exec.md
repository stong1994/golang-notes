报错信息：
```
fork/exec The process cannot access the file because it is being used by another process.
```
一开始以为当前程序创建的脚本文件，当程序还在运行时，就占有该脚本的控制权，因此不能在程序运行时执行该脚本。

后来觉得不对，当文件`Close()`之后，程序就失去了控制权。  

经检查， 果然忘记`Close()`了。

然后想重现这个bug，发现即使没有`Close()`，也不会再报这个错。。。

有可能是其他原因？

错误总是这样，当你不想得到的时候，她总来烦你，当你想得到的时候她就不来了。

有缘再见吧。