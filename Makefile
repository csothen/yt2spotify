include .env

build-sql-init:
	@./db/setup.sh ${MYSQL_DB_USER} ${MYSQL_DB_NAME}

teardown:
	@docker-compose down
	@-./scripts/teardown.sh

start: teardown build-sql-init
	@docker-compose --env-file ./.env up

db-cli:
	@docker exec -it ${APP_NAME}_db mysql -h localhost -P 3306 -u ${MYSQL_DB_USER} -p${MYSQL_DB_PASSWORD} ${MYSQL_DB_NAME}