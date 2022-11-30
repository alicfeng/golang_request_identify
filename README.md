<div align="center">
    <h1>
        <a href="https://devops.coding.smoa.cloud/p/smore-vimo/d/golang_request_identify/git">
            Distributed Request Identity
        </a>
    </h1>
    <p>
       分布式请求链路追踪组件
    </p>
</div>



## 🪤 组件简介

该组件为 `Golang` 的开箱即用 `SDK`，是一款分布式请求链路追踪插件，基于请求上下文传播机制实现，提供请求链路追踪能力。



## ✨ 支持特性

- [x] 开箱即用，基于中间件注册方式使用

- [x] 支持 `kratos` 框架服务端与客户端中间件



## 🔨 快速安装

基于 `go` 包管理器

```shell
go get e.coding.smoa.cloud:devops/smore-vimo/golang_request_identify@v1.0.0
```



## 🚀 快速使用

##### 基于 `kratos` 框架引擎

因 `kratos` 的 `rpc and http` 协议均基于 `metadata` 作为请求载体，因此需要提前启用 `metadata` 模块

- 服务端注册

  ```
  var opts = []http.ServerOption{
      http.Middleware(
          // ... ...
          metadata.Server(),
          kratos_request_identify.Server(),
          // ... ...
      ),
  }
  ```

- 客户端注册

  ```
  conn, err := grpc.DialInsecure(
      context.Background(),
      grpc.WithEndpoint("127.0.0.1:10000"),
      grpc.WithMiddleware(
          // ... ...
          metadata.Client(),
          kratos_request_identify.Client(),
          // ... ...
      ),
  )
  ```



##### 获取请求标识值

根据 `context.Context` 即可获取

```
request_identify.GetRequestIdentify(ctx)
```



## 📒 请求验证

 ```shell
 curl -I -X GET 127.0.0.1:8000/helloworld/samego
 HTTP/1.1 200 OK
 Content-Type: application/json
 X-Md-Global-Request-Identity: e64b0044e2e476edd1a0d68586387f8a
 Date: Wed, 30 Nov 2022 10:51:39 GMT
 Content-Length: 26
 ```

响应头上会增加一个头部属性 `X-Md-Global-Request-Identity`，此值就是分布式传递的标识符，值由 `UUID4` 生成且转小写构建而来
