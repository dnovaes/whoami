## Who am I - Hotsite

### Gif of actual progress:

![crd_working](https://user-images.githubusercontent.com/3251916/92190351-51ba5780-ee37-11ea-98aa-3d5b4b92493e.gif)

### Requisites:

1 - Create a portifolio file containing the credentials for you to connect to your mongodb server

`~/CredentialsConfig/whoami-cert.pem`

2 - define your environment variables:

MONGO_CLUSTER_HOST
MONGO_USERNAME

3  - Install dependencies

```
go mod tidy
go build
```

4 - start server
`go run server.go`
