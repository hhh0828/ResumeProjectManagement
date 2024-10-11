pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building... test git push triggered '
                sh 'go mod tidy'
                
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test .'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...test onemor againe'
                
            }
        }
    }
}
