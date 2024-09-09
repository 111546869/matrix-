# 文件说明

## todo-list

todo-list文件夹中的Golang代码实现的功能是待办列表的创建和增删改查

### 使用方法

1. 到达todo-list目录下，输入以下命令，初始化一个go项目并利用 `go mod`来安装和管理所需的依赖：

   ```bash
   #先要进入到wsl环境中
   cd path/to/todo-list
   go mod init todo-list
   go mod tidy
   ```

2. 运行代码： 

   ```bash
   go run main.go
   ```

   运行完之后会出现以下界面

   ```
   [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
   
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] POST   /users/register           --> todo-list/controllers.Register (3 handlers)
   [GIN-debug] POST   /users/login              --> todo-list/controllers.Login (3 handlers)
   [GIN-debug] GET    /todos/                   --> todo-list/controllers.GetTodos (4 handlers)
   [GIN-debug] POST   /todos/                   --> todo-list/controllers.CreateTodo (4 handlers)
   [GIN-debug] PUT    /todos/:id                --> todo-list/controllers.UpdateTodo (4 handlers)
   [GIN-debug] DELETE /todos/:id                --> todo-list/controllers.DeleteTodo (4 handlers)
   [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
   Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
   [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
   [GIN-debug] Listening and serving HTTP on :8080
   ```

   

3. 新开一个终端，利用本地电脑进行验证：

   ```bash
   #注册
   curl -X POST http://localhost:8080/users/register -d '{"username": "testuser", "password": "password"}'
   #登录
   curl -X POST http://localhost:8080/users/login -d '{"username": "testuser", "password": "password"}'
   #增加todo任务
   curl -X POST http://localhost:8080/todos -H "Authorization: Bearer <your_token>" -d '{"title": "Test Todo", "description": "Test description"}'
   #查询所有 Todo（GET 请求）
   curl -X GET http://localhost:8080/todos -H "Authorization: Bearer <your_token>"
   #查询指定 ID 的 Todo（GET 请求）
   curl -X GET http://localhost:8080/todos/<id> -H "Authorization: Bearer <your_token>"
    #删除指定 ID 的 Todo（DELETE 请求）
   curl -X DELETE http://localhost:8080/todos/<id> -H "Authorization: Bearer <your_token>"
   
   
   ```

   

### 代码原理

1. `main.go`

先初始化一个默认路由r来处理http请求，比如GET,POST,PUT,DELETE，然后与数据库建立连接并把数据格式传给数据库，以便用户上传的数据能按照设定的格式持久性存储到数据库中。

然后就是通过利用自己写好的函数接口和数据结构处理用户上传的数据，protected中定义的GET,POST,PUT,DELETE方法只有经过JWT Token验证才能使用。

最后再运行即可。

2. `jwt.go`和 `auth_middlewares.go`

   `jwt.go`主要是分为生成token和解析token，生成token部分是利用某种算法生成一个时间期限为72小时的token，解析token主要是检查用户提交的token的合法性

   `auth_middlewares.go`则是调用`jwt.go`中的函数接口来处理用户上传的token

3. `todo.go`和`user.go`

   这两个文件是设计待办事项和用户的信息

4. `todo_controller.go`和`user_controller.go`

   `todo_controller.go`实现对于用户的增删改查的请求，主要是与数据库进行交互，更改数据库中存储的内容

   `user_controller.go`实现用户的注册和登录功能，其中为了保护用户密码的隐私，我们存储的是哈希加密后的密钥

5. `database.go`

   主要是初始化自己的数据库并建立连接，返回建立的数据库

## docker_learning

### 5

利用nginx.conf配置两个server，分别处理a.site和b.site两个域名的请求

使用方法：

```bash
cd path/to/5
docker build -t 5:latest .
docker run -d -p 80:80 --name my-nginx 5:latest
curl http://a.site
curl http://b.site
```

### 6，7

利用docker-compose部署网页访问计数和Redis存储服务，使用方法：

```
cd path/to/6
docker-compose up
#然后访问提示的网址即可
```

## go_learning

主要是学习interface` , `defer` , `chan`, `go` ,`select这些关键字的用法以及参考 《build web application with golang》第三章搭建简单的web服务
