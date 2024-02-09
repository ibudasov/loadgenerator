# Load generator

This is a simple load generation script, which execures `scenario.go` of user behavior taking in account multiple parallel users. 

## Use

```sh
# run the load from your machine. `-h` is the host without the path, `-u` is the amount of users
go run . -h "http://localhost" -u 2 2>&1

# also tere is a conteinerized way:
docker build -t loadgenerator .
docker run -d -e FRONTEND_ADDR=frontendhost -e USERS=1 --name loadgenerator-running loadgenerator 
docker logs loadgenerator-running

# testing before pushing. It would run tests in dry-run mode, without making HTTP requests
go test
```