# short-it
URL shortner in go

## getting started

### setup
1. clone
```
git clone git@github.com:thebluespecs/short-it.git
```
2. installing dependencies
> [!note] building or running the server takes care of this in golang.

### server
1. Build, the server uses make as a build engine. This takes care of the dependencies of the project by itself.
```
make build
```
2. Running the server
```
make run 
```
3. Stop the process. While a simple `CTRL + c` should kill the process. The supporting container in the background
would need the below command to stop
```
make stop
```

## Functionality and endpoints
> [!note] assuming we are running on loopback adresses with port 8000
1. shorten a URL
```
curl --location 'localhost:8000/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://github.com/"
}'
```
2. get info about the shortened url
```
curl --location 'localhost:8000/:some_id/info'
```
3. redirection
```
curl --location 'localhost:8000/:some_id'
```
