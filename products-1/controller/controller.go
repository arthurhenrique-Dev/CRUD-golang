package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController{
	return productController{
		productUseCase: usecase,
	}
}
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}
func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err:= ctx.BindJSON(&product)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return	
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

		if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return	
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
func (p *productController) GetProductsById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if (id == ""){
		response := model.Response{
			Message: "Id não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

		if err != nil{

		response := model.Response{
			Message: "Id precisa ser numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductsById(productId)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil{
			response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}