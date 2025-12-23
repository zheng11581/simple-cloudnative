all: threads-cpu
threads-cpu: threads-cpu.c
	gcc -o threads-cpu threads-cpu.c -lpthread
image: threads-cpu
	docker build -t threads-cpu .
clean:
	rm -f *.o threads-cpu