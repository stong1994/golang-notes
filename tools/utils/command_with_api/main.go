package main

import (
	"bufio"
	cRand "crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	mRand "math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	h bool
	p string
	c string
	k string
	t int // t秒后仍未执行完，强制退出
)

func init() {
	flag.BoolVar(&h, "h", false, "help info")
	flag.StringVar(&p, "p", "8080", "listen port")
	flag.StringVar(&c, "c", "echo program started!", "command")
	flag.StringVar(&k, "k", genRandStr(), "use to check permission")
	flag.IntVar(&t, "t", 4, "after t second, exit if program is not finish")

	flag.Usage = usage // 覆盖默认函数
}

func usage() {
	fmt.Fprintf(os.Stderr, `command exec api
Usage: command [-p port] [-c command] [-k key] [-t limit_time]
	-h: 帮助命令
	-p: 端口号，默认8080
	-c: 启动时执行的命令
	-k: 设定的秘钥key，用于api访问时权限校验
	-t: 限定的命令执行时间，如ping命令能够一直执行，为了防止这种情况，增加t来限制时间
`)
	//flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if h { // 即使访问时，h没有参数，默认为true，但是如果没有使用命令h，则为设置的默认值false
		flag.Usage()
		return
	}
	Init()
	addr := fmt.Sprintf(":%s", p)
	fmt.Printf("listening %s ;key is %s\n", addr, k)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func Init() {
	_, err := Exec(c)
	if err != nil {
		panic(err)
	}
	route()
}

func parseCommand(c string) []string {
	return strings.Split(c, " ")
}

func Exec(c string) ([]byte, error) {
	var (
		data []byte
	)
	commands := parseCommand(c)
	if len(commands) < 1 {
		return nil, fmt.Errorf("command can not be empty")
	}
	var cmd *exec.Cmd
	if len(commands) == 0 {
		cmd = exec.Command(commands[0])
	} else {
		cmd = exec.Command(commands[0], commands[1:]...)
	}
	done := make(chan struct{}, 1)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(stdout)

	go func() {
		time.Sleep(time.Second * time.Duration(t))
		close(done)
	}()
	//实时循环读取输出流中的一行内容
	for {
		select {
		case <-done:
			goto end
		default:
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				goto end
			}
			fmt.Println(line)
			data = append(data, []byte(line)...)
		}

	}
	cmd.Wait() // TODO 这里不调用会不会有资源泄露的问题， 代码中没有说会有资源泄露问题，只是等待执行完毕
end:

	if len(data) == 0 {
		data = []byte("exec success")
	}
	return data, nil
}

func route() {
	http.HandleFunc("/command", func(writer http.ResponseWriter, request *http.Request) {
		var (
			result []byte
		)
		err := request.ParseForm()
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}
		c := request.Form.Get("c")
		key := request.Form.Get("key")
		if key != k {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("key is not valid"))
			return
		}

		result, err = Exec(c)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
		writer.Write(result)
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf(":%s/command?c={command}&key={key}", p)))
	})
}

func genRandStr() string {
	var allNum = "23456789QWERTYUPASDFGHJKLZXCVBNMqwertyupasdfghjkzxcvbnm"
	l := len(allNum)
	var data []byte
	for i := 0; i < 10; i++ {
		r := getRand(int64(l))
		data = append(data, allNum[r])
	}
	return string(data)
}

func getRand(m int64) int64 {
	n, err := cRand.Int(cRand.Reader, big.NewInt(m))
	if err != nil {
		mRand.Seed(time.Now().UnixNano())
		return mRand.Int63n(m)
	}
	return n.Int64()
}
