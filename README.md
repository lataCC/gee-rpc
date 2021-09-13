# gee-rpc
参考于github开源项目https://github.com/geektutu/7days-golang/tree/master/gee-rpc
<br/>Go 语言广泛地应用于云计算和微服务，成熟的 RPC 框架和微服务框架汗牛充栋。grpc、rpcx、go-micro 等都是非常成熟的框架。一般而言，RPC 是微服务框架的一个子集，微服务框架可以自己实现 RPC 部分，当然，也可以选择不同的 RPC 框架作为通信基座。

考虑性能和功能，上述成熟的框架代码量都比较庞大，而且通常和第三方库，例如 protobuf、etcd、zookeeper 等有比较深的耦合，难以直观地窥视框架的本质。GeeRPC 的目的是以最少的代码，实现 RPC 框架中最为重要的部分，帮助大家理解 RPC 框架在设计时需要考虑什么。代码简洁是第一位的，功能是第二位的。

因此，GeeRPC 选择从零实现 Go 语言官方的标准库 net/rpc，并在此基础上，新增了协议交换(protocol exchange)、注册中心(registry)、服务发现(service discovery)、负载均衡(load balance)、超时处理(timeout processing)等特性