报错信息
```
/usr/local/go/pkg/tool/linux_amd64/link: running gcc failed: exit status 1
/usr/bin/ld: i386 architecture of input file `/tmp/go-link-180479662/000000.o' is incompatible with i386:x86-64 output
collect2: error: ld returned 1 exit status
```
第一感觉是GCC或者go的版本问题，google到[答案](https://github.com/golang/go/issues/12448#issuecomment-137279343)： 
`Ah, I figure out. I put syso file in the directory. It contains windows resource objects.
 Sorry.`  
 
原来是为了获取windows下的管理员权限，产生了一个`syso`文件，删掉这个文件就可以在linux上进行`build`了
