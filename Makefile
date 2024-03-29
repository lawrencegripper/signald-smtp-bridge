run:
	go run ./main.go

test:
	go test --count=1 .

signald:
	docker run -v /Users/lawrencegripper/.signald:/signald finn/signald

build:
	go build .

publish:
	echo "v0.1.$$(date +%s)" > tag.txt
	docker build -t ghcr.io/lawrencegripper/signald-smtp-bridge:$$(cat tag.txt) -t ghcr.io/lawrencegripper/signald-smtp-bridge:v0.1 -t ghcr.io/lawrencegripper/signald-smtp-bridge:latest .
	docker push ghcr.io/lawrencegripper/signald-smtp-bridge:$$(cat tag.txt)
	docker push ghcr.io/lawrencegripper/signald-smtp-bridge:latest
	docker push ghcr.io/lawrencegripper/signald-smtp-bridge:v0.1

deploy: publish
	kubectl apply -f ./Deployment.yaml