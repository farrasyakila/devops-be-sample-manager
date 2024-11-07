pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                script {
                if (env.BRANCH_NAME == 'dev') {
            sh '''
            cd code/
            docker build -t farrasyakila/api-go-dev:$BUILD_NUMBER -f Dockerfile-dev .
            '''
                }
                else if (env.BRANCH_NAME == 'staging') {
            sh '''
            cd code/
            docker build -t farrasyakila/api-go-stg:$BUILD_NUMBER -f Dockerfile-stg .
            '''   
               }
                else if (env.BRANCH_NAME == 'prod') {
            sh '''
            cd code/
            docker build -t farrasyakila/api-go-prod:$BUILD_NUMBER -f Dockerfile-prod .
            '''   
               }
                else {
                    sh 'echo Nothing to Build'
                }
            }
        }
    }
        stage('push') {
            steps {
                script {
                if (env.BRANCH_NAME == 'dev') {
                sh 'docker push farrasyakila/api-go-dev:$BUILD_NUMBER' 
                    }
                    else if (env.BRANCH_NAME == 'staging') {
                sh 'docker push farrasyakila/api-go-stg:$BUILD_NUMBER'
                }
                    else if (env.BRANCH_NAME == 'prod') {
                sh 'docker push farrasyakila/api-go-prod:$BUILD_NUMBER'
                }
                    else {
                        sh 'echo Nothing to Push'
                    }
                }
            }
        } 
        stage('deploy') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'dev') {
                    sh '''
                        cd k8s/dev/
                        # Replace image tag in deploy.yaml with the current BUILD_NUMBER
                        sed -i "s|farrasyakila/api-go-dev|farrasyakila/api-go-dev:${BUILD_NUMBER}|" deploy.yaml
                    
                        cat deploy.yaml | grep image:

                        # Apply Kubernetes configurations
                        kubectl apply -f namespace.yaml
                        kubectl apply -f secret.yaml
                        kubectl apply -f service.yaml
                        kubectl apply -f deploy.yaml
                    '''
                }
                    else if (env.BRANCH_NAME == 'stg') {
                    sh '''
                    cd k8s/staging/
                        # Replace image tag in deploy.yaml with the current BUILD_NUMBER
                        sed -i "s|farrasyakila/api-go-stg|farrasyakila/api-go-stg:${BUILD_NUMBER}|" deploy.yaml
                    
                        cat deploy.yaml | grep image:

                        # Apply Kubernetes configurations
                        kubectl apply -f namespace.yaml
                        kubectl apply -f secret.yaml
                        kubectl apply -f service.yaml
                        kubectl apply -f deploy.yaml
                    '''
                    }
                    else if (env.BRANCH_NAME == 'prod') {
                    sh '''
                    cd k8s/dev/
                    whoami
                        # Replace image tag in deploy.yaml with the current BUILD_NUMBER
                        sed -i "s|farrasyakila/api-go-prod|farrasyakila/api-go-prod:${BUILD_NUMBER}|" deploy.yaml
                    
                        cat deploy.yaml | grep image:

                        # Apply Kubernetes configurations
                        kubectl apply -f namespace.yaml
                        kubectl apply -f secret.yaml
                        kubectl apply -f service.yaml
                        kubectl apply -f deploy.yaml
                    '''
                    }
                    else {
                    sh 'echo Nothing to Deploy'
                    }
                }
            }
        }
    }
}