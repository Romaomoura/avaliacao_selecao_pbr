# avaliacao_selecao_pbr:
- Teste em java de uma pequena api que realizar um cadastro com 3 informações ( Nome completo, telefone e email ).

#Atenção, para rodar a aplicação é necessario:
- framework Spring boot
- Criar um banco de dados chamado api_pbr
- colocar senha do seu banco no application.properties

#INTALE as dependencias utilizadas:
- spring-boot-starter-actuator
- spring-boot-starter-data-jpa'
- spring-boot-starter-web'
- spring-boot-starter-thymeleaf'
- spring-boot-devtools'
- postgresql

#IMPORTANTE
- As configurações do Dockerfile e docker-compose estão incompletas, as imagens são geradas apartir do comando:
  *docker build --tag=openjdk:8-jre* .e *docker-compose up --build --force-recreate* , mas não estão funcionando no docker.
