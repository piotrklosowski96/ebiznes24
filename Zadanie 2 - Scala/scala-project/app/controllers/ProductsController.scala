package controllers

import models.product._
import play.api.libs.json.{JsError, JsSuccess, JsValue, Json}
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}
import repositories.{DuplicatedIndexError, Repository, ResourceNotFound}
import validation.{MalformedInput, ProductDoesNotExist, ValidationError}

import java.util.UUID
import javax.inject.Inject

class ProductsController @Inject()(val controllerComponents: ControllerComponents, val productsRepository: Repository[ProductModel, UUID]) extends BaseController {
  def addProduct(): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[NewProductModel] match {
      case JsSuccess(newProductModel, _) =>
        val result = for {
          newProduct <- productsRepository.insert(ProductModel(
            UUID.randomUUID(),
            newProductModel.name,
            newProductModel.description.getOrElse(""),
          ))
        } yield newProduct
        result match {
          case Left(error) => error.getResult
          case Right(newProduct) => Ok(Json.toJson(newProduct))
        }

      case JsError(errors) => MalformedInput(errors).getResult
    }
  }

  def getProducts: Action[AnyContent] = Action {
    val result = for {
      products <- productsRepository.find
    } yield products
    result match {
      case Left(error) => error.getResult
      case Right(products) => Ok(Json.toJson(products))
    }
  }

  def getProductById(productId: UUID): Action[AnyContent] = Action {
    val result = for {
      product <- validateProductExists(productId)
    } yield product
    result match {
      case Left(error) => error.getResult
      case Right(product) => Ok(Json.toJson(product))
    }
  }

  def updateProduct(productId: UUID): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[UpdateProductModel] match {
      case JsSuccess(updateProductModel, _) =>
        val result = for {
          product <- validateProductExists(productId)
          updatedProduct <- productsRepository.update(product.id, ProductModel(
            product.id,
            updateProductModel.name.getOrElse(product.name),
            updateProductModel.description.getOrElse(product.description)
          ))
        } yield updatedProduct
        result match {
          case Left(error) => error.getResult
          case Right(updatedProduct) => Ok(Json.toJson(updatedProduct))
        }

      case JsError(errors) => MalformedInput(errors).getResult
    }
  }

  def deleteProduct(productId: UUID): Action[AnyContent] = Action {
    val result = for {
      _ <- productsRepository.delete(productId)
    } yield ()
    result match {
      case Left(error) => error match {
        case ResourceNotFound(_) => NoContent
        case error: Error => error.getResult
      }
      case Right(_) => NoContent
    }
  }

  private def validateProductExists(productId: UUID): Either[ValidationError, ProductModel] = {
    productsRepository.findOne(productId) match {
      case Right(product) => Right(product)
      case Left(_) => Left(ProductDoesNotExist(productId))
    }
  }
}
