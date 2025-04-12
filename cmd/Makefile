docker-run-sql:
	docker run -d --name mysql --privileged=true -e MYSQL_ROOT_PASSWORD="ead8686ba57479778a76e" -e MYSQL_USER="food_delivery" -e MYSQL_PASSWORD="19e5a718a54a9fe0559dfbce6908" -e MYSQL_DATABASE="food_delivery" -p 3308:3308 bitnami/mysql:5.7
mysql-string: 
	"root:2008@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"