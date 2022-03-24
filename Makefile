build:
	docker build -t ervanhohoho/lsip-mock-be:0.0.8 .
deploy:
	docker tag ervanhohoho/lsip-mock-be:0.0.8 registry.heroku.com/isip-mock-be/web
	docker push registry.heroku.com/isip-mock-be/web
	heroku container:push web
	heroku container:release web