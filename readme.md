# GoLog

## 主要技术栈

### 后端

* GIN
* GORM-> 数据库
* JWT-GO 鉴权
* Casbin->权限管理
* Zap->日志管理
* Viper->配置读取
* redis-go 缓存

### 前端

* Vue-cli [VUE 3.x]
* Element UI

## 主要功能

* 注册，登录

* 用户权限管理

* 基于角色的访问控制模型

* 数据库初始化和CRUD

* redis 缓存

* operation record

* 不同角色具体权限：

  * 用户：
    * 用户信息管理
      * CRUD
    * 博客
      * CRUD
  * 管理员：
    * 用户管理
    * 内容管理
    * 数据管理
    * 系统管理【各项统计数据报表监控】

  

  

  