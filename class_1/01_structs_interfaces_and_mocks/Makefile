image:
	DOCKER_BUILDKIT=1 docker build --ssh default -t 00b_base_lesson_with_server -f Dockerfile .

agentlist:
	docker run --rm -ti  -v /run/host-services/ssh-auth.sock:/run/host-services/ssh-auth.sock -e SSH_AUTH_SOCK=/run/host-services/ssh-auth.sock -v "${HOME}"/.ssh:/root/.ssh golang ssh-add -L

agentload:
	docker run --rm  -ti -v /run/host-services/ssh-auth.sock:/run/host-services/ssh-auth.sock -e SSH_AUTH_SOCK=/run/host-services/ssh-auth.sock -v "${HOME}"/.ssh:/root/.ssh golang ssh-add

