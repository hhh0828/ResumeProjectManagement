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
                echo 'is it working ?'
            }
        }
        stage('Deploy') {
            steps {

                
                echo 'Deploying...test onemor againe'
                //git hub -push and sync trigerred
                //sh ssh 211.221.147.21 access and
                //sh cd /Resumemanagement - folder move
                //sh sudo git pull
                //sh docker build --pull --rm -f "Dockerfile_FS" -t hyunhohong/resume:testver .
                //sh docker push hyunhohong/resume:testver
                //sh docker container run --name resumeapi -d -p 8700:8700 -v ./:/usr/src/app hyunhohong/resume:testver
            }
        }
    }
}
