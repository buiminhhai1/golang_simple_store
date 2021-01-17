# Some libs have require
* system: macos

* need instal golang-migrate

* run command to generates migrate db: ```migrate create -ext sql -dir db/migration -seq init_schema```

* using sqlc to CRUD database, must install on local machine: ```brew install sqlc``` for more detail https://github.com/kyleconroy/sqlc
 
* To run test code you need to install pg driver
```go get github.com/lib/pg```

* To check test result using testify package
``` go get github.com/stretchr/testify```

* You can test code coverage ```make test```

* In this project using ```go gin``` framework to make server http more easy for more detail https://github.com/gin-gonic/gin

* Config load env from app.env file using ```viper``` for more detail https://github.com/spf13/viper

* Using DB stubs: GOMOCK to generate and build stubs that returns hard-coded values, for more detail https://github.com/golang/mock

* You can start server by the following step below
0. using ```make install`` to install dependencies package
1. create a database using ```make createdb``` or drop database using ```make dropdb```
2. run migrate to generate all tables with them relations you defined on folder db/migrate using ```make migrateup``` or drop them all using ```make migratedown```
3. create func to access and modify database with ```make sqlc```
4. start server using ```make server```