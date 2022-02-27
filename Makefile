buildweb:
	cd web/depfootprint && npm run export

pushweb: buildweb
	aws s3 sync ./web/depfootprint/out s3://depfootprint-static-assets --acl public-read

buildapi:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda/api cmd/lambda/*.go
	zip -j ./bin/lambda/api.zip ./bin/lambda/api

pushapi: buildapi
	aws s3 cp ./bin/lambda/api.zip s3://depfootprint-api-source/api.zip
