CFLAGS=
CC=8g
LD=8l
NAME=swif

all:
	$(CC) *.go $(CFLAGS)
	$(LD) -o $(NAME) *.8 
run:
	./$(NAME)
clean:
	@-rm *.8 $(NAME)
