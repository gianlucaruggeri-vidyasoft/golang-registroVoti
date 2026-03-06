pipeline {
    // Diciamo a Jenkins di usare il tuo nuovo Agente "1" per tutta la pipeline
    agent { label 'linux-worker' } 

    options {
        timeout(time: 10, unit: 'MINUTES')
    }

    environment {
        NOME_IMMAGINE = 'mia-api-go'
        NOME_CONTAINER = 'registro-voti-api'
        PORTA_HOST = '8082'
        PORTA_APP = '3000' 
        DB_NAME = 'registro_voti'
        // Usiamo il nome del container Mongo nella rete registro-net
        MONGO_URI = 'mongodb://root:secret@demo-mongodb-1:27017/?authSource=admin'
    }

    stages {
        stage('Esecuzione Test') {
            agent {
                docker {
                    image 'golang:1.26' 
                    // reuseNode true è fondamentale: dice a Docker di girare DENTRO l'Agente "1"
                    reuseNode true
                }
            }
            steps {
                // Checkout del codice
                git branch: 'main', url: 'https://github.com/gianlucaruggeri-vidyasoft/golang-registroVoti.git'
                // Esecuzione dei test Go
                sh 'go test -v ./...'
            }
        }

        stage('Build e Deploy Docker') {
            // Qui non serve specificare l'agente, usa quello globale (linux-worker)
            steps {
                // Creazione dell'immagine usando il Dockerfile del progetto
                sh "docker build -t ${NOME_IMMAGINE}:latest ."
                
                script {
                    // Pulizia: fermiamo il vecchio container se esiste
                    sh "docker rm -f ${NOME_CONTAINER} || true"
                    
                    echo "Avvio sulla rete registro-net con porta ${PORTA_HOST}..."
                    // Lancio del nuovo container collegato a MongoDB
                    sh """
                        docker run -d \
                        --name ${NOME_CONTAINER} \
                        --network registro-net \
                        -p ${PORTA_HOST}:${PORTA_APP} \
                        -e MONGO_DB=${DB_NAME} \
                        -e MONGO_URI=${MONGO_URI} \
                        ${NOME_IMMAGINE}:latest
                    """
                }
            }
        }
    }

    post {
        success {
            echo "-----------------------------------------------------------"
            echo "PIPELINE SUCCESS! L'API è ora collegata a MongoDB."
            echo "API disponibile su: http://localhost:${PORTA_HOST}"
            echo "-----------------------------------------------------------"
        }
        failure {
            echo "La pipeline è fallita. Controlla i log per gli errori."
        }
    }
}