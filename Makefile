generate_cache:
	templ generate
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	.out/generator -source cache

generate_prod:
	templ generate
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	.out/generator -source psql

start:
	.out/generator -source cache
	.out/gqs

build:
	CGO_ENABLED=0 GOOS=linux go build -o .out/gqs cmd/gqs/main.go 
	CGO_ENABLED=0 GOOS=linux go build -o .out/generator cmd/generator/main.go

setup:
	go get .
	npm install

clear:
	rm .out/*
	rm ./assets/dist/*
	rm ./assets/static/*