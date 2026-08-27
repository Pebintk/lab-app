pipeline {
    agent any

    environment {
        PATH     = "/usr/local/go/bin:${env.PATH}"
        AR_HOST  = "us-central1-docker.pkg.dev"
        PROJECT_ID = "ops-lab-506804"
        AR_REPO  = "lab-images"
        IMAGE_NAME = "hello"
        IMG      = "${AR_HOST}/${PROJECT_ID}/${AR_REPO}/${IMAGE_NAME}"
    }

    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                script {
                    env.TAG = env.GIT_COMMIT.take(7)
                }
                sh "docker build --build-arg VERSION=${env.TAG} -t ${IMG}:${env.TAG} ."
            }
        }

        stage('Push') {
            steps {
                sh "docker push ${IMG}:${env.TAG}"
            }
        }
    }

    post {
        always {
            sh 'docker image prune -f'
        }
    }
}
