# 로그인 서버

## user struct

```go
type User struct {
	ID       int
	Username string `json:"username"`
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}
```
DB 연결
# 
---

### docker mysql container

1. Pull Image

```
docker pull mysql:oraclelinux9
```

2. Run Container with tags

local 환경에서 3306포트를 mysql이 사용중이라 3307 사용했음.

{password}에 원하는 비밀번호값 입력

```
docker run -d -p 3307:3306 --name mysql -e MYSQL_ROOT_PASSWORD={password} mysql:oraclelinux9
```


3. Create Database & Table
```
# 컨테이너 접속
docker exec -it CONTAINER_NAME mysql -uroot -p

# QUERY
CREATE DATABASE login;
USE login;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    user_id VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

4. root directory에 .env파일 생성, 다음 입력
아까 2.에서 입력한 {password}
```
MYSQL_URI=”root:{password}@tcp(loaclhost:3307)/login?parseTime=true”
```
