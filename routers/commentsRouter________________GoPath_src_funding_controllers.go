package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "GetAddresses",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "DeleteAddress",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "NewAddress",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "UpdateAddress",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "AddCart",
            Router: `/addCart`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartDel",
            Router: `/cartDel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartEdit",
            Router: `/cartEdit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartList",
            Router: `/cartList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "DelCartChecked",
            Router: `/delCartChecked`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "EditCheckAll",
            Router: `/editCheckAll`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "AddOrder",
            Router: `/addOrder`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "OrderInIds",
            Router: `/orderInIds`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "OrderList",
            Router: `/orderList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOrderListToSeller",
            Router: `/orderListToSeller`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "OrderPay",
            Router: `/orderPay`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "SendOutOrder",
            Router: `/sendOutOrder`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetCheckoutPkgInfo",
            Router: `/checkoutPkgInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductWithPkg",
            Router: `/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetHome",
            Router: `/home`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductByPage",
            Router: `/productList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductTypeList",
            Router: `/typeList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetVerifyProduct",
            Router: `/verify`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "VerifyProduct",
            Router: `/verify/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductMangerController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductMangerController"],
        beego.ControllerComments{
            Method: "GetPkgListByProductId",
            Router: `/pkgListByProductId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductMangerController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductMangerController"],
        beego.ControllerComments{
            Method: "GetProductById",
            Router: `/productById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductMangerController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductMangerController"],
        beego.ControllerComments{
            Method: "SaveProduct",
            Router: `/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductMangerController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductMangerController"],
        beego.ControllerComments{
            Method: "SaveProductPackage",
            Router: `/saveProductPackage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "GetUserById",
            Router: `/id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "UploadImage",
            Router: `/uploadImage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "OptionsUploadImage",
            Router: `/uploadImage`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
