.PHONY: gen-api
gen-api:
	cwgo server --type HTTP --idl ../../idl/api/auth_api.proto --service api --module Gomall/app/api -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module Gomall/rpc_gen -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module Gomall/app/user --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module Gomall/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/pr product && cwgo server --type RPC --service product --module Gomall/app/product --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module Gomall/rpc_gen -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module Gomall/app/cart --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto



