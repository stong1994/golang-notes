# RESTful与CORS
> 现在的WEB项目经常使用RESTful风格的Api，那随之而来的就是跨域问题，找了几篇好的博客，在这边记录下（写一写总是有好处的，原文在最下边）
*PS: 阮一峰老师写的博客真的是精炼，像我这么懒的人，几乎都复制了过来。*
## RESTful API 的总结
> REST，即Representational State Transfer的缩写。我对这个词组的翻译是"表现层状态转化"

> URI：Uniform Resource Identifier，统一资源标识符  
  URL：Uniform Resource Location，统一资源定位符
  
#### RESTful架构：  
（1）每一个URI代表一种资源；  
（2）客户端和服务器之间，传递这种资源的某种表现层；  
（3）客户端通过HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"。

#### 五个常用的HTTP动词：
- GET（SELECT）：从服务器取出资源（一项或多项）。
- POST（CREATE）：在服务器新建一个资源。
- PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。
- PATCH（UPDATE）：在服务器更新资源（客户端提供改变的属性）。
- DELETE（DELETE）：从服务器删除资源。

#### 两个不常用的动词：
- HEAD：获取资源的元数据。与GET相同，只是不返回报文主体部分，用于确认URI的有效性及资源更新的日期时间等。
- OPTIONS：用来查询针对请求URI指定的资源支持的方法


**关于PATCH**: 
> 一个提升互通性并且阻止错误产生的新方法是必要的。 PUT已经定义为用一个全新的body来覆盖资源，并且不能在局部改变时重复使用。另一方面，代理、缓存甚至客户端和服务端可能会
对得到的结果感到困惑。POST已经被使用，但是没有广泛的互通性。早期的HTTP提到了PATCH但是没有，但是没有定义。

> A new method is necessary to improve interoperability and prevent
     errors.  The PUT method is already defined to overwrite a resource
     with a complete new body, and cannot be reused to do partial changes.
     Otherwise, proxies and caches, and even clients and servers, may get
     confused as to the result of the operation.  POST is already used but
     without broad interoperability (for one, there is no standard way to
     discover patch format support).  PATCH was mentioned in earlier HTTP
     specifications, but not completely defined.

> PATCH方法是在请求URI中定义一些描述改变的东西，如果请求URI没有指向一个存在的资源，这个服务有可能会创建一个新的依赖patch文档类型资源和权限

> The PATCH method requests that a set of changes described in the
     request entity be applied to the resource identified by the Request-
     URI.  The set of changes is represented in a format called a "patch
     document" identified by a media type.  If the Request-URI does not
     point to an existing resource, the server MAY create a new resource,
     depending on the patch document type (whether it can logically modify
     a null resource) and permissions, etc.
     
> PUT和PATCH请求之间的区别在于服务器处理在URI中封闭的实体来更改资源。在PUT请求中，封闭的实体被认为是存储在
源服务器上的资源的修改版本，客户端请求替换存储的版本。在PATCH请求中，封闭的实体包含了一系列指令来描述如何修改
源服务器上存储的资源来生成新的版本。PATCH方法依赖URI来影响资源，可能会对其他资源产生副作用。**应用可以通过PATCH来
创建新资源，或者修改现有资源**

>  The difference between the PUT and PATCH requests is reflected in the
     way the server processes the enclosed entity to modify the resource
     identified by the Request-URI.  In a PUT request, the enclosed entity
     is considered to be a modified version of the resource stored on the
     origin server, and the client is requesting that the stored version
     be replaced.  With PATCH, however, the enclosed entity contains a set
     of instructions describing how a resource currently residing on the
     origin server should be modified to produce a new version.  The PATCH
     method affects the resource identified by the Request-URI, and it
     also MAY have side effects on other resources; i.e., new resources
     may be created, or existing ones modified, by the application of a
     PATCH.
     
> PATCH请求可以以幂等的方式来发出，能够避免两个PATCH在同一时间请求同一个资源带来的风险。

#### 常用的状态码
- 200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
- 201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
- 202 Accepted - [*]：表示一个请求已经进入后台排队（异步任务）
- 204 NO CONTENT - [DELETE]：用户删除数据成功。
- 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
- 401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
- 403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
- 404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
- 406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
- 410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
- 422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。
- 500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。

## 跨域资源共享 CORS 详解
> CORS是一个W3C标准，全称是"跨域资源共享"（Cross-origin resource sharing）。  
> 它允许浏览器向跨源服务器，发出XMLHttpRequest请求，从而克服了AJAX只能同源使用的限制。  

#### 简单请求  
（1）请求方法是以下三种方法之一
- HEAD
- GET
- POST

（2）HTTP的头信息不超出以下几种字段：
- Accept
- Accept-Language
- Content-Language
- Last-Event-ID
- Content-Type：只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain

*如果不能满足以上需求，就是非简单请求。*

对于简单请求，浏览器直接发出CORS请求。具体来说，就是在头部信息中，增加一个**Origin**字段。
**Origin**表明了请求信息（如：Origin: http://api.bob.com），服务器根据这个值来判断是否同意该请求。

如果**Origin**指定的域名**能够被访问**，服务器返回的响应中，会多出几个头部信息
```
Access-Control-Allow-Origin: http://api.bob.com  
Access-Control-Allow-Credentials: true  
Access-Control-Expose-Headers: FooBar  
Content-Type: text/html; charset=utf-8
```
注意上边四个头部信息中，有三个以`Access-Control`开头。
- Access-Control-Allow-Origin：该字段是**必须**的。它的值要么是请求时Origin字段的值，要么是一个*，表示接受任意域名的请求。
- Access-Control-Allow-Credentials：该字段**可选**。它的值是一个布尔值，表示**是否允许发送Cookie**。**默认情况下，Cookie不包括在CORS请求之中**。**设为true，即表示服务器明确许可**，Cookie可以包含在请求中，一起发给服务器。这个值也**只能设为true**，如果**服务器不要浏览器发送Cookie，删除该字段即可**。
- Access-Control-Expose-Headers:该字段**可选**。CORS请求时，**XMLHttpRequest**对象的**getResponseHeader()**方法只能拿到6个基本字段：`Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma`。如果想拿到其他字段，就必须在Access-Control-Expose-Headers里面指定。上面的例子指定，getResponseHeader('FooBar')可以返回FooBar字段的值。

如果**Origin**指定的域名**不能够被访问**，服务器会返回一个正常的HTTP响应，浏览器发现，头部信息中没有包含**Access-Control-Allow-Origin**。
于是我们就知道出错了，从而抛出一个错误，被**XMLHttpRequest**的**onerror**回调函数捕获。（HTTP的响应状态码可能为200，因此不能通过状态码来识别该错误）

**withCredentials** 属性

CORS请求默认不发送Cookie和HTTP认证信息。如果要把Cookie发到服务器，一方面要服务器同意，指定Access-Control-Allow-Credentials字段。

另一方面，开发者必须在AJAX请求中打开`withCredentials`属性。
```
var xhr = new XMLHttpRequest();
xhr.withCredentials = true;
```
需要注意的是，如果要发送Cookie，Access-Control-Allow-Origin就**不能设为星号**，必须指定明确的、与请求网页一致的域名。

#### 非简单请求  
不符合上述简单请求，即为非简单请求，比如请求方法是`PUT`或`DELETE`，或者`Content-Type`字段的类型是`application/json`等。

非简单请求的CORS请求，会在正式通信之前，增加一次HTTP查询请求，称为**"预检"请求**（preflight）。

浏览器先**询问服务器**，当前网页所在的域名是否在服务器的**许可名单**之中，以及可以使用哪些**HTTP动词**和**头信息字段**。只有得到肯定答复，浏览器才会发出正式的`XMLHttpRequest`请求，否则就报错。

如下边的浏览器JS脚本
```javascript
var url = 'http://api.alice.com/cors';
var xhr = new XMLHttpRequest();
xhr.open('PUT', url, true);
xhr.setRequestHeader('X-Custom-Header', 'value');
xhr.send();
```
上面代码中，HTTP请求的方法是`PUT`，并且发送一个自定义头信息`X-Custom-Header`。

浏览器发现，这是一个非简单请求，就自动发出一个"预检"请求，要求服务器确认可以这样请求。下面是这个"预检"请求的HTTP头信息。
```
OPTIONS /cors HTTP/1.1
Origin: http://api.bob.com
Access-Control-Request-Method: PUT
Access-Control-Request-Headers: X-Custom-Header
Host: api.alice.com
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```
注意到请求方法是`OPTIONS`,`Origin`中的值表示**请求来自哪个源**。`Access-Control-Request-Method`和`Access-Control-Request-Headers`中表示“非简单”的访问方法和头部信息

- Access-Control-Request-Method

该字段是**必须**的，用来列出浏览器的CORS请求会用到哪些**HTTP方法**，上例是PUT。

- Access-Control-Request-Headers

该字段是一个**逗号分隔的字符串**，指定浏览器CORS请求会**额外发送的头信息字段**，上例是X-Custom-Header。

服务器收到"预检"请求以后，检查了`Origin、Access-Control-Request-Method`和`Access-Control-Request-Headers`字段以后，**确认允许跨源请求**，就可以做出回应。

```
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:15:39 GMT
Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http://api.bob.com
Access-Control-Allow-Methods: GET, POST, PUT
Access-Control-Allow-Headers: X-Custom-Header
Content-Type: text/html; charset=utf-8
Content-Encoding: gzip
Content-Length: 0
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
Content-Type: text/plain
```

上面的HTTP回应中，关键的是`Access-Control-Allow-Origin`字段，表示http://api.bob.com可以请求数据。该字段也可以设为**星号**，表示**同意任意跨源请求**。

如果浏览器**否定**了"预检"请求，会返回一个**正常的HTTP回应**，但是没有任何CORS相关的头信息字段。这时，浏览器就会认定，服务器不同意预检请求，因此**触发一个错误**，被`XMLHttpRequest`对象的`onerror`回调函数捕获。控制台会打印出如下的报错信息。

```
XMLHttpRequest cannot load http://api.alice.com.
Origin http://api.bob.com is not allowed by Access-Control-Allow-Origin.
```

服务器回应的其他CORS相关字段如下。

```
Access-Control-Allow-Methods: GET, POST, PUT
Access-Control-Allow-Headers: X-Custom-Header
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 1728000
```
- Access-Control-Allow-Methods

该字段**必需**，它的值是逗号分隔的一个字符串，表明服务器支持的所有跨域请求的方法。注意，返回的是所有支持的方法，而不单是浏览器请求的那个方法。这是为了避免多次"预检"请求。

- Access-Control-Allow-Headers

如果浏览器请求包括Access-Control-Request-Headers字段，则Access-Control-Allow-Headers字段是必需的。它也是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段。

- Access-Control-Allow-Credentials

该字段与简单请求时的含义相同。

- Access-Control-Max-Age

该字段可选，用来指定**本次预检请求的有效期**，单位为秒。上面结果中，有效期是20天（1728000秒），即允许缓存该条回应1728000秒（即20天），在此期间，不用发出另一条预检请求。

一旦服务器**通过了"预检"请求**，以后**每次**浏览器正常的CORS请求，就都跟简单请求一样，会有一个`Origin`头信息字段。服务器的回应，也都会有一个`Access-Control-Allow-Origin`头信息字段。

下面是"预检"请求之后，浏览器的正常CORS请求。
```
PUT /cors HTTP/1.1
Origin: http://api.bob.com
Host: api.alice.com
X-Custom-Header: value
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```
上面头信息的`Origin`字段是浏览器**自动添加**的。

下面是服务器正常的回应。
```
Access-Control-Allow-Origin: http://api.bob.com
Content-Type: text/html; charset=utf-8
```
上面头信息中，`Access-Control-Allow-Origin`字段是每次回应都**必定包含**的。

### 参考资料  
- [PATCH的定义](https://tools.ietf.org/html/rfc5789)
- [深入理解RESTful API](https://www.jianshu.com/p/84568e364ee8)
- [RESTful API 设计指南](http://www.ruanyifeng.com/blog/2014/05/restful_api.html)
- [跨域资源共享 CORS 详解](http://www.ruanyifeng.com/blog/2016/04/cors.html)
- [不要再问我跨域的问题了](https://segmentfault.com/a/1190000015597029)
- [请求路由-压缩字典树](https://github.com/chai2010/advanced-go-programming-book/blob/master/ch5-web/ch5-02-router.md)