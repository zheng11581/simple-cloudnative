all: compile image
compile: 
	gcc -o fork-child fork-child.c

image: compile
	docker build -t registry/fwd_sig:v1 .
clean: 
	rm -f *.o fork-child
	docker rmi registry/fwd_sig:v1