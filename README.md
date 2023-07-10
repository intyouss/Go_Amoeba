# Go_Amoeba

---
__基于Go实现分布式服务系统框架__

所有服务将通过 registry 注册模块，进行服务注册。<br>
registry 注册模块将启动服务发送给依赖服务，进行服务发现。<br>
service 基础服务调度模块，提供全局基础服务启动功能。<br>
所有基础服务可互相依赖，并分布式提供服务。<br>

## 模块分组

---
### registry 注册模块

提供:
- 注册，删除服务更新功能。 
- 被动通知和主动发现服务依赖功能。 
- 心跳监测功能。

### service 基础服务调度模块

提供:
- 基础服务启动功能。

### log 日志服务模块

提供:
- 自定义日志输出与文件存储功能。

### grades 学生系统分数模块

提供:
- 学生分数信息查询功能
- 学生分数信息求值功能

依赖:
- log 日志服务

### portal 学生系统分数查询 web 服务模块

提供:
- web显示
- 学生分数添加，删除功能

依赖:
- log 日志服务
- grades 分数查询服务
