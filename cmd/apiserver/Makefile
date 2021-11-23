.PHONY: api
apiHost=wedo-workflow.niuqiang.vip

api:
	env GOOS=linux go build -ldflags="-s -w" -o wedo-workflow && upx -1 -k -v wedo-workflow
	rsync -e 'ssh -p 22' -c wedo-workflow root@$(apiHost):/data/apps/wedo-workflow
	ssh root@$(apiHost) -p 22 "supervisorctl restart wedo-workflow"
	rm -fr wedo-workflow*
