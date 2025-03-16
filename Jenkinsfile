pipeline {
    agent any

    environment {
        REMOTE_USER = 'stockhive-mdi'
        REMOTE_HOST = '94.21.0.144' 
        REMOTE_SSH_KEY = credentials('ssh-private-key-id') 
        IMAGE_NAME = 'stockhive-dev'
        CONTAINER_NAME = 'stockhive'
    }

    stages {
        stage('Checkout Code') {
            when {
                branch 'development'
            }
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            when {
                branch 'development'
            }
            steps {
                script {
                    sh "docker build -t ${IMAGE_NAME}:latest ."
                }
            }
        }

        stage('Push Docker Image (optional)') {
            when {
                branch 'development'
            }
            steps {
                sh "docker tag ${IMAGE_NAME}:latest stockhive/${IMAGE_NAME}:latest"
                sh "docker push stockhive/${IMAGE_NAME}:latest"
            }
        }

        stage('Deploy to Cloud via SSH') {
            when {
                branch 'development'
            }
            steps {
                script {
                    sshagent(credentials: ['ssh-private-key-id']) {
                        sh """
                        ssh -o StrictHostKeyChecking=no ${REMOTE_USER}@${REMOTE_HOST} << 'EOF'
                            docker stop ${CONTAINER_NAME} || true
                            docker rm ${CONTAINER_NAME} || true
                            docker rmi ${IMAGE_NAME}:latest || true

                            # Clone ulang repo jika perlu
                            cd ~/apps/${IMAGE_NAME} || mkdir -p ~/apps/${IMAGE_NAME} && cd ~/apps/${IMAGE_NAME}
                            git pull origin development

                            docker build -t ${IMAGE_NAME}:latest .
                            docker run -d --name ${CONTAINER_NAME} -p 80:80 ${IMAGE_NAME}:latest
                        EOF
                        """
                    }
                }
            }
        }
    }
}
