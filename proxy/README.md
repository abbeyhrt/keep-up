# proxy

## Getting Started

When working with this dockerfile in local development, the following commands
may be helpful:

```bash
# Ran from the root of the project
docker build -t keepupcompany/proxy:0.1 proxy
docker run -p "3000:443" -d --rm keepupcompany/proxy:0.1
```

### Default block debugging

For debugging server blocks, the following snippet seems to be a good getting
started point:

```conf
server {
    listen       80 default_server;
    server_name  _;

    #charset koi8-r;
    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

This is taken and modified from the `nginx` dockerfile `conf.d/default.conf` file.

## Reference

- [Understanding the Nginx Configuration File Structure and Configuration Context](https://www.digitalocean.com/community/tutorials/understanding-the-nginx-configuration-file-structure-and-configuration-contexts)
- Folder reference: https://github.com/Codingpedia/codingmarks-api/wiki/Nginx-Setup--in-Production#server-configuration
- https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-server-blocks-virtual-hosts-on-ubuntu-16-04

### Location block tips

- `location /foo/` will match anything with foo at the beginning
  - `proxy_pass http://some-host` will mean that we end up going to `some-host/foo`
  - `proxy_pass http://some-host/` will mean we end up going to `some-host`
- `location = /foo` will match only `/foo`
