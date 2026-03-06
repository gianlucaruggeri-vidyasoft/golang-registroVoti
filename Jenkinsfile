pipeline {
    agent none 

    options {
        timeout(time: 10, unit: 'MINUTES')
    }

    environment {
        NOME_IMMAGINE = 'mia-api-go'
        NOME_CONTAINER = 'registro-voti-api'
        PORTA_HOST = '8082'
        PORTA_APP = '8082' // Allineata alla porta che usi solitamente
        DB_NAME = 'registro_voti'
        // CORREZIONE 1: Usiamo il nome del container 'demo-mongodb-1' invece dell'ID
        MONGO_URI = 'mongodb://root:example@demo-mongodb-1:27017'
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
                git branch: 'main', url: 'https://github.com/gianlucaruggeri-vidyasoft/golang-registroVoti.git'
                sh 'go test -v ./...'
            }
        }

        stage('Build e Deploy Docker') {
            agent any 
            steps {
                sh "docker build -t ${NOME_IMMAGINE}:latest ."
                script {
                    sh "docker rm -f ${NOME_CONTAINER} || true"
                    echo "Avvio sulla rete registro-net con porta ${PORTA_HOST}..."
                    // CORREZIONE 2: Aggiunto --network registro-net
                    sh "docker run -d --name ${NOME_CONTAINER} --network registro-net -p ${PORTA_HOST}:${PORTA_APP} -e MONGO_DB=${DB_NAME} -e MONGO_URI=${MONGO_URI} ${NOME_IMMAGINE}:latest"
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
    }
}