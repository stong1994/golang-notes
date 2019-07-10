# RESTful API 的总结

五个常用的HTTP动词：
- GET（SELECT）：从服务器取出资源（一项或多项）。
- POST（CREATE）：在服务器新建一个资源。
- PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。
- PATCH（UPDATE）：在服务器更新资源（客户端提供改变的属性）。
- DELETE（DELETE）：从服务器删除资源。

两个不常用的动词：
- HEAD：获取资源的元数据。
- OPTIONS：获取信息，关于资源的哪些属性是客户端可以改变的。


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

### 常用的状态码
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

**参考资料**：  
- [PATCH的定义](https://tools.ietf.org/html/rfc5789)
- [深入理解RESTful API](https://www.jianshu.com/p/84568e364ee8)
- [RESTful API 设计指南](http://www.ruanyifeng.com/blog/2014/05/restful_api.html)