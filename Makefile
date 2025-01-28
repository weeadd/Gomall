

clean:
	rm -rf rpc
	rm -rf go.mod
	rm -rf kitex_gen


init:
	kitex -module Gomall -I idl idl/auth.proto
	kitex -module Gomall -I idl idl/user.proto
	kitex -module Gomall -I idl idl/product.proto
	kitex -module Gomall -I idl idl/cart.proto
	kitex -module Gomall -I idl idl/order.proto
	kitex -module Gomall -I idl idl/checkout.proto
	kitex -module Gomall -I idl idl/payment.proto

dir: init
	mkdir -p rpc/auth rpc/user rpc/product rpc/cart rpc/order rpc/checkout rpc/payment


