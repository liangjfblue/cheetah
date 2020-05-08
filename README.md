# cheetah

[![Build Status](https://travis-ci.org/liangjfblue/cheetah.svg?branch=master)](https://travis-ci.org/github/liangjfblue/cheetah)

## 💬介绍
😎😎😎cheetah(猎豹)是一个微服务，分布式架构的任务调度中心

## ☁架构


## 🌟组件


## ⚙特性
- 微服务架构。随心所欲的水平扩展，高可靠，高可用
- 模块组件解耦，方便扩展。网关-web-调度器master-worker分层
- 用户权限管理，接口权限管理
- 支持定时任务，延时任务（redis实现）
- 实时查看任务执行进度（websocket实现）
- 支持多种类型的任务调度
    - 接口级别
    - 脚本启动级别
    - 代码级别（支持shell，golang，python）
- 支持多种任务调度方式
    - 单任务调度
    - 多任务并行调度
    - 因果关系任务调度（任务A->任务B->任务C->任务D，中间有任务运行失败，整个任务调度失败）
- 支持插件化调度算法
    - 内部提供**随机**、**轮训**、**Worker权重**、**Worker最少负载**调度算法
    - 重写提供的接口，可自定义调度算法
- 支持多平台报警
    - 邮件
    - 企业微信
    - web callback url

## 👏技术栈
- go-micro（微服务框架）
- gin（http服务）
- docker（容器化）
- etcd（服务发现注册中心）
- gorm（数据库orm）
- mysql
- redis
- OpenTracing（分布式链路追踪）
- casbin（权限管理）
- Traefik（反向代理）


## 🧪使用
### 1、编译
`./scripts/build.sh all`


### 2、生成Dockerfile
`./scripts/dockerfile.sh all`

### 3、运行
创建`deployments/db/mysql_data` 目录

#### 3.0 打包
进入deployments目录: `sudo docker-compose build`

#### 3.1、运行
进入deployments目录: `sudo docker-compose up`

#### 3.2、停止
进入deployments目录: `sudo docker-compose down`

#### 3.3、水平扩展master，worker
进入deployments目录: `sudo docker-compose --scala srv_xxx=3`


## 分布式链路监控
```http://172.16.7.16:16686/```

## 🗨️TODO
- k8s部署


## 赞助

