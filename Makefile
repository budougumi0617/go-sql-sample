.PHONY: mysql.start mysql.stop

mysql.start:
	docker run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
		-p $(SHOTEN6_MYSQL_PORT):3306 --name mysql_tmp mysql:5.7

mysql.stop:
	docker stop mysql_tmp
