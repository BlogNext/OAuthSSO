# oauth功能模块说明


- 只提供预授权码一种模式颁发AccessToken



## 简单实现方案

1. 只建oauth_client表
2. 预授权码存存在缓存中（内存）,lru,失败后，重新走流程
3. preAuthCode、AccessToken，RefreshToken都不做存储，用jwt做，都给客户端保存，客户端丢失，重新授权走流程
4. 目前只开放用户登录资源,开放的资源是用户信息，返回用户的id和昵称



## 接口

### 创建预授权码

```html
POST /api/oauth/create_pre_auth_code

```

请求参数

|  属性   | 类型  | 默认值  | 必填 | 说明 |
|  ----  | ---- |  ----  |----  |----  |
| nickname  | string |  | 是 | 登录昵称 |
| password  | string |  | 是 | 登录密码 |
| client_id  | string |  | 是 | 应用client_id |
| redirect_url  | string |  | 是 | 成功获得预授权码后重定向地址 |

响应

|  属性   | 类型  | 说明 |
|  ----  | ----  |----  |
| pre_auth_code  | string | 预授权码 |
| redirect_url  | string | 成功获得预授权码后重定向地址 |


### pre_auth_code换取access_token

```html
POST /api/oauth/pre_auth_code_access_token

```

请求参数

|  属性   | 类型  | 默认值  | 必填 | 说明 |
|  ----  | ---- |  ----  |----  |----  |
| client_id  | string |  | 是 | 应用client_id |
| client_secret  | string |  | 是 | 应用client_secret |
| pre_auth_code  | string |  | 是 | 预授权码 |

响应

|  属性   | 类型  | 说明 |
|  ----  | ----  |----  |
| access_token  | string | accessToken有效期2小时 |
| refresh_token  | string | refreshToken无限有效时间，服务端保存，丢失后只能重新授权登录 |


### refresh_token刷新access_token

```html
POST /api/oauth/refresh_token

```


请求参数

|  属性   | 类型  | 默认值  | 必填 | 说明 |
|  ----  | ---- |  ----  |----  |----  |
| refresh_token  | string |  | 是 | refreshToken无限有效时间，服务端保存，丢失后只能重新授权登录  |


响应

|  属性   | 类型  | 说明 |
|  ----  | ----  |----  |
| access_token  | string | accessToken有效期2小时 |
| refresh_token  | string | refreshToken无限有效时间，服务端保存，丢失后只能重新授权登录 |



### verify_access_token资源服务器验证accessToken是否有权限访问资源

```html
POST /api/oauth/verify_access_token

```


请求参数

|  属性   | 类型  | 默认值  | 必填 | 说明 |
|  ----  | ---- |  ----  |----  |----  |
| access_token  | string |  | 是 | accessToken有效期2小时 |


响应

|  属性   | 类型  | 说明 |
|  ----  | ----  |----  |
| is_power  | bool | true有权，false无权 |




