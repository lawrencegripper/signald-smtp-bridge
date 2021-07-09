run:
	go run ./main.go

test:
	go test --count=1 .

signald:
	docker run -v /Users/lawrencegripper/.signald:/signald finn/signald