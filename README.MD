<a name="RfyrV"></a>
## 项目简介
项目基于 kratos-mono 结构，对现有公司项目“在线教育系统”进行微服务架构改造。系统以学员购买课程，完成学习任务， 统计分析学习数据为主，以课程、班级营销运营为辅，分为用户、课程、班级、题库、资料库、虚拟仿真库、运营、营销、订单、线下活动等 10 个核心模块以及虚拟仿真客户端。
<a name="oOxYl"></a>

## 项目工程化目录结构

-  `.vscode` 
   - 个人配置喜好，已经添加到`.gitignore`
-  `/api` 
   - 与`app`目录的服务一一对应，维护各个服务的`proto`文件与生成的源码
   - 内部服务一般只需定义`rpc`接口，BFF 服务则需要定义`option (google.api.http)`
   - 也可以去掉`/api`目录，就像`kratos`提供的单例服务`layout`一样，将`api/`放到`app`目录内独立维护
-  `/app/${server_name}/service/internel/` 
   - 需要在不同模块中注入不同的依赖(`ProviderSet`)
   - `BIZ`， 业务组装，定义数据操作方法，供`Service`层调用
   - `Data`，实现数据库/缓存等能力，实现`Biz`定义的接口，供`BIZ`层调用
   - `Conf`，解析`../../configs`内的`yaml`配置文件，供`Data Server`使用（这两个目录会进行一些服务初始化需要用到配置信息，内含监听配置文件简单用法）
   - `Server`，服务注册，提供服务初始化依赖注入
   - `Service`，校验数据，调用`BIZ`组装好的用例，实现GRPC接口
   - 各个模块不能乱跨模块调用: `Service`实现RPC ==> `Biz`定义数据操作方法并组装业务 ==> `Data`数据操作
-  `/app/${server_name}/admin` 运营管理后端微服务
-  `/app/${server_name}/job` 异步工作流服务
-  `/app/${server_name}/task` 任务调度服务
-  `/app/${server_name}/interface` BFF 层，封装组合数据
-  `/deploy` 
   - 存放一些自动化脚本
   - 本地测试使用起来挺方便
-  `pkg` 
   - 存放项目下所有服务均可使用的公共代码
   - 在这里添加公共代码需要按照规范分好目录层级（语意化），不能一坨放一起
-  `third_party` 
   - 存放三方依赖，目前是一些`proto`文件
   - 查看`third_party/google/app/http.proto`可以了解到接口的配置跟参数匹配机制
-  `app_makefile` 
   - `app`目录下每个服务根路径都会有一个`Makefile`文件，内部 `include ../../../app_makefile`
   - 最好不要私自改动`app_makefile`,需要添加功能可以在自己维护的服务内的`Makefile`中添加
-  `Makefile` 
   - 整个项目根目录的`Makefile`
   - 维护一些全局的`Makefile` 命令
-  其他目录自行了解 
<a name="Xskrc"></a>
## 开发规范
<a name="GDsVc"></a>
### 1. 前置条件

- `[Kratos](https://go-kratos.dev/en/docs/getting-started/start/)` `docker docker-compose` 正确安装
<a name="PTkbt"></a>
### 2. 启动

1. `make initdb`  // 初始化环境
2. `make run app=<theServerName>` // 运行服务，服务需要先搭建完毕才能启动
3. `docker`文件位于`/deploy`目录下，路径不对的自行切换
<a name="EXBg0"></a>
### 3. 新增服务

- 新增 course 服务:`make app name=course`
<a name="NuNZI"></a>
### 4. Commit 规范格式如下
例：`type(scope)`:本次提交概述 `git commit -m "Docs(README.md):添加编码规范"` 

-  `type`: 本次 commit 的类型，诸如 Bugfix Docs Style 等，参考如下: 
   - `Init`：初始化
   - `Feat`：添加新功能
   - `Fix`：修补缺陷
   - `Docs`：修改文档
   - `Style`：修改格式
   - `Refactor`：重构
   - `Perf`：优化
   - `Test`：增加测试
   - `Chore`：构建过程或辅助工具的变动
   - `Revert`：回滚到上一个版本
   - `Merge`：合并
-  `scope`: 本次 `commit` 波及的范围 `*` 表示全部工程 
-  `subject`: 简明扼要的阐述下本次 `commit` 的主旨，在原文中特意强调了几点： 
   1. 使用祈使句
   2. 首字母要大写
   3. 结尾无需添加标点
<a name="Vp5cp"></a>
### 5.分支记录

- `dev`：开发版 `prod`：发布版 `stage`：稳定版
- 命令记录：
```shell
git checkout -b dev   // 新建分支
git push origin dev   // 推送新分支到线上
git branch -d dev     // 删除本地分支
git push origin :dev  // 删除远程分支,需要设置权限

git checkout pro // 切换到发布分支
git  merge dev   // 再合并开发分支
```
<a name="hsU4z"></a>
## API 设计
<a name="s4e3h"></a>
#### Api 定义

