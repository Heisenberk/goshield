#Makefile Goshield

GOSHIELD=github.com/Heisenberk/goshield

run: compil
	./goshield
	
all: compil run 

compil:
	go build $(GOSHIELD)
	
test: 
	go test $(GOSHIELD)/command
	go test $(GOSHIELD)/crypto
	
install: 
	go build $(GOSHIELD) 
	mv ./goshield /usr/local/bin
	
clean: 
	rm -f goshield
	
doc: 
	godoc -http=:8080
	firefox http://localhost:8080/pkg/github.com/Heisenberk/goshield/
	
	
