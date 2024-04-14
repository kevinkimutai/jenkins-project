pipeline {
    agent any
    
   tools { go '1.22.1' }
    
    stages {
        stage('Checkout') {
            steps {
                echo '--- Checking out the code from version control ---'
                // Checkout your code from version control (e.g., Git)
                git 'https://github.com/kevinkimutai/jenkins-project.git'
            }
        }
        
        stage('Build') {
            steps {
                echo '--- Building the GoLang application ---'
                // Build your GoLang application
                sh 'go build -o main ./cmd/main.go'
            }
        }
        
        stage('Test') {
            steps {
                echo '--- Running tests ---'
                // Run tests if any
                sh 'go test ./...'
            }
        }
        
        stage('Deploy') {
            steps {
                echo '--- Deploying the application ---'
                // Deploy your application
                // You may use tools like Docker, Kubernetes, etc. for deployment
                sh 'echo "Deploying the application"'
                // Example: Deploy to Kubernetes
                // sh 'kubectl apply -f deployment.yaml'
            }
        }
    }
    
    post {
        always {
            echo '--- Cleaning up workspace ---'
            // Clean up workspace after build
            cleanWs()
        }
    }
}
