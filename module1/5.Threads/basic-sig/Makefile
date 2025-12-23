all: sigterm sigkill

sigterm: reg-sigterm.c
	gcc -o reg-sigterm reg-sigterm.c

sigkill: reg-sigkill.c
	gcc -o reg-sigkill reg-sigkill.c

clean:
	rm -f *.o reg-sigterm reg-sigkill