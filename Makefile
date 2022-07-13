# apiサーバーを起動する
webapi:
	docker-compose up

# テストの実行
test:
	docker-compose run app go test -v ./...
