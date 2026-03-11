pipeline {
    agent none

    options {
        timeout(time: 10, unit: 'MINUTES')
    }

    stages {
        stage('Controllo Qualita (Golang)') {
            agent { label 'agent-golang' }
            steps {
                echo "Scarico il codice sull'agente Go..."
                checkout scm
                
                echo "Verifico l'ambiente Go..."
                sh 'go version'
                echo "L'agente Go ha finito il suo lavoro!"
            }
        }

        stage('Build (Docker)') {
            agent { label 'agent-docker' }
            steps {
                echo "Scarico il codice sull'agente Docker..."
                checkout scm
                
                dir('src/.docker') {
                    echo "Costruisco l'immagine Docker dell'API..."
                    // Costruiamo solo l'immagine, non alziamo più il compose!
                    sh 'docker build -f service.Dockerfile -t docker-app:latest ../..'
                }
            }
        }

       stage('Deploy (Kubernetes)') {
            agent { label 'agent-kubectl' }
            steps {
                echo "Scarico il codice sull'agente Kubectl..."
                checkout scm
                
                dir('src/.docker/kubernetes') {
                    echo "Configuro l'accesso al cluster (copia temporanea)..."
                    sh '''
                        # Creiamo una cartella temporanea scrivibile
                        mkdir -p /tmp/k8s
                        # Copiamo il file originale lì
                        cp /home/jenkins/.kube/config /tmp/k8s/config
                        # Diciamo a kubectl di usare la copia
                        export KUBECONFIG=/tmp/k8s/config
                        
                        # Ora possiamo modificarlo senza errori!
                        kubectl config set-cluster docker-desktop --server=https://host.docker.internal:6443 --insecure-skip-tls-verify=true
                        
                        echo "Lancio i manifest sul Cluster Kubernetes..."
                        kubectl apply -f pod.yml
                        kubectl apply -f service.yml
                    '''
                }
            }
        }

    post {
        failure {
            echo '❌ La pipeline è fallita. Controlla i log per gli errori.'
        }
        success {
            echo '✅ Pipeline completata! L\'app sta girando su Kubernetes.'
        }
    }
}