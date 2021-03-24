# service结构目录


>> 功能模块名

>>> model 层与数据库表一一对应

>>> service 层提供服务(不会对数据做任何的验证处理，到service的数据都是对的)

>>> eneity (与用户数据之间的数据传输)

>> controller.go 控制器（做数据验证，服务调用、构建服务调用需要的数据）

> exception 自定义的运行时异常

> main.go 入口函数

