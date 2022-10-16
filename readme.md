# GoLog

## 主要技术栈

### 后端

* :white_check_mark: GIN
* :white_check_mark: GORM-> 数据库
* JWT-GO 鉴权
* Casbin->权限管理
* :white_check_mark: Zap->日志管理
* :white_check_mark: Viper->配置读取
* redis-go 缓存

### 前端

* Vue-cli [VUE 3.x]
* Element UI

## 主要功能

* 用户信息相关：

  * :white_check_mark: 注册 `user/signup`

  * :white_check_mark: 登录 `user/signin`

  * :white_check_mark: 修改密码 `user/change_pwd`

  * :white_check_mark: 编辑更新信息 `user/getEditUser` `user/postEditUser`

* 文章功能：

  * :white_check_mark: 发表文章 `article/postPublishArticle`

  * :white_check_mark: 编辑文章 `article/getEditArticle` `article/postEditArticle`

  * :white_check_mark: 删除文章 `article/deleteArtile/:artilceId`

  * :white_check_mark: 获取某用户的文章 `article/getUserArticles?id=xxx`

  * :white_check_mark: 根据文章id查询文章 `article/getArticle?=id=xxx`

* 评论功能：

  * :white_check_mark: 对文章/评论进行评论 `comment/postCreateComments` 

  * :white_check_mark: 获取文章/评论的所有回复 `comment/postGetComments`

  * :white_check_mark: 获取某用户的所有回复（包括评论和文章） `comment/getUserComments?id=xxx`

* 点赞功能：

  * :white_check_mark: 对文章/评论点赞 `user_like/postLike`

  * :white_check_mark: 对文章/评论取消点赞 `user_like/postUnlike`

  * :white_check_mark: 获取某用户对文章/评论的所有点赞 `user_lie/getUserLike?entity_type=xxx&user_id=xxx`

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

  

  

  