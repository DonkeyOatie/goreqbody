# goreqbody
Code showing error in request body parsing

## Build
go get github.com/julienschmidt/httprouter
go build

## Run
./goreqbody

## Output

### go 1.8.7
curl -H "Content-Type: application/json" -X POST -d '{"jsonParam":"xyz"}' http://localhost:8000/dosomething

localStr not empty

#### logs 
localStr: xyz

### go 1.9.4
curl -H "Content-Type: application/json" -X POST -d '{"jsonParam":"xyz"}' http://localhost:8000/dosomething

localStr empty

#### logs 
localStr: 