- 项目中定义 proto ，以 api 包名为根目录
- 路径为：api.<theServiceName>.service.v1，v1 表示 api 版本号
<a name="Id3em"></a>
#### Api 鉴权
在 API Gateway 进行统一的认证拦截，一旦认证成功，我们会使用 JWT 方式通过 RPC 元数据传递的方式带到 BFF 层。BFF 校验 Token 完整性后把身份信息注入到应用的 Context 中，BFF 到其他下层的微服务，建议是直接在 RPC Request 中带入用户身份信息(UserID)请求服务。
<a name="JSJhF"></a>
#### 依赖倒置 

- 使用`wire`维护依赖注入 (`/service/cmd/server/wire.go`)
- 可以观察`/service/cmd/server/wire_gen.go` 中的代码
<a name="UrAG4"></a>
#### 服务拆分

1. `user`服务：用户信息服务
1. `order`服务：订单服务
1. `course`服务：课程服务
1. `classroom`服务：班级服务
1. `questionBank`服务：题库服务
1. `simulation`服务：仿真服务
1. `materialLib`服务：资料库服务
1. ....
<a name="pxK59"></a>
#### 错误码规范

1. 基础服务需要根据业务抛出合适的错误，在`pkg/errors/normal`中自定义`kratos`错误；基础服务应该返回的是`kratos`错误类型
1. BFF层处理(记录日志等操作)具体的错误并返回请求，解析`kratos`错误得到具体底层抛出的错误信息:`e := errors.FromError(err)`
<a name="vV5w7"></a>
#### 日志级别
info：记录提示信息<br />debug：记录调试信息<br />error: 记录错误信息<br />fatal: 记录灾难性错误信息
<a name="JvQiR"></a>
## 微服务基础实施
<a name="Nx8fk"></a>
### 微服务中间件

-  配置中心：zookeeper
- 注册中心：`Nacos`，记录`Nacos`单例模式后台地址:`http://127.0.0.1:8848/nacos/` `login = nacos:nacos` 
- 日志系统：ELK
- 链路追踪：jaeger，docker run --rm --name jaeger -p14268:14268-p16686:16686 jaegertracing/all-in-one
   - trace: endpoint: http://127.0.0.1:14268/api/traces	
- 消息队列：kafka
<a name="SFNvD"></a>
### 其他辅助工具

-  生产推荐 `gitlab`+`k8s`+`lens` 
-  Grpc 测试工具 
   - `[bloomrpc](https://github.com/bloomrpc/bloomrpc)`
   - `[grpcui](https://github.com/fullstorydev/grpcui#installation)` (推荐)
- Http 接口测试工具
   - APIFox（推荐）
<a name="c0800710"></a>
#### 

---

<a name="8f57b97a"></a>
##### 


## 当前计划任务
### 《表设计》
	# courses 课程表 

	# users 用户表 

	# orders 订单表 - 记录用户下单，课程、班级都都可看做订单的商品。

	# pay_cashflows 支付流水表 - 记录用户支付详情。

	# course_members  课程成员表- 用户下单支付成功后，加入到课程成员表中，才有资格访问课程资源 

	# user_learn_stats 用户学习数据统计，统计用户加入多少课程，完成多少课时任务

	# operation  操作记录表，主要用在 admin 后台修改数据的流水信息，用于业务数据"对账"
	
### 《grpc服务》
````
	user/interface  addr: 0.0.0.0:8001 -> BFF 分别从 service 和其它服务组合数据，起来返回给调用者 
   user/service    addr: 0.0.0.0:8002	-> 用户主体的微服务 ，获取用户信息等  
	user/admin      addr: 0.0.0.0:8003	-> 运营管理后端服务 修改用户，并保存修改记录   

	course/interface  addr: 0.0.0.0:8011 -> BFF 分别从 service 和其它服务组合数据，起来返回给调用者 
   course/service    addr: 0.0.0.0:8012	-> 课程主体的微服务 ，获取课程列表，获取课程详情，课程学习等  
	course/admin      addr: 0.0.0.0:8013	-> 运营管理后端服务 修改课程，并保存修改记录   

	order/interface  addr: 0.0.0.0:8021 -> BFF 分别从 service 和task 查 用户的订单和支付数据，组合起来返回给调用者 
	order/service    addr: 0.0.0.0:8022	-> 订单主体的微服务 ，查订单，下订单 等  
	order/admin      addr: 0.0.0.0:8023	-> 运营管理后端服务 修改订单，并保存修改记录  
	order/job 		 addr: 0.0.0.0:8024	-> 从kafka 中捞取订单，把用户加入到课程成员表中     
	order/task 		 addr: 0.0.0.0:8025  -> 计划任务处理过期未支付订单

````
 由于工作时间原因，只能先搭个架子了，后续会继续完善下去