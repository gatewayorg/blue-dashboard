module github.com/gatewayorg/blue-dashboard

go 1.15

replace github.com/gatewayorg/blue-dashboard/api v0.0.0 => ./api

require (
	github.com/Ankr-network/kit v1.8.7
	github.com/cornelk/hashmap v1.0.1
	github.com/gatewayorg/blue-dashboard/api v0.0.0
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.16
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
)
