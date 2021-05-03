# ml-x-men

## Application Architecture

I used a clean architecture called `Hexagonal Architecture`. It allows an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases

Some info about it:
- https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)
- https://fideloper.com/hexagonal-architecture
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

Github repo example:
- https://github.com/err0r500/go-realworld-clean

## Set up

## Database

### Dockerized:
If you want run the mysql server provided by a docker image see the detail below:

Inside the dockerized folder create an `.env` file with the following environment variables:
```
DB_HOST_PORT=3308
MYSQL_DATABASE=ml_x_men
MYSQL_ROOT_PASSWORD=root
```
Now, run the command below:
```
cd dockerized/ && docker-compose up
```
### Note

If you change your DB data, please update the `config/config.yaml` file with the new one


### Database script

If you don't want to use docker, you can find the sql with the data struct in the `./dockerized/db/dump.sql`

## Application

Run the main package:

```
go run cmd/app/main.go 
```


## Nivel 1 exercise 


