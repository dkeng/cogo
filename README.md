# cogo (开发中)

微服务配置文件中心


# Restful Api


## JWT验证

* 请求路由：`/login`
* 请求方式：`POST`
* 请求参数：无

* 请求路由：`/auth/refresh_token`
* 请求方式：`GET`
* 请求参数：无


## 错误响应

|参数名|数据类型|说明|
|--|--|--|
|error|字符串|错误信息|

### 应用

* 请求路由：`/applications/:id`
* 请求方式：`GET`
* 请求参数：无

* 请求路由：`/applications`
* 请求方式：`GET`
* 请求参数：无
  

- 200 响应参数：

  | 参数名     | 数据类型 | 说明 |
  | :--------- | :------- | :--- |
  | app_id     |          |      |
  | app_key    |          |      |
  | app_secret |          |      |
  | created_at |          |      |

  

* 请求路由：`/applications`
* 请求方式：`POST`
* 请求参数：无

* 201 响应参数：

|参数名|数据类型|说明|
|-------|--------|------|
|app_id|数值|应用ID|
|app_key|字符串|应用Key|
|app_secret|字符串|应用密钥|


* 请求路由：`/applications/:id`
* 请求方式：`DELETE`
* 请求参数：无


### 配置文件

* 请求路由：`/configs/:id`
* 请求方式：`GET`
* 请求参数：无

* 请求路由：`/configs`
* 请求方式：`GET`
* 请求参数：无


* 请求路由：`/configs`
* 请求方式：`POST`
* 请求参数：无


* 请求路由：`/configs/:id`
* 请求方式：`PUT`
* 请求参数：无


* 请求路由：`/configs/:id`
* 请求方式：`DELETE`
* 请求参数：无