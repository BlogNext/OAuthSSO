# user模块


- 该模块是开放用户信息的资源



### 获取用户信息

```html
POST /api/resource/user/user_info

```


请求参数

|  属性   | 类型  | 默认值  | 必填 | 说明 |
|  ----  | ---- |  ----  |----  |----  |
| access_token  | string |  | 是 | 授权码  |


响应

|  属性   | 类型  | 说明 |
|  ----  | ----  |----  |
| user_id  | uint64 | 用户的唯一标识 |
| nickname  | string | 用户昵称 |


