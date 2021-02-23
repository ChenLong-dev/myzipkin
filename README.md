# myzipkin
zipkin的使用

## 一、Zipkin 介绍
#### Zipkin 是什么？
Zipkin的官方介绍：https://zipkin.apache.org/

Zipkin是一款开源的分布式实时数据追踪系统（Distributed Tracking System），基于 Google Dapper的论文设计而来，由 Twitter 公司开发贡献。其主要功能是聚集来自各个异构系统的实时监控数据。分布式跟踪系统还有其他比较成熟的实现，例如：Naver的Pinpoint、Apache的HTrace、阿里的鹰眼Tracing、京东的Hydra、新浪的Watchman，美团点评的CAT，skywalking等。

#### 为什么用 Zipkin？
随着业务越来越复杂，系统也随之进行各种拆分，特别是随着微服务架构和容器技术的兴起，看似简单的一个应用，后台可能有几十个甚至几百个服务在支撑；一个前端的请求可能需要多次的服务调用最后才能完成；当请求变慢或者不可用时，我们无法得知是哪个后台服务引起的，这时就需要解决如何快速定位服务故障点，Zipkin分布式跟踪系统就能很好的解决这样的问题。

#### Zipkin的一些基本概念？
##### Brave
Brave 是用来装备 Java 程序的类库，提供了面向 Standard Servlet、Spring MVC、Http Client、JAX RS、Jersey、Resteasy 和 MySQL 等接口的装备能力，可以通过编写简单的配置和代码，让基于这些框架构建的应用可以向 Zipkin报告数据。同时 Brave 也提供了非常简单且标准化的接口，在以上封装无法满足要求的时候可以方便扩展与定制。

如下图是 Brave 的结构图。Brave 利用 reporter 向 Zipkin的 Collector 发送 trace 信息。
![avatar](resources/brave.png)

Brave 主要是利用拦截器在请求前和请求后分别埋点。例如 Spingmvc 监控使用 Interceptors，Mysql 监控使用 statementInterceptors。同理 Dubbo 的监控是利用 com.alibaba.dubbo.rpc.Filter 来过滤生产者和消费者的请求。

##### traceId
一次请求全局只有一个traceId。用来在海量的请求中找到同一链路的几次请求。比如servlet服务器接收到用户请求，调用dubbo服务，然后将结果返回给用户，整条链路只有一个traceId。开始于用户请求，结束于用户收到结果。

##### spanId
一个链路中每次请求都会有一个spanId。例如一次rpc，一次sql都会有一个单独的spanId从属于traceId。

##### cs
Clent Sent 客户端发起请求的时间，比如 dubbo 调用端开始执行远程调用之前。

##### cr
Client Receive 客户端收到处理完请求的时间。

##### ss
Server Receive 服务端处理完逻辑的时间。

##### sr
Server Receive 服务端收到调用端请求的时间。

```
 sr - cs = 请求在网络上的耗时
 ss - sr = 服务端处理请求的耗时
 cr - ss = 回应在网络上的耗时
 cr - cs = 一次调用的整体耗时
```

#### Zipkin的工作过程
当用户发起一次调用时，Zipkin 的客户端会在入口处为整条调用链路生成一个全局唯一的 trace id，并为这条链路中的每一次分布式调用生成一个 span id。span 与 span 之间可以有父子嵌套关系，代表分布式调用中的上下游关系。span 和 span 之间可以是兄弟关系，代表当前调用下的两次子调用。一个 trace 由一组 span 组成，可以看成是由 trace 为根节点，span 为若干个子节点的一棵树。

Zipkin 会将 trace 相关的信息在调用链路上传递，并在每个调用边界结束时异步的把当前调用的耗时信息上报给 Zipkin Server。Zipkin Server 在收到 trace 信息后，将其存储起来。随后 Zipkin 的 Web UI 会通过 API 访问的方式从存储中将 trace 信息提取出来分析并展示。

![avatar](resources/process.png)

## 二、Zipkin的部署与运行
Zipkin的 github 地址：https://github.com/apache/incubator-zipkin

Zipkin的 官网 地址：https://zipkin.io/

#### Docker 方式
```
docker run -d -p 9411:9411 openzipkin/zipkin
```

#### Jar 包方式（JDK8）
```
curl -sSL https://zipkin.io/quickstart.sh | bash -s
java -jar zipkin.jar
```
注意：以上方式的 Zipkin 都是基于内存存储，Zipkin 重启后数据会丢失，建议测试环境使用。Zipkin 支持的存储类型有 inMemory、MySql、Cassandra、以及 ElasticsSearch 几种方式。正式环境推荐使用 Cassandra 和 ElasticSearch。

这里介绍一下 zipkin 基于 mysql 存储进行启动的方式：
```
STORAGE_TYPE=mysql
MYSQL_DB=zipkin
MYSQL_USER=root
MYSQL_PASS=123456
MYSQL_HOST=127.0.0.1
MYSQL_TCP_PORT=3306
```

```
STORAGE_TYPE=mysql MYSQL_DB=zipkin MYSQL_USER=root MYSQL_PASS='123456' MYSQL_HOST='127.0.0.1' MYSQL_TCP_PORT=3306 java -jar zipkin.jar > start.logger 2>&1 &
```
启动后，访问 http://127.0.0.1:9411 可以看到效果：
![avatar](resources/show.png)

## 三、代码执行
#### proto文件生成
```
# cd myzipkin/server/proto/hello
# protoc -I. --proto_path=.  --go_out=plugins=grpc:. ./*.proto
```

#### 服务端编译
```
# cd myzipkin/server
# go build
```

#### 客户端编译
```
# cd myzipkin/client
# go build
```

#### 


