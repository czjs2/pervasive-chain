module pervasive-chain

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/gorilla/websocket v1.4.2
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/tebeka/strftime v0.1.4 // indirect
	github.com/tidwall/gjson v1.6.3
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42
)

replace xjrwws.com/ws => ./extern/xjrwwebsocket/ws
