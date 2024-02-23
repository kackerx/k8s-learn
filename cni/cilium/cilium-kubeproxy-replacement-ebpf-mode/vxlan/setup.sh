set -v

cat <<EOF | kind create cluster --name=cilium-kubeproxy-replacement-ebpf-vxlan --image=kindest/node:v1.23.4 --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  disableDefaultCNI: true
  # kubeProxyMode: "node" # enable kubeproxy
nodes:
  - role: control-plane
  - role: worker
  - role: worker
EOF

# remove taints
controller_node_ip=172.18.0.7
kubectl taint nodes node/cilium-kubeproxy-control-plane node-role.kubernetes.io/master:NoSchedule-

# install cni
helm repo add cilium https://helm.cilium.io >/dev/null 2>&1
helm repo update >/dev/null 2>&1

# VxLAN Routing Options
# --set kubeProxyReplacement=strict
# --set tunnel=disabled # 关闭隧道, 就是native routing
# --set autoDirectNodeRoutes=true
# --set ipv4NativeRoutingCIDR="10.0.0.0/8"
# Host Routing[EBPF] --set bpf.masquerade=true
helm install cilium cilium/cilium --set k8sServiceHost=$controller_node_ip --set k8sServicePort=6443 --version 1.13.0-rc5 \
  --namespace kube-system --set debug.enabled=true --set debug.verbose=datapath \
  --set monitorAggregation=none --set ipam.mode=cluster-pool \
  --set cluster.name=cilium-kubeproxy-replacement-ebpf-vxlan --set kubeProxyReplacement=strict \
  --set bpf.masquerade=true
