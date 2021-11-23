.PHONY: api
apiHost=wedo-workflow.niuqiang.vip

api:
	env GOOS=linux go build -o wedo-workflow # && upx -1 -k -v wedo-workflow
	rsync -e 'ssh -p 22' wedo-workflow root@$(apiHost):/data/apps/wedo-workflow/wedo-workflow
	ssh root@$(apiHost) -p 22 "supervisorctl restart wedo-workflow"
	rm -fr wedo-workflow*

run:
	env CONFIG_PATH="config/apiserver.json" go run apiserver.go
