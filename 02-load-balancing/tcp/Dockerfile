FROM nginx:stable

# Install mysql-client
RUN apt-get update && apt-get install -y default-mysql-client && rm -rf /var/lib/apt/lists/*

RUN cat << 'EOF' >> /etc/nginx/nginx.conf
stream {
	include /etc/nginx/stream_conf.d/*.conf;
}
EOF

COPY config /etc/nginx/stream_conf.d

EXPOSE 3075
