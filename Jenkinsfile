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
                echo 'Testing...'
                sh 'go test .'
                echo 'is it working ?'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...test onemor againe'
                
            }
        }
    }
}
