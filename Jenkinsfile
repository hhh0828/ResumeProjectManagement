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
                echo 'Build Docker image'
                sh '''
                ssh -v hyunho@211.221.147.21 "cd /ResumeProjectManagement && 
                git pull &&
                docker build --pull --rm -f Dockerfile_FS -t hyunhohong/resume:testver . &&
                docker push hyunhohong/resume:testver &&
                echo 'Docker image built and pushed successfully'"
                '''
            }
        }
        stage('Deploy') {
            steps {
                echo '*********start build***********'
                echo '*********make ssh connection and set a path for jobs***********'
                //docker stop resumeapi && docker rm resumeapi &&
                sh '''
                ssh hyunho@211.221.147.21 "docker pull hyunhohong/resume:testver &&
                docker run -d -p 8771:8771 -v ./:/usr/src/app hyunhohong/resume:testver &&
                echo 'Container started successfully'"
                '''
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