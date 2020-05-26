## web

##### nginx 配置

```
server {
    listen        80;
    server_name   localhost;

    location / {
        root    "/xxx/laixhe-web";
        index   index.html;
    }
}
```
