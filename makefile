docli:
	go build -o docli

migrate: docli
	cp docli /usr/local/bin/

clean:
	rm docli

