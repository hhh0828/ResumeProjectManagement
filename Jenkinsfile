pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go version'
                sh 'go mod tidy'
                // 여기서 빌드 명령어 실행 (예: mvn clean install)
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                // 여기서 테스트 명령어 실행 (예: mvn test)
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
                // 배포 단계 실행 (예: scp, rsync 등)
            }
        }
    }
}
