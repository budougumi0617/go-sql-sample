.PHONY: create up e2e mysql.start mysql.stop

create:
	mysql -h 127.0.0.1 --port ${SHOTEN6_MYSQL_PORT} -u${SHOTEN6_MYSQL_USER} < db/database.sql

up:
	sql-migrate up

e2e:
	go test -v -tags e2e ./...

mysql.start:
	docker run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
		-p $(SHOTEN6_MYSQL_PORT):3306 --name mysql_tmp mysql:5.7

mysql.stop:
	docker stop mysql_tmp
