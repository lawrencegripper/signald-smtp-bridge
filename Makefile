run:
	go run ./main.go

test:
	cat "mail.txt" |while read L; do sleep "2"; echo "$$L"; done | netcat "localhost" "25"