# service结构目录


>> 功能模块名

>>> model 层与数据库表一一对应

>>> service 层提供服务(不会对数据做任何的验证处理，到service的数据都是对的)

>>> entity (定义接口的request和response),request可以实现validator.v10的数据验证功能

>> controller.go 控制器（做数据验证，服务调用、构建服务调用需要的数据）

> exception 自定义的运行时异常

> main.go 入口函数,路由入口

> help 一些帮助的方法,一些快捷方法

> config 配置文件

> common 放一些自己写的类库

>> cache 

>>> lru lru算法实现的内存缓存

>>>> lru_sync.go 线程安全的lru

>>>> lru.go 线程不安全的lru(未实现)

>> db 统一的db连接

>>> mysql.go gorm官网连接mysql

>> config 统一的配置类库

