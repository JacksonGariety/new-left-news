install:
	go get github.com/tools/godep
	go get bitbucket.org/liamstask/goose/cmd/goose
	godep restore

install_dev:
	go get github.com/tools/godep
	go get bitbucket.org/liamstask/goose/cmd/goose
	godep restore
	go get github.com/pilu/fresh
	go get github.com/stretchr/testify

watch:
	env env=development port=8080 session_hash=dev_hash base_path=$$GOPATH/src/github.com/JacksonGariety/new-left-news fresh

test:
	dropdb new_left_news_test --if-exists
	createdb new_left_news_test
	goose -env test up
	env env=test port=8080 session_hash=test_hash base_path=$$GOPATH/src/github.com/JacksonGariety/new-left-news/ go test ./app/...
	dropdb new_left_news_test

run:
	go build
	env env=production port=8002 session_hash=needed_hash base_path=$$GOPATH/src/github.com/JacksonGariety/new-left-news ./new-left-news
