FROM postgres:15.1-alpine

EXPOSE 5432

COPY *.sql /docker-entrypoint-initdb.d/