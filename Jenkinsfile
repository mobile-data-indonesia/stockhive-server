pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/mobile-data-indonesia/stockhive-server.git', branch: 'main'
            }
        }

        stage('Build') {
            steps {
                echo 'Building From Jenkins'
            }
        }
    }
}
