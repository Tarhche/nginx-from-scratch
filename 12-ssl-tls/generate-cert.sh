#!/bin/sh

# Create SSL directory if it doesn't exist
mkdir -p /ssl

# Generate SSL certificate with default values (non-interactive)
openssl req -x509 -newkey rsa:2048 \
    -nodes \
    -sha256 \
    -days 365 \
    -keyout /ssl/private.key \
    -out /ssl/public.pem \
    -subj "/C=US/ST=State/L=City/O=Organization/OU=OrganizationalUnit/CN=localhost"
