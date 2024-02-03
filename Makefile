generate_cache:
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	.out/generator -source cache

generate_prod:
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	.out/generator -source psql

start:
	.out/gqs

build:
	templ generate
	CGO_ENABLED=0 GOOS=linux go build -o .out/gqs cmd/gqs/main.go 
	CGO_ENABLED=0 GOOS=linux go build -o .out/generator cmd/generator/main.go

run:
	templ generate
	CGO_ENABLED=0 GOOS=linux go build -o .out/gqs cmd/gqs/main.go 
	CGO_ENABLED=0 GOOS=linux go build -o .out/generator cmd/generator/main.go
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	.out/generator -source cache
	.out/gqs

setup:
	go get .
	npm install

clear:
	rm .out/*
	rm ./assets/dist/*
	rm ./assets/static/*