dev:
	docker-compose up --remove-orphans

clean:
	docker-compose down && docker volume rm candyshop_go-modules candyshop_pgdata