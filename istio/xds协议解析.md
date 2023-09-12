## xDS是什么

1. xDS是一类发现服务的总称，包含LDS，RDS，CDS，EDS以及SDS
2. Envoy 通过xDS API 可以动态获取Listener(监听器)，Route（路由），Cluster（集群），Endpoint（集群成员）以及Secret（证书）配置

### LDS

1. Listener 发现服务
2. Listener 监听器控制Envoy启动端口监听（目前只支持TCP协议），并配置L3/L4层过滤器，当网络连接达到后，配置好的网络过滤器堆栈开始处理后续事件。这种通用的监听器体系结构用于执行大多数不同的处理任务（限流，客户端认证，HTTP连接管理，TCP代理等）

### RDS

1. Route发现服务，用于HTTP连接管理过滤器动态获取路由配置。
2. 路由配置包含HTTP头部修改（增加、删除HTTP头部键值），virtual hosts （虚拟主机），以及virtual hosts 定义的各个路由条目。

### CDS

1. Cluster发现服务，用于动态获取Cluster信息。（cluster类似于service）
2. Envoy cluster管理器管理着所有的上游cluster。
3. 鉴于上游cluster或者主机可用于任何代理转发任务，所以上游cluster一般从Listener或Route中抽象出来。

### EDS

1. Endpoint发现服务。（endpoint类似于pod ip）
   1. 在Envoy术语中， Cluster成员就叫 Endpoint，对于每个Cluster， Envoy通过EDS API 动态获取Endpoint。
2. EDS 作为首选的服务发现的原因有两点：
   1. 与通过DNS解析的负载均衡器进行路由相比， Envoy能明确的知道每个上游主机的信息，因而可以做出更加智能的负载均衡决策。
   2. Endpoint配置包含负载均衡权重、可用域等附加主机属性，这些属性可用域服务网格负载均衡，统计收集等过程中。

### SDS

1. Secret发现服务，用于运行时动态获取TLS证书
2. 若没有SDS特性，在k8s环境中，必须创建包含证书的Secret，代理启动前Secret必须挂载到sidecar容器中，如果证书过期，则需要重新部署。使用SDS，集中式的SDS 服务器将证书分发给所有的Envoy实例，如果证书过期，服务器会将新的证书分发， Envoy 接收到新的证书后重新加载儿不用重新部署

## ADS

ADS 是一种xDS 的实现，它基于gRPC长连接

在istio 0.8以前，Pilot提供单一资源的DS

- 每种资源需要一条单独的连接
- Istio 高可用环境下，可能部署多个Pilot

带来的挑战：

- 没办法保证配置资源更新的顺序
- 多Pilot配置资源的一致性没法保证

综合以上两个问题，很容易出现配置更新过程中网络流量丢失带来的网络错误

解决：ADS允许通过一条连接（gRPC的同一stream），发送多种资源的请求和响应

1. 能够保证请求一定落在同一Pilot上，解决多个管理服务器配置不一致的问题
2. 通过顺序的配置分发，轻松解决资源更新顺序的问题











