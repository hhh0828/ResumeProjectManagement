pipeline {
    agent any
    tools { go '1.23.2' }
    stages {
        stage('Build') {
            steps {
                echo 'Building... test git push triggered '
                sh 'go version'
                sh 'go mod tidy'
                
            }
        }
        stage('Test') {
            steps {
                echo 'Testing... the token system'
                sh 'go test . -v'
            }
        }
        stage('Docker Build') {
            steps {
                sshagent(['34c716a6-aa67-4d0d-bfcf-75b86238421f']){
                    echo 'Build Docker image'
                    sh '''
                    ssh -v hyunho@211.221.147.21 "cd /ResumeProjectManagement && 
                    git pull &&
                    docker build --pull --rm -f Dockerfile_FS -t hyunhohong/resume:latest . &&
                    echo 'Docker image built successfully'"
                    '''
                    //docker push hyunhohong/resume:testver &&
                    //배포후 최근 5개이내 버전만 남기고 지우는 로직 추가 - capacity 자동 관리
                }
            }
        }
        // 서버 docker stop / rm  컨테이너 실행중지하는 코드 넣어주면 좋을듯.... 
        // 나중에~~~ 쿠버네티스에 파드로 배포할때 도커허브로 푸쉬하고 가져오는, 새로써야할듯. 

        stage('Deploy') {
            steps { 
                sshagent(['34c716a6-aa67-4d0d-bfcf-75b86238421f']) {
                    echo '*********start build***********'
                    echo '*********make ssh connection and set a path for jobs***********'
                    //temporary 임시포트 대체, 이름랜덤.
                    //docker pull - image - from hub. 
                    sh '''
                    ssh hyunho@211.221.147.21 "echo 'start deployment' &&
                    docker stop resumeapi &&
                    docker rm resumeapi &&
                    docker run -d -p 8700:8700 -v ./ResumeProjectManagement:/usr/src/app --name resumeapi hyunhohong/resume:latest &&
                    echo 'Container started successfully' &&
                    docker logs resumeapi" 
                    '''
                }
            }
        }
    }
}
/*
generate ssh key 
ssh-keygen -t rsa -b 4096 -C "jenkins-server"

- set passphrase / no need to set path
- copy the hash256

ssh-copy-id -i ~/.ssh/id_rsa.pub hyunho@211.221.147.21

- enther passphrase

set the permission of file directory
sudo chown -R hyunho:hyunho /ResumeProjectManagement
*/