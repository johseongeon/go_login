# go_login

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

---
# Database : docker mysql container

docker-compose.yml
``` yaml
version: "3.9"

services:
  mysql:
    image: mysql:oraclelinux9
    container_name: mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: {password}
      MYSQL_DATABASE: login
    ports:
      - "{external port}:3306"
    command: --default-authentication-plugin=mysql_native_password
    volumes:
      - mysql_data:/var/lib/mysql
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql

volumes:
  mysql_data:

```
You should set root password & external port in docker-compose.yml

