## gin

gin.Context を使うことで、URL に付随したパラメータの取得や POST で送信されたデータの取得などを行うことが可能

## layer

repository -> usecase -> controller -> router

domain で usecase のインターフェースを定義したら、usecase ファイルのインターフェース内で定義したメソッド群をポインタレシーバにする usecase を作成

### Router

First of all, the request comes to the Router.

divided into two routers as follows:

- Public Router: All the public APIs should go through this router.

- Protected Router: All the private APIs should go through this router.

In between both routers, a middleware gets added to check the validity of the access token

the private request with the invalid access token should not reach the protected router at all

Then, it gets distributed to the corresponding router

the router will call its corresponding controller

to call the controller, we need the usecase, as the controller is dependent on the usecase

We also need a repository as the usecase is dependent on the repository

Now, we have the repository and we pass it to the usecase

After that, we have the usecase, we pass it to the controller

Finally, our controller is ready to use inside the router

Each request to the backend is eventually executed by a controller

A list of routes is defined which maps a given request to a controller and an action

### Controller

validate the data present inside the request

If anything is invalid, it returns a "400 Bad Request" as the error response

If everything is valid inside the request, it will call the usecase layer to perform an operation

### Usecase

uses the repository layer to perform an operation. It is completely up to the repository how it is going to perform an operation

### Repository

The repository layer is free to choose any database, in fact, it can call any other independent services based on the requirement

In the project, the repository layer makes the database query for performing operations asked by the Usecase layer

### Domain

In the domain layer,

- Models for request, and response.

- Entities for the database.

- Interfaces for usecases, and repositories.

Domain, model, and entity get used in the controller, usecase, and repository

## authentication

jwt_custom.go -> tokenutil.go -> jwt_auth_middleware.go

create jwt_custom.go

create .env

create tokenutil.go

tokenutil.go have following func

- CreateAccessToken

takes three parameters: user, access secret, and expiry

creates a token by encoding the payload that consists of the user Name and ID with the given expiry time signed with the given access secret

- CreateRefreshToken

takes three parameters: user, refresh secret, and expiry.

creates a token by encoding the payload that consists of the user ID with the given expiry time signed with the given refresh secret.

- IsAuthorized

equestToken secret

function does the task of checking if the given token is authorized or not

- ExtractIDFromToken

decodes and extracts the ID that was encoded while creating the token

jwt_auth_middleware.go

takes the access secret key as an input parameter

```sh

brew services start mongodb-community

brew services stop mongodb-community

```
