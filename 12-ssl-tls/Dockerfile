# building stage - generate ssl key
FROM alpine:3.22 AS builder

WORKDIR /ssl

RUN apk add --no-cache openssl

COPY generate-cert.sh .
RUN chmod +x generate-cert.sh && ./generate-cert.sh

# runtime stage - nginx
FROM nginx:stable

COPY --from=builder /ssl/private.key /etc/nginx/ssl/private.key
COPY --from=builder /ssl/public.pem /etc/nginx/ssl/public.pem

COPY ./config /etc/nginx/conf.d/

EXPOSE 8080
