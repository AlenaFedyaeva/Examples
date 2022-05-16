# Swagger

https://blog.logrocket.com/documenting-go-web-apis-with-swag/

# Mux + swag
https://medium.com/@pointgoal/gorilla-mux-101-rk-boot-swagger-ui-a63c8a900c79 

# Install Swag
from official site
https://github.com/swaggo/swag#getting-started

```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest

```
# Init swagger in work directoory 
```
swag init
~/go/bin/swag init
```

# Import generate

in main.go
```
import(
    _ "swagexample/docs"
)	
```
