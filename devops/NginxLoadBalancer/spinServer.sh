# # Terminal 1
# python3 -m http.server 8001 --directory /tmp &

# # Terminal 2
# python3 -m http.server 8002 --directory /tmp &

# echo "Hello from backend-1" >/tmp/index.html
# # For backend-2:
# echo "Hello from backend-2" | sudo tee /tmp/index.html  # if same file reused, skip


mkdir -p /tmp/app1 /tmp/app2
echo "Hello from backend-1" > /tmp/app1/index.html
echo "Hello from backend-2" > /tmp/app2/index.html

# Start the two backends
python3 -m http.server 8001 --directory /tmp/app1 &
python3 -m http.server 8002 --directory /tmp/app2 &

sudo nginx -t && sudo nginx -s reload

sudo apt-get update && sudo apt-get install -y nginx

sudo tee /etc/nginx/sites-available/lab.conf >/dev/null <<'NGX'
upstream app_pool {
    # Default algorithm is round_robin
    server 127.0.0.1:8001 max_fails=3 fail_timeout=10s;
    server 127.0.0.1:8002 max_fails=3 fail_timeout=10s;

    # Examples (uncomment ONE to try)
    # least_conn;          # Choose server with fewest active connections
    # hash $remote_addr;   # "Sticky" based on client IP
}

server {
    listen 8080;
    server_name _;

    # Basic passive health/failover via max_fails/fail_timeout above
    # (Active health checks require NGINX Plus or a 3rd-party module.)

    location / {
        proxy_pass http://app_pool;
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # Simple status page (requires stub_status)
    location /nginx_status {
        stub_status;
        allow 127.0.0.1;
        deny all;
    }
}
NGX

sudo ln -s /etc/nginx/sites-available/lab.conf /etc/nginx/sites-enabled/lab.conf || true
echo "server_tokens off;" | sudo tee /etc/nginx/conf.d/hardening.conf >/dev/null
sudo nginx -t
sudo service nginx start
sudo service nginx stop
sudo service nginx restart


curl -s http://localhost:8080 | head -1

curl -I http://localhost:8080 | grep X-Backend