build:
	docker build -t ervanhohoho/lsip-mock-be:0.0.2 .
deploy:
	docker tag ervanhohoho/lsip-mock-be:0.0.2 registry.heroku.com/isip-mock-be/web
	docker push registry.heroku.com/isip-mock-be/web
	heroku container:push web
	heroku container:release web