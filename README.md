# Description
This is a small project written in [golang](https://golang.org/) that implements hexagonal architecture. The project is based on a [YouTube tutorial](https://www.youtube.com/watch?v=ygjTG5qeU2o&list=PLle8fNcWU5atqOGLtviMMHMb6v9dSU_Zc) by [CPoP](https://www.youtube.com/channel/UCAWdZyUJs2RonhbspUR6slg/about), which follows a project structure presented by Kat Zien during [dotGo 2019](https://www.youtube.com/watch?v=vKbVrsMnhDc&t=960s).

The project is a simple CRUD application called "candyshop" with endpoints for ADDING, READING, UPDATING and DELETING candies in a [Postgres](https://www.postgresql.org/) database. The project uses docker-compose to run the application, database and database admin UI, [Adminer](https://www.adminer.org/).

Docker Compose has been set up to persist data added to the Postgres DB while running the application and also to cache the imported golang modules to avoid downloading them every time the application is started.

An ```init.sql``` file has been added that will set up the candies table in the Postgres container the first time that the service is run, so no additional steps are required to set up the database.

---

# Requirements
As the project uses docker-compose, the only requirement for running the application is to have the following installed locally:
* Golang >= 1.16
* Docker/Docker Compose
* [Postman](https://www.postman.com/) (For testing the API endpoints)

---

# Running the application
After cloning this repo, navigate to the root directory of the project and enter the following command in the terminal:

```make dev```

This command will start and initialize the Postgres database, Adminer (DB admin UI) and the candyshop application.

## Cleaning
A command has been added to ssimplify the cleaning up of docker images & volumes:

```make clean```

This command will remove all dangling docker images (if any) and delete all volumes (clear cached postgres and go module imports data)

## Adminer
Adminer is a simple, lightweight admin UI for viewing and editing a number of databases, including Postgres. Follow the following steps to access the database using Adminer once the ```make dev``` command has been entered into the terminal as above:
1. Open a browser and navigate to ```http://localhost:1000```
2. Enter/select the following fields on the login page:

|Field |Value |
|:-----|:-----|
|_System_|**PostgreSQL**|
|_Server_|**app_db**|
|_Username_|**app**|
|_Password_|**dbpassword**|
|_Database_|**candiesDB**|

---

# Endpoints
The application has the following endpoints for adding, reading, updating and deleting candies form the database:
## Adding
### Adds a single candy to the database.

_URL_: ```http://localhost:5000/api/candy```
_Method_: ```POST```
_Request body_:

```json
{ 
    "category": "SOME CATEGORY NAME",
    "name": "SOME CANDY NAME", 
    "price": "SOME PRICE (NUMERIC)"
}
```
## Reading
### Returns a list of all candies stored in the database.
_URL_: ```http://localhost:5000/api/candies```
_Method_: ```GET```

## Updating
### Updates the [category] of an existing candy.
_URL_: ```localhost:5000/api/candy/{id}/category```
_Method_: ```PATCH```
_Request body_:
```json
{ 
    "category": "SOME CATEGORY NAME"
}
```
_The_ ```{id}``` _param refers to the auto generated uuid(v4) id for a specific candy stored in the database._
### Updates the [name] of an existing candy.
_URL_: ```localhost:5000/api/candy/{id}/name```
_Method_: ```PATCH```
_Request body_:
```json
{ 
    "name": "SOME CANDY NAME"
}
```
_The_ ```{id}``` _param refers to the auto generated uuid(v4) id for a specific candy stored in the database._
### Updates the [price] of an existing candy.
_URL_: ```localhost:5000/api/candy/{id}/price```
_Method_: ```PATCH```
_Request body_:
```json
{ 
    "price": "SOME PRICE (NUMERIC)"
}
```
_The_ ```{id}``` _param refers to the auto generated uuid(v4) id for a specific candy stored in the database._

## Deleting
### Deletes a specific candy from the database
_URL_: ```http://localhost:5000/api/candy/{id}```
_Method_: ```DELETE```

_The_ ```{id}``` _param refers to the auto generated uuid(v4) id for a specific candy stored in the database._

---

# Todo
* Add http request payload validation to all endpoints.
* Add endpoint to get information for a specific candy.
* Implement logging service.
* Implement error handling service.

