FROM httpd:latest

COPY dist/chat /usr/local/apache2/htdocs/

ENTRYPOINT ./chat-app

COPY .htaccess /usr/local/apache2/htdocs/
COPY httpd.conf /usr/local/apache2/conf/httpd.conf

RUN chmod -R 755 /usr/local/apache2/htdocs/

EXPOSE 4200

EXPOSE 8080

EXPOSE 5432
