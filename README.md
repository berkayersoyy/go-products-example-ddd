# <div align="center">go-products-example-ddd</div>
</br>
DDD version of Go application created with Gin and Gorm basicly. <br>
Previous version of application -> https://github.com/berkayersoyy/go-products-example
</br>

# Project Structure

```
pkg
├── application
|   ├── util
|   └── services *
├── domain                    
└── infrastructure
|   ├── dynamodb
|   ├── mysql
|   └── redis
|
└── presentation
    ├── http
    └── middleware
```


# 🚀 Building and Running for Production

1. Follow these steps to get your development environment set up:

2. At the root directory which include docker-compose.yml files, run below command:

        docker-compose up -d --build

3. You can launch application as below url:

* http://localhost:8080/swagger/index.html
