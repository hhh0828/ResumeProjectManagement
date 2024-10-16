ResumeManagement Web appication.
Kindly look arround in here regarding the codes i have written

Infra Tools or External Resource

Language : Golang x GORM
Database : Postgres
HTML/CSS/JS : Bootstrap Template.
Proxy : NginX
IDE : Visual Code

Authkey activating - Certbot let's encrypt
Develope env : Linux unbutu 20.04, 24.04 LTD , Windows 11 Pro
H/W resource : AWS cloud EC2 t2.micro freetier for nginx server.

Container runtime interface - Docker
Ochestration - Kubernetes

추가 예정 
Prometheus, Jenkins CICD workload, helm package

추가 개발 예정 사항

Feedback 기반 개선 사항 개발 9월 15일 - no feedback received  :( 
업로드 및 DB관련 추가 API 개발 9월 10일 - 완
- ERD 및 시스템 설계도 9월 15일 - 미완 계획 미정


9월6일 업데이트 - 사진 업로드 기능 추가 파일서버 구현 및 Nginx X fileserveer location 추가 



어플리케이션을 하이퍼 V 또는 로컬 환경에서 구동하기위한 도커 및 쿠버네티스 매니페스트파일, GO 언어로 이루어진 기본적인 웹 API 구성을 작성하였습니다.


Test env implemented.

회원가입 로그인 초안 구현 - 9월23일 완료 
JWT 구현,Token provisioning-  token issuing, validating 구현 - 9월 24일 완료
authmidware 구현 -  and providing validating - 9월 25일 백엔드 서버 업그레이드 완료.
미들웨어 구현 완료, 토큰 Issuing and Validating. 

9월 27일 - 로그인 테스트 환경 구축...


목적 : 게시글 수정 권한 차등화

10월 10일
로그인 페이지 구현완료.
로그인 상태 별 쿠기 셋.
로그인 상태에 따른 index page 수정 완료.
block disk cache using, and edit paths of all the css js 


10월 11일부터~ CICD 자동화를 위한 git push trigger server 구현
Jenkins 사용기
ENVPATH =_= 

tool set and git & Jenkins hook setting - completed.

init docker based jenkins.- completed.

Set go plugin completed. set properties - completed.
add tool part code on Jenkinsfile
checked that the tmp user works as Jenkins system in declared pipeline.
test build successed.

generate token for git access on github. and add credential with token to request src - completed


Basic CI CD 

Create Docker based CD.
Create test code for all the components.
so that any change must not make any issue. 

1. build docker image on jenkins server. and push them on the docker hub.
2. access to the server that will have new image, and create docker container with new image. 

Pipeline created
10월15일 Basic CI/CD completed.
Test branch published - 
only need push event for triggered according to a main merge action.
Webhook trigger > Jenkins server > Check if Main branch has any changes, if yes do a test and deploy app on server with ssh agent.

test code will be deployed on test branch in the future, and it will be merged on main.

10월16일
Oauth system migration with main log in system.
-register the member profile using user access token // db set-
-issue a jwt to a new logged user with Naver account.
-add a feature that validate jwt payload claim -
    permission info- Webmaster only can access edit page.
    modify some code
        -validation part
        func ValidateToken(receivedjwt string) bool >>>>>> func ValidateToken(receivedjwt string) (bool, string)
        
        -oauth signed up user check part
        if Checkuser() bool else 

        -below func added
        Checkuser()
        Createuser()

        -interface added just in case adding a new login system like Oauth Kakaotalk, etc. 
        LoginInterface

-fix some bugs
    validation codes were not wokring properly
    comparing oid with userid -> fixed
        Oauth.go >>
        if model.oid == oid  >>>>>>> if model.Userid == oid

Get some capacity of a Node which has docker CRI, Removed some images.