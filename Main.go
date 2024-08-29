package main

import "net/http"

func main() {
	//go routines and server is alivinig in hereee!!
	//AdjustingScale("./home/assets/resumemanagement.png", 300, 300)
	http.ListenAndServe("0.0.0.0:8700", NewHandlers())

}

/*
Master Node
laptop - Asus G14
requirement below

kube
ccm
cm
etcd
apiserver


nodes
requirement below

kubelet
kubeadm
docker CRI


worker node1번
HyperV unbunto 20.04 lts
Pod1
Volume1 resource provider - Local directory mount 사용

Pod2
Container4 - ImgFS - transfer file. when the


worker node2번
HyperV unbunto 20.04 lts

Pod1
Container1 - NginX Web Server - proxy and transfer static file. port 80 to 8800 / Pod2 IP address
Container2 - Resumemanagement Server - Web application server port 8800
Pod2 -
Container3 - PostgreDB - DB server
https://kubernetes.io/ko/docs/tutorials/kubernetes-basics/create-cluster/cluster-intro/

*/
