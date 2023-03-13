## goin-mframe
goin-mframe是使用golang gin Web框架和gorm关系型数据库面向对象数据库操作进行二次开发的新Web框架。框架对业务流程开发的每个环节进行的代码封装，开发人员只需要关心相关的接口配置和业务逻辑
开发。框架提升了开发人员的开发效率和规范了开发人员的代码设计。后期会增加cache、crontab、script、httplib、config等功能引入。
# 例子
## Restful Api接口
---
- 查询指定数据：http://localhost:8080/v1/api/user/list?uname=xxxxxx 
- 查询表所有数据： http://localhost:8080/v1/api/user/list 
- 创建用户： http://localhost:8080/v1/api/user/create
- 更新用户： http://localhost:8080/v1/api/user/update
- 用户详情：http://localhost:8080/v1/api/user/view?id=20
---
## Admin管理页面
---
- 后台首页：http://localhost:8080/admin/index
---
## 版本更新
---
- 3月10日初次提交goin-mframe MVC Web框架。
- 3月11日提交框架日志类GoinLog。
- 3月13日提交Dockerfile Docker容器化部署。
---
