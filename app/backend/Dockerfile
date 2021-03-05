FROM mysql:latest

ARG mysql_user
ARG mysql_password

ENV MYSQL_RANDOM_ROOT_PASSWORD=true
ENV MYSQL_DATABASE=movie
ENV MYSQL_USER=${mysql_user}
ENV MYSQL_PASSWORD=${mysql_password}

COPY init_db.sql /docker-entrypoint-initdb.d
