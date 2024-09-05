package main

import "net/http"

func main() {
	//go routines and server is alivinig in hereee!!
	//AdjustingScale("./home/assets/resumemanagement.png", 300, 300)
	http.ListenAndServe("0.0.0.0:8700", NewHandlers())

}

/* 0905 todo Dynamic Infra 환경 구축

1. AWS > EC2. T2.Micro > nginX LT Container 설치 - completed
2. nginX 및 인증서 설치 - completed
3. AWS 공인아이피 도메인 등록 - completed
4. nginX proxy pass 설정 - completed - after domain registered on the DNS servers globally
[5. 쿠버네티스 로컬환경 인그레스 서비스 설정 - not completed need to case study.]
6. 로컬환경 방화벽 설정 포트포워딩 - completed

7. resumemanagement app Helm Package and deploying -
[options]
7-1.add a func> query owned IP address when it changes. and send it out to the NginX server and update Nginx Proxy
// 위 상황은 로컬환경의 공인아이피가 변경되었을때, 노드에서 현 공인아이피를 확인 후, 외부서버로 전달 한다. External Name service가 있기때문에 내부로 들어오는건 막혔지만 외부 통신은 가능하다.
7-2. add a pod with above func and change service.yaml.
// 배포할 deployment의 service 매니페스트의 로컬 부분을 수정한다.
// 수정후 배포 완료하고 NginX에 OK사인 전달, 서비스 재개


7-3. 업로드 및 DB 관련 추가 api 개발
	0905 added > projectuploadpage added.
7-4. 리팩토링
	0905 added > data base table interface upload feature implemented.

7-5. 개선점 피드백 받기
7-6. 개선점 픽스
7-7. 아키텍팅 도식화 추상화된 내용 정리 및 resumesite에 업로드

//8. Kubernetes network 정복
//9. Kubernetes statefulset 정복
//10. Kubernetes & Jenkins,helm,Workload 정복
//11. Kuber & prometheus 정복
//12. AWS cloud


*/
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

//todo

// bring the data from server when the user request the page. the data should be a type of writable.
// front page creating required.
// DB remaster >> when server sends data to client, the data should have a PKey.
// DB architecture refactoring. - high level.
//
