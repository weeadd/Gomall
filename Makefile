.PHONY: first-init
first-init:
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/auth.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/cart.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/user.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/product.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/payment.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/order.proto
	cwgo server -I idl --type RPC --module Gomall --service Gomall --idl idl/checkout.proto
	go mod tidy
