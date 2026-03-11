FROM jenkins/inbound-agent:latest

USER root

# Installiamo curl e scarichiamo l'eseguibile di kubectl
RUN apt-get update && apt-get install -y curl && \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
    RUN mkdir -p /var/jenkins/workspace && chown -R jenkins:jenkins /var/jenkins

USER jenkins