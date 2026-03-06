pipeline {
    agent none 

    options {
        timeout(time: 10, unit: 'MINUTES')
    }

    environment {
        NOME_IMMAGINE = 'mia-api-go'
        NOME_CONTAINER = 'registro-voti-api'
        PORTA_HOST = '8082'
        PORTA_APP = '3000'
        DB_NAME = 'registro_voti'
        MONGO_URI = 'mongodb://host.docker.internal:27017'
    }

    stages {
        stage('Esecuzione Test') {
         agent {
                docker {
                    image 'golang:1.26' 
                    reuseNode true
                }
            }
            steps {
                // Questo checkout serve perché siamo in un nuovo contenitore isolato
                git branch: 'main', url: 'https://github.com/gianlucaruggeri-vidyasoft/golang-registroVoti.git'
                sh 'go test -v ./...'
            }
        }

        stage('Build e Deploy Docker') {
            // Per costruire l'immagine e farla partire, torniamo sull'agente "host" 
            // che ha il motore Docker installato
            agent any 
            steps {
                sh "docker build -t ${NOME_IMMAGINE}:latest ."
                script {
                    sh "docker rm -f ${NOME_CONTAINER} || true"
                    echo "Avvio con porta interna ${PORTA_APP} e variabili DB..."
                    sh "docker run -d --name ${NOME_CONTAINER} -p ${PORTA_HOST}:${PORTA_APP} -e MONGO_DB=${DB_NAME} -e MONGO_URI=${MONGO_URI} ${NOME_IMMAGINE}:latest"
                }
            }
        }
    }

    post {
        success {
            echo "-----------------------------------------------------------"
            echo "PIPELINE SUCCESS CON GLI AGENTS!"
            echo "API disponibile su: http://localhost:${PORTA_HOST}"
            echo "-----------------------------------------------------------"
        }
    }
}