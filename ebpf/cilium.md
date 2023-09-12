## cilium数据面路径

### 跨节点Pod -> Pod

1. 宿主机的 veth lxc 接口的 tc ingress hook 的 eBPF 程序处理
2. kernel stack 进行路由查找
3. 物理口 eth0 的 tc egress 发出到达 Node2 的eth0
4. Node2 的 eth0 的 tc ingress hook 的 eBPF处理
5. kernel stack 进行路由查找
6. cilium_host 的 tc egress hook
7. redirect 到目的 Pod的宿主机 lxc 接口上

### 同节点 pod -> pod

1. 宿主机 lxc 接口的tc ingress hook的eBPF 程序处理
2. eBPF最终会查到目的pod
3. 通过redirect 方法将流量直接越过kernel stack，送给目的Pod的lxc口

### 访问NodePort Local Endpoint

Client->NodePort

1. 流量从 Node2 的 eth0 口进入
2. 经过 tc ingress hook eBPF Dnat 处理后
3. 将流量发给 kernel，kernel 查路由转发给 cilium_host 接口
4. cilium_host 接口的 tc ingress 收到流量后直接 redirect 流量到目的 pod 的宿主机 lxc 接口上

NodePort->Client

1. 经过 veth 宿主机侧接口口 lxc 的 tc ingress hook eBPF 反 Dnat 处理
2. eBPF 程序直接将流量从定向到物理口，该过程不经过 kernel
3. 最终在经过物理口的 tc egress hook 点后发给 Client

### 访问 NodePort Remote EndPoint

1. 流量发给 Node1 接口 eth0 
2. 经过 tc ingress hook eBPF 先进行 Dnat 将目的地址变成目的 Pod 的 IP 和 Port
3. 后进行 Snat 将源地址变成 Node1 物理口的 IP 和 Port
4. 将流量 redirect 到 Node1 的 eth0 物理口上，流量会经过 tc egress hook 点后发出

### Pod 访问外网

流量发出

1. 经过宿主机的 veth 的 lxc 口的 tc ingress hook 的 eBPF 程序处理后
2. 送给 kernel stack 进行路由查找，确定需要从 eth0 发出
3. 流量回发给物理口 eth0
4. 经过物理口 eth0 的 tc egress eBPF 程序做个 Snat，将源地址转换成 Node 节点的物理接口的 IP 和 Port 发出

反向流量

1. 从外网回来的反向流量，经过 Node 节点物理口 eth0 tc ingress eBPF 程序进行 Snat 还原
2. 将流量发给 kernel 查路由，流量流到 cilium_host 接口后，经过 tc egress eBPF 程序
3. 将流量直接 redirect 到目的 Pod 的宿主机 lxc 接口上

### 主机访问pod

主机访问 Pod 流量使用 cilium_host 口发出，所以在 tc egress hook 点 eBPF 程序直接 redirect 到目的 Pod 的宿主机 lxc 接口上，最终发给 Pod。

反向流量，从 Pod 发出到宿主机的接口 lxc，经过 tc ingress hook eBPF 识别送给 kernel stack，回到宿主机上。



### 内核5.10以后，cilium新增eBPF Host-Routing功能，不需要经过宿主机的lxc接口，直接导入pod内部的veth网卡













