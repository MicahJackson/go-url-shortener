redis-bg-start:
	brew services start redis

redis-bg-stop:
	brew services stop redis

redis-bg-status:
	brew services info redis

build:
	docker build -t my-url-shortener .

prune:
	docker system prune -af