
# api 指令请在api目录下执行
.PHONY: gen-api-auth
gen-api-auth:
	cwgo server --type HTTP --idl ../../idl/api/auth_api.proto --service api --module Gomall/app/api -I ../../idl

.PHONY: gen-api-user
gen-api-user:
	cwgo server --type HTTP --idl ../../idl/api/user_api.proto --service api --module Gomall/app/api -I ../../idl


.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module Gomall/rpc_gen -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module Gomall/app/user --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module Gomall/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module Gomall/app/product --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto


.PHONY: gen-auth
gen-auth:
	@cd rpc_gen && cwgo client --type RPC --service auth --module Gomall/rpc_gen -I ../idl --idl ../idl/auth.proto
	@cd app/auth && cwgo server --type RPC --service auth --module Gomall/app/auth --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/auth.proto



