# Hướng dẫn


1. Cài đặt [docker](https://www.docker.com/products/docker-desktop) 
1. Cài [mysql](https://www.mysql.com/downloads/) với các config:
    - Username: root
    - Password: root
    - Port: 3306
1. Tạo database: `lessin`
1. Kiểm tra ip hiện tại của máy. Tiến hành edit ./leesin/cmd/main.go
 localhost -> your IP
    ```golang
    db, err := database.NewDatabase("root", "root", "leesin", "localhost", "3306")
    ```
1. Vào terminal gõ: `docker build -t leesin_service:v0.0.1 -f ./leesin/docker/Dockerfile_service .
`

1. sau khi build thành công gõ: `docker run leesin_service:v0.0.1
`

1. Xem log để biết các API nào public và cào posman để test kết quả
