# GORM CRUD

https://gorm.io/docs/connecting_to_the_database.html



```
# Start your Postgres with Docker
docker-compose up

# Start your go module
go mod init github.com/FranciscoOrtizCastillo/gorm-example

go mod tidy

go run main.go


git init
git add .
git commit -m "first commit"   


docker build --tag docker-gorm-example  .

docker run docker-gorm-example

#Stop the container(s) using the following command:
docker-compose down

#Delete all containers using the following command:
docker rm -f $(docker ps -a -q)

#Delete all volumes using the following command:
docker volume rm $(docker volume ls -q)

#Restart the containers using the following command:
docker-compose up -d

```



