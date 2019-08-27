报错信息  
`segmentation fault`  
原因：
#### 第一个：编译问题
有些代码编译的时候会根据系统信息编译成对应的版本，因此将centos编译好的文件放到ubuntu上执行会报错  
####第二个：文件传输问题
将centos虚拟机编译好的文件拖动到本地（windows10）,然后再拖动到ubuntu，发现文件大小变了，因此执行报错。改用`scp`传输文件后，解决问题。
感谢[文章](https://studygolang.com/topics/1733)解惑。

由上述第二个问题，学到了一个新的命令来查看编译后的文件信息：readelf

执行命令
`readelf -h filename`  
查看结果，发现报错。
```go
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00 
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x459910
  Start of program headers:          64 (bytes into file)
  Start of section headers:          624 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         10
  Size of section headers:           64 (bytes)
  Number of section headers:         37
  Section header string table index: 9
readelf: linux_client: Error: Reading 0x1ff bytes extends past end of file for string table
readelf: linux_client: Error: Reading 0x378 bytes extends past end of file for symbols
readelf: linux_client: Error: the dynamic segment offset + size exceeds the size of the file
readelf: linux_client: Error: no .dynamic section in the dynamic segment

```
