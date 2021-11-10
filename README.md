# 1110 Docker

docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:5.7.36

docker run --name phpmyadmin -d --link mysql -e PMA_HOST="mysql" -p 8081:80 phpmyadmin

docker build . -t todolist 

docker run --name todo -d --link mysql:mysql -p 8080:80 todolist

dsn := "root:root@(mysql:3306)/todolist?charset=utf8&parseTime=True&loc=Local"
