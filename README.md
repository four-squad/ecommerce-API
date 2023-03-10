# Ecommerce-API

RESTful API for our  [E-Commerce App Project](https://github.com/four-squad/ecommerce-FE)

Where users can post products, buy & sell their goods. It is built with Golang, MySQL, Docker, AWS EC2, etc.

# β‘Features
 - [x] CRUD (Users, Cart, Transactions)
 - [x] Hashing passwordgit 
 - [x] Authentication & Authorization
 - [x] Database Migration
 - [x] Automated deployment with GitHub Actions, DockerHub & AWS EC2
 - [ ] Midtrans Integration as Payment Gateway

# π Folder Structure Pattern

```
βββ .github
β   βββ workflows
β       βββ main.yml
βββ config
β   βββ cloudinary.go
β   βββ config.go
β   βββ db.go
βββ features
β   βββ cart
β   β   βββ data
β   β   β   βββ model.go
β   β   β   βββ query.go
β   β   βββ handler
β   β   β   βββ handler.go
β   β   β   βββ request.go
β   β   β   βββ response.go
β   β   βββ services
β   β   β   βββ service_test.go
β   β   β   βββ service.go
β   β   βββ entity.go
β   βββ product
β   β   βββ data
β   β   β   βββ model.go
β   β   β   βββ query.go
β   β   βββ handler
β   β   β   βββ handler.go
β   β   β   βββ request.go
β   β   β   βββ response.go
β   β   βββ services
β   β   β   βββ service_test.go
β   β   β   βββ service.go
β   β   βββ entity.go
β   βββ user
β   β   βββ data
β   β   β   βββ model.go
β   β   β   βββ query.go
β   β   βββ handler
β   β   β   βββ handler.go
β   β   β   βββ request.go
β   β   β   βββ response.go
β   β   βββ services
β   β   β   βββ service_test.go
β   β   β   βββ service.go
β   β   βββ entity.go
β   βββ transaction
β   β   βββ data
β   β   β   βββ model.go
β   β   β   βββ query.go
β   β   βββ handler
β   β   β   βββ handler.go
β   β   β   βββ request.go
β   β   β   βββ response.go
β   β   βββ services
β   β   β   βββ service_test.go
β   β   β   βββ service.go
β   β   βββ entity.go
βββ helper
β   βββ cloudinary_helper.go
β   βββ dto.go
β   βββ jwt.go
β   βββ pwd.go
β   βββ response.go
βββ mocks
βββ .gitignore
βββ dockerfile
βββ ecommerce.yaml
βββ erd.png
βββ go.mod
βββ go.sum
βββ LICENSE
βββ local.env.example
βββ main.go
βββ README.md
```

# π₯ Open API

Simply [click here](https://app.swaggerhub.com/apis-docs/TECHMIN7_1/Ecommerce/1.0.0#/) to see the details of endpoints we have agreed with our FE team.
# π ERD

![run](./erd.png)

# π» Tools & Stacks
- Backend Stacks :
  - [Golang](https://go.dev/) : Programming Language
  - [Viper](https://github.com/spf13/viper) : Environment Reader
  - [Echo](https://echo.labstack.com/) : Web Framework
  - [JWT](https://jwt.io/) : Authentication & Authorization
  - [GORM](https://gorm.io/) : ORM Library
  - [MySQL](https://gorm.io/) : Database Management System
- Documentation :
  - [Postman](https://www.postman.com/) : API Testing & Documentation
  - [Swagger](https://swagger.io/) : Open API Documentation
- Deployment & Storage :
  - [Ubuntu](https://ubuntu.com/) : Development & Deployment OS
  - [Docker](https://docker.com/) : Containerization
  - [Amazon EC2](https://aws.amazon.com/) : Cloud server
  - [Cloudinary](https://cloudinary.com/) : Store and retrieve images
  - [Cloudflare](https://www.cloudflare.com/) : SSL & Proxy

# π οΈ How to Run Locally

- Clone it

```
$ git clone https://github.com/four-squad/ecommerce-API
```

- Go to directory

```
$ cd ecommerce-api
```

- Create a new database

- Rename `local.env.example` to `local.env`
- Adjust `local.env` as your environment settings

- Run the project

```
$ go run .
```
- Voila! π« 

# π€ Author

- Devangga Wiku :

    [![GitHub](https://img.shields.io/badge/-Deva-black?style=for-the-badge&logo=github&logoColor=white)](https://github.com/DevWiku) 

- Muhammad Habibullah :

    [![GitHub](https://img.shields.io/badge/-Habib-black?style=for-the-badge&logo=github&logoColor=white)](https://github.com/hebobibun) 


<h5>
<p align="center">Built with β€οΈ by Four Squad Β©οΈ 2023</p>
</h5>