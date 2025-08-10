# Go_login

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

---
# Database : docker mysql container

db/docker-compose.yml
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

and run

```cmd
cd db
docker compose up --build
```

and you should create .env and write

```
MYSQL_URI="root:{password}@tcp(localhost:{external port})/login?charset=utf8mb4&parseTime=true"
JWT_SECRET="jwtsecretkeyshouldbelongandsecure" #example
```

---

# APIs

## /register
<img width="1519" height="683" alt="image" src="https://github.com/user-attachments/assets/a44aa008-8e34-440b-ac18-b00344647b17" />
Content-Type: application/json

response : success or error message

---

## /login
<img width="1533" height="750" alt="image" src="https://github.com/user-attachments/assets/b8225cd2-7a87-4aab-963d-038f88cdc434" />

Content-Type: application/json

response : created jwt
