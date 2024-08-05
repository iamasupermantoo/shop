### 项目编译方法
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o src/service main.go


### 服务器环境安装
ubuntu 22 运行环境 配置
安装redis 服务
1. apt install redis-server

安装 mysql8.0 服务
1. apt install mysql-server
2. mysql 命令行 use mysql;alter user 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'Aa123098..';flush privileges;

安装 nginx 服务
1. apt install nginx
2. 配置文件目录 /etc/nginx/conf.d/

安装申请 Let’s Encrypt 证书
1. apt install certbot
2. certbot certonly --webroot -w /var/www/src/public -d example.com -d www.example.com
3. 证书生成完毕后，我们可以在 /etc/letsencrypt/live/ 目录下看到对应域名的文件夹，里面存放了指向证书的一些快捷方式。 这时候我们的第一生成证书已经完成了，接下来就是配置我们的web服务器，启用HTTPS。

### Nginx 配置
server {
   listen 80;
   server_name www.cskf6.com cskf6.com;
   index index.html;
   root /var/www/online/public/web;

   location / {
      try_files $uri $uri/ /index.html;
   }
}

server {
   listen 80;
   server_name app.cskf6.com;
   index index.html;

   location / {
      proxy_pass http://localhost:3020;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header REMOTE-HOST $remote_addr;
      proxy_http_version 1.1;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection Upgrade;
   }
}

server {
   listen 80;

   listen 443;
   ssl on;
   ssl_certificate /etc/letsencrypt/live/dfex.online/fullchain.pem;
   ssl_certificate_key /etc/letsencrypt/live/dfex.online/privkey.pem;

   server_name api.home.basic.com;
   location / {
      proxy_pass http://localhost:3050;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header REMOTE-HOST $remote_addr;
      proxy_http_version 1.1;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection Upgrade;
   }
}

安装 前端代码 环境
1. apt install curl;curl https://raw.githubusercontent.com/creationix/nvm/master/install.sh | bash
2. 安装提示 添加环境变量
3. export NVM_DIR="$HOME/.nvm"
   [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
   [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
4. nvm install 18
5. npm install -g yarn
6. cd 项目目录下 运行 yarn install 


### 其他功能
alter table user auto_increment = 8888

### 导出产品sql
mysqldump -u root -p shop product product_attrs_key product_attrs_sku product_attrs_val > app/models/service/consoleService/initDatabase/datas/sql/shop_backup.sql