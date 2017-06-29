FROM nginx:stable

RUN rm /etc/nginx/conf.d/default.conf
RUN mkdir /etc/nginx/certs
#RUN apt-get update && apt-get install -y curl

CMD ["nginx", "-g", "daemon off;"]
#CMD ["/usr/bin/supervisord"]

EXPOSE 443

COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/certs/domain.crt /etc/nginx/certs/domain.crt
COPY conf/certs/domain.key /etc/nginx/certs/domain.key


