module github.com/omniboost/go-exactglobe-webservices

go 1.14

require (
	github.com/Azure/go-ntlmssp v0.0.0-20191115210519-2b2be6cc8ed4
	github.com/gorilla/schema v1.1.0
	github.com/joho/godotenv v1.3.1-0.20181120194748-69ed1d913aa8
	github.com/kr/pretty v0.1.0 // indirect
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200204104054-c9f3fb736b72 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/guregu/null.v3 v3.5.0
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
