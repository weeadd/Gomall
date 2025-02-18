.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module Gomall/rpc_gen -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module Gomall/app/user --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-auth
gen-auth:
	@cd rpc_gen && cwgo client --type RPC --service auth --module Gomall/rpc_gen -I ../idl --idl ../idl/auth.proto
	@cd app/auth && cwgo server --type RPC --service auth --module Gomall/app/auth --pass "-use Gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/auth.proto



