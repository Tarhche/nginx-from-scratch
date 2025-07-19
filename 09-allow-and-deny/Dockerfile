FROM nginx:1.28-alpine

# Install apache2-utils (htpasswd) on Alpine
RUN apk add --no-cache apache2-utils

# use htpasswd to create file for authentication
RUN htpasswd -cb /etc/nginx/.htpasswd admin pass123
RUN htpasswd -b /etc/nginx/.htpasswd user pass456

COPY config /etc/nginx/conf.d

COPY public /var/www/public

EXPOSE 9050-9052
