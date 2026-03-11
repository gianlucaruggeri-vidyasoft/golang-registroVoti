FROM golang:latest

RUN apt-get update && \
    apt-get install -y curl default-jre && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/jenkins

WORKDIR /var/jenkins

CMD curl -sO http://jenkins:8080/jnlpJars/agent.jar && \
    java -jar agent.jar \
    -url http://jenkins:8080/ \
    -secret 501bc408098e9610cc1c4295608f8297f9cfa8eef8440c6dd3ae4a6a7e483009 \
    -name "agent-golang" \
    -webSocket \
    -workDir "/var/jenkins"