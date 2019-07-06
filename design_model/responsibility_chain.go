package design_model

/**
责任链模式
1. 定义
	在责任链模式里，很多对象由每一个对象对其下家的引用而连接起来形成一条链。请求在这个链上传递，直到链上的某一个对象决定处理此请求。
2. 使用步骤
3. 使用场景
	对一个对象进行"流水线"处理
4. 优点
	降低耦合度。它将请求的发送者和接收者解耦。
	简化了对象。使得对象不需要知道链的结构。
	增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。
	增加新的请求处理类很方便。
5. 缺点
	不能保证请求一定被接收。
	系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。
	可能不容易观察运行时的特征，有碍于除错。
参考文章: https://blog.csdn.net/lkysyzxz/article/details/79548853
 */

 /*
 模拟http的请求处理
  */

type Values map[string][]string

// request
type Request struct {
	Method string
	Url string
	PostForm Values
}

func NewRequest() *Request {
	r := Request{}
	r.PostForm = make(Values)
	return &r
}

func SendRequest(request *Request, mux *Mux)  {
	mux.root.Process(request)
}

func (r *Request) SetValues(key string, values ...string) {
	r.PostForm[key] = append(r.PostForm[key], values...)
}

func (r *Request) Post(url string, mux *Mux) {
	r.Method = "POST"
	r.Url = url
	SendRequest(r, mux)
}

// 处理器接口
type IProcess interface {
	Process(request *Request)
}

type HandlersCollection map[string]IProcess

// 处理器基类
type Processer struct {
	Handlers HandlersCollection
}

func (p *Processer) SetHandler(key string, process IProcess) {
	p.Handlers[key] = process
}

func (p *Processer) Init() {
	p.Handlers = make(HandlersCollection)
}

// 根处理器
type RootProcesser struct {
	Processer
}

func (p *RootProcesser) Process(request *Request) {
	p.Handlers[request.Method].Process(request)
}

func NewRootProcesser() *RootProcesser {
	r := &RootProcesser{}
	r.Init()
	return r
}

// POST请求处理器
type PostProcesser struct {
	Processer
	PMux *Mux
}

func NewPostProcesser(mux *Mux) *PostProcesser {
	post := new(PostProcesser)
	post.Processer.Init()
	post.PMux = mux
	return post
}

func (pp *PostProcesser) Process(request *Request) {
	pp.PMux.mux[request.Url](request)
}

type HandlerFunc func(request *Request)
type muxEntry map[string]HandlerFunc // 存放url和对应的处理请求的方法

// 多路复用器
type Mux struct {
	mux muxEntry
	root IProcess
}

func (m *Mux) Handle(url string, handlerFunc HandlerFunc) {
	m.mux[url] = handlerFunc
}

func (m *Mux) SetRootProcess(root IProcess) {
	m.root = root
}

func NewMux() *Mux {
	mux := Mux{}
	mux.mux = make(muxEntry)

	root := NewRootProcesser()
	post := NewPostProcesser(&mux)
	root.SetHandler("POST", post)

	mux.root = root
	return &mux
}