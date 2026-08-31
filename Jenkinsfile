pipeline {
    agent any

    environment {
        PATH     = "/usr/local/go/bin:${env.PATH}"
        AR_HOST  = "us-central1-docker.pkg.dev"
        PROJECT_ID = "ops-lab-506804"
        AR_REPO  = "lab-images"
        IMAGE_NAME = "lab-app"
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

        stage('Bump manifest') {
        when { branch 'main' }
        environment {
            GITOPS_REPO = "github.com/Pebintk/lab-git-ops.git"
            OVERLAY     = "manifests/lab-app/overlays/dev"
        }
        steps {
            withCredentials([string(credentialsId: 'gitops-token', variable: 'TOKEN')]) {
            sh '''
                set -euo pipefail
                rm -rf gitops
                git clone --depth 1 "https://x-access-token:${TOKEN}@${GITOPS_REPO}" gitops
                cd gitops/${OVERLAY}

                kustomize edit set image ${IMG}=${IMG}:${TAG}

                cd "$(git rev-parse --show-toplevel)"
                if git diff --quiet; then
                echo "manifest already at ${TAG}, nothing to commit"
                exit 0
                fi

                git config user.email "jenkins@lab.local"
                git config user.name  "jenkins-ci"
                git commit -am "deploy lab-app ${TAG} (build ${BUILD_NUMBER})"
                git push origin HEAD:main
            '''
            }
        }
        }
    }

    post {
        always {
            sh 'docker image prune -f'
        }
    }
}
