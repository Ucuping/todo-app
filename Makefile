# Makefiles are used by the `make` command.
run:
	CompileDaemon -command="./todo-app" -exclude="Makefile,README.md,request.rest,.env.example" -exclude-dir=".git,test"