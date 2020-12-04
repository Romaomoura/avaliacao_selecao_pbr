FROM openjdk:8-jre

ARG PROFILE
ARG ADDITIONAL_OPTS

ENV PROFILE=${PROFILE}
ENV ADDITIONAL_OPTS=${ADDITIONAL_OPTS}

WORKDIR /opt/pbrapp

# Falta configurar esse caminho para a copia da imagem
COPY target/pbrapp*.jar pbrapp.jar 

SHELL ["/bin/sh", "-c"]

EXPOSE 5005
EXPOSE 8080

CMD java ${ADDITIONAL_OPTS} -jar pbrapp.jar --spring.profiles.active=${PROFILE}