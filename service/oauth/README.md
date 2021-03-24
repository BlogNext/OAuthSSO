# oauth功能模块说明


- 只提供预授权码一种模式颁发AccessToken



## 简单实现方案

1. 只建oauth_client表
2. 预授权码存存在缓存中（内存）,lru,失败后，重新走流程
3. preAuthCode、AccessToken，RefreshToken都不做存储，用jwt做，都给客户端保存，客户端丢失，重新授权走流程
4. 目前只开放用户登录资源,开放的资源是用户信息，返回用户的id和昵称

