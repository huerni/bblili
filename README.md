<img src="./logo.jpg" style="zoom:50%;" />

<div align="center">
<h1>bblili</h1>
</div>

<div align="center">
<b>仿bilibili后端项目</b>
</div>

<div align="center">
<img src="https://img.shields.io/badge/Golang-1.20-orange"/>
<img src="https://img.shields.io/badge/GoZero-1.6.0-green"/>
<img src="https://img.shields.io/badge/MySQL-8.0-yellowgreen"/>
<img src="https://img.shields.io/badge/Docker-24.0.6-blue"/>
</div>

<div align="center">

<img src="https://img.shields.io/badge/-Etcd-brightgreen"/>
<img src="https://img.shields.io/badge/-Redis-blue"/>
<img src="https://img.shields.io/badge/-ElasticSearch-lightgrey"/>
<img src="https://img.shields.io/badge/-MinIO-blueviolet"/>
</div>

## 项目简介
参考BiliBili视频网站，采用微服务架构实现**用户、视频、弹幕、三连**等核心功能。底层采用缓存、消息队列以及分布式文件服务器等技术支持海量用户操作。  
> 该项目为纯后端项目，无前端页面

功能：
- 用户登陆注册，信息修改
- 实现全文检索并且高亮搜索字段
- 视频分片上传、断点续传、在线播放
- 实时弹幕，在线观看人数统计
- 视频评论，点赞投币收藏

## 项目启动
1. 启动docker组件
```shell
docker-compose up -d
```
2. 启动各个微服务，包含user，file，search，video
```shell
go mod tidy
go run /service/user/user.go
```
3. 启动API服务
```shell
go run /api/bblili.go
```
4. 访问接口 查看/api/bblili.api文件
