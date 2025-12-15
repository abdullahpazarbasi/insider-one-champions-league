FROM node:25-bookworm-slim AS build

WORKDIR /

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

RUN npm run build

FROM nginx:1.29-alpine AS runtime

RUN rm -f /etc/nginx/conf.d/default.conf

COPY <<'EOF' /etc/nginx/conf.d/app.conf
server {
  listen 8080;
  server_name _;

  root /usr/share/nginx/html;
  index index.html;

  server_tokens off;
  charset utf-8;

  location /assets/ {
    try_files $uri =404;
    expires 365d;
    add_header Cache-Control "public, immutable";
  }

  location / {
    try_files $uri $uri/ /index.html;
    add_header Cache-Control "no-store";
  }

  location = /healthz {
    return 200 "ok\n";
    add_header Content-Type text/plain;
  }
}
EOF

COPY --from=build /dist /usr/share/nginx/html

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
