FROM mysql:8

COPY ./sql/*.sql /docker-entrypoint-initdb.d/

COPY ./mysqld.cnf /etc/mysql/conf.d/

RUN mkdir -p /var/log/mysql/

RUN chmod 777 /var/log/mysql/

RUN mysql_tzinfo_to_sql /usr/share/zoneinfo

EXPOSE 3306

CMD ["mysqld"]