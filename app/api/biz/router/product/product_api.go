// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	product "Gomall/app/api/biz/handler/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_product := root.Group("/product", _productMw()...)
		_product.GET("/get", append(_getproductMw(), product.GetProduct)...)
		_product.GET("/search", append(_searchproductMw(), product.SearchProduct)...)
	}
}
