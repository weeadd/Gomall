# api 指令请在api目录下执行
.PHONY: gen-api-auth
gen-api-auth:
	cwgo server --type HTTP --idl ../../idl/api/auth_api.proto --service api --module Gomall/app/api -I ../../idl

.PHONY: gen-api-user
gen-api-user:
	cwgo server --type HTTP --idl ../../idl/api/user_api.proto --service api --module Gomall/app/api -I ../../idl


