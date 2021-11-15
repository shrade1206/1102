# 1110 Docker

不使用compose 啟動加連線

docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:5.7.36

docker run --name phpmyadmin -d --link mysql -e PMA_HOST="mysql" -p 8081:8081 phpmyadmin

docker build . -t todolist 

docker run --name todo -d --link mysql:mysql -p 8080:8080 todolist

dsn := "root:root@(mysql:3306)/todolist?charset=utf8&parseTime=True&loc=Local"

# 1111 Docker 更新

docker build . -t todolist 

docker-compose up -d

可直接啟動，MySQL的DATABASE預設utf8還沒完成

# 1112 Docker 更新

基本compose自動化，預設utf8已完成(重新build cnf的部分)

docker-compose up -d --build

docker-compose ps

docker-compose build (compose 設定的 name)

docker-compose stop

docker-compose restart (compose 設定的 name)

docker image prune -f -a (-f代表不用確認(Y/N) 刪除所有沒使用的image)

docker build -f ./mysql/db.Dockerfile -t (imageName) . (.代表執行工作區，務必確認Dockerfile的執行地點)
			這裡的-f 是檔案位置
-h 可以得到詳細說明
