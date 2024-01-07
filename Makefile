generate_cache:
	templ generate
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	go build .
	./gqs -mode gen -source cache

generate_prod:
	templ generate
	tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
	go build .
	./gqs -mode gen -source psql

start:
	./gqs -mode serve

setup:
	go get .
	npm install

clear:
	rm ./gqs
	rm ./assets/dist/*
	rm ./assets/static/*