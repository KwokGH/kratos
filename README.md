# Kratos项目实践

##开发流程 
1. 定义接口  
kratos的接口统一使用protobuf定义的，无论是走grpc协议还是走http协议。
把接口的定义和具体的协议隔离出来，并用统一的定义语言定义接口入参和回参以及其他的信息。
kratos提供了http协议的protoc插件，在生成grpc的代码时，也会生成http的代码。
定义http接口的方式使用的是google的规范。
```bash
    kratos proto add api/account/account.proto
```
2. 使用``` make api ``` 生成相关pb代码
3. 设计数据库表（如果需要的话）
4. 添加相关配置（如果需要的话）
   1. 例如jwt token 密钥信息的配置，修改internal/conf/conf.proto添加相应的配置
   2. 修改Bootstrap对象，增加对应的配置项,并且在yaml文件中新增对应的节点。
   3. 执行```make config```生成pb文件
5. 在biz层定义接口，编写逻辑代码
6. 在data层实现biz定影的接口
