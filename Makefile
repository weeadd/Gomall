
# api 指令请在api目录下执行
.PHONY: gen-api-auth
gen-api-auth:
	cwgo server --type HTTP --idl ../../idl/api/auth_api.proto --service api --module Gomall/app/api -I ../../idl

.PHONY: gen-api-user
gen-api-user:
	cwgo server --type HTTP --idl ../../idl/api/user_api.proto --service api --module Gomall/app/api -I ../../idl

.PHONY: gen-api-cart
gen-api-cart:
	cwgo server --type HTTP --idl ../../idl/api/cart_api.proto --service api --module Gomall/app/api -I ../../idl

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


.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module Gomall/rpc_gen -I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module Gomall/app/order --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module Gomall/rpc_gen -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module Gomall/app/cart --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module Gomall/rpc_gen -I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module Gomall/app/payment --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto
	
.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module Gomall/rpc_gen -I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module Gomall/app/checkout --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto

