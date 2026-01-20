.PHONY: docker
docker:
	@rm golang || true
	@go mod tidy
	@GOOS=linux GOARCH=arm64 go build -tags=k8s -o golang .
	@docker image rm -f chixuy/golang-record:v0.0.1
	@docker buildx build --platform linux/arm64  -t chixuy/golang-record:v0.0.1 .
	@minikube image load chixuy/golang-record:v0.0.1
	@kubectl apply -f golang-deployment.yaml
	@kubectl get pods | grep record
	@kubectl apply -f golang-service.yaml
	@kubectl get deployment

# 执行日志
# kubectl logs $(kubectl get pods -l app=golang-record -o jsonpath="{.items[0].metadata.name}")