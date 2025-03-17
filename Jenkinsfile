pipeline {
    agent any

    environment {
        DEPLOY_SERVER = "4.237.71.147"
        DEPLOY_USER = "mdi"
        SSH_KEY = "e4bfa090-50e0-428e-93e3-bb43fc9f1b19"
        DEPLOY_PATH = "/home/mdi/stockhive"
    }

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
        stage('Deploy via SSH') {
            steps {
                sshagent (credentials: ["${SSH_KEY}"]) {
                    sh """
                    ssh -o StrictHostKeyChecking=no ${DEPLOY_USER}@${DEPLOY_SERVER} << 'ENDSSH'
                        cd ${DEPLOY_PATH}
                        mkdir testingdirectory
                    ENDSSH
                    """
                }
            }
        }
    }
}
