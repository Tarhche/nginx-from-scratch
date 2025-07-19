FROM nginx:1.28-alpine

# Install mysql-client
RUN apk add --no-cache mysql-client

RUN cat << 'EOF' >> /etc/nginx/nginx.conf
stream {
	include /etc/nginx/stream_conf.d/*.conf;
}
EOF

COPY config /etc/nginx/stream_conf.d

EXPOSE 3075
