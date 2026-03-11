FROM docker:latest

RUN apk add --no-cache curl openjdk17-jre

RUN mkdir -p /var/jenkins

WORKDIR /var/jenkins

CMD curl -sO http://jenkins:8080/jnlpJars/agent.jar && \
    java -jar agent.jar \
    -url http://jenkins:8080/ \
    -secret fb0fb6ff45af25301812d939db5efd917186d907c659acb5a739c025ab916c62 \
    -name "agent-docker" \
    -webSocket \
    -workDir "/var/jenkins"