include config.mk

all:
	cd swif ; $(MAKE) $(MFLAGS)
	$(CC) $(CFLAGS) -I ./swif/ *.go
	$(LD) -o $(NAME) -L swif *.$(ARCH) 
run:
	./$(NAME)
clean:
	cd swif; $(MAKE) $(MFLAGS) clean
	@-rm *.$(ARCH) $(NAME) 
