FROM mysql:8

ADD ./init.sql /docker-entrypoint-initdb.d/init.sql
ADD ./my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf