run:
	go run ./main.go

test:
	cat "mail.txt" |while read L; do sleep "2"; echo "$$L"; done | netcat "localhost" "25"

signald:
	docker run -v /Users/lawrencegripper/.signald:/signald finn/signald