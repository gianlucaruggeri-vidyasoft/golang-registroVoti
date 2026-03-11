pipeline {
    // Diciamo a Jenkins di non assegnare un agente globale, 
    // perché decideremo noi chi fa cosa in ogni stage!
    agent none 

    options {
        timeout(time: 10, unit: 'MINUTES')
    }

    stages {
        // --- PRIMO TEMPO: Tocca all'esperto Go ---
        stage('Controllo Qualita (Golang)') {
            agent { label 'agent-golang' }
            steps {
                echo "Scarico il codice sull'agente Go..."
                git branch: 'main', url: 'https://github.com/gianlucaruggeri-vidyasoft/golang-registroVoti.git'
                
                echo "Verifico l'ambiente Go..."
                sh 'go version'
                
                // In futuro qui potrai lanciare: sh 'go test ./...'
                echo "L'agente Go ha finito il suo lavoro!"
            }
        }

        // --- SECONDO TEMPO: Tocca all'imballatore Docker ---
        stage('Build e Deploy (Docker)') {
            agent { label 'agent-docker' }
            steps {
                echo "Scarico il codice sull'agente Docker..."
                git branch: 'main', url: 'https://github.com/gianlucaruggeri-vidyasoft/golang-registroVoti.git'
                
                // Entriamo nella cartella dove c'è il file compose dell'app
                dir('src/.docker') {
                    echo "Accendo il Database, LocalStack e l'API..."
                    sh 'docker compose up -d --build'
                }
            }
        }
    }

    post {
        success {
            echo "-----------------------------------------------------------"
            echo "🚀 PIPELINE COMPLETATA CON SUCCESSO!"
            echo "L'API e il Database sono online."
            echo "-----------------------------------------------------------"
        }
        failure {
            echo "❌ La pipeline è fallita. Controlla i log per gli errori."
        }
    }
}