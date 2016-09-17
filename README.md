douban-rent-tools 
=================
[豆瓣小组抓取](https://github.com/soyking/douban-group-spider) 查询前端, 基于 Golang 和 React

## 使用

运行

```
# 前端
cd static
npm install
npm run build
# 调试 npm run watch 

# 后端
go build
./douban-rent-tools
```

帮助

```
./douban-rent-tools -h
```

页面

![image](https://github.com/soyking/douban-rent-tools/raw/master/images/page.png)

## 参数

```
  # 服务器设置
  -port string
        监听端口, 默认 8080

  # 存储设置      
  -es_addr string
        默认存储是 ElasticSearch, 默认 127.0.0.1:9200
  -es_index string
        ElasticSearch 索引, 如果不存在会自动建立 mapping, 默认 db_rent
  -mongo bool
        使用 MongoDB, 但其中文搜索支持比较麻烦, 所以只做正则查询    
  -mg_addr string
        MongoDB 地址, 默认 127.0.0.1:27017
  -mg_db string
        MongoDB 数据库, 默认 db_rent
  -mg_usr string
        MongoDB 用户名, 默认空
  -mg_pwd string
        MongoDB 密码, 默认空
```

## 自定义

自定义 server (`server/server.go#Server`)

* 存储接口 `storage/storage.go#StorageSave`

* 关键词扩展, 目前支持地铁路线查询扩展`subway.go`(输入 地铁：6), 房间数扩展`room.go`(输入 房间：2)
