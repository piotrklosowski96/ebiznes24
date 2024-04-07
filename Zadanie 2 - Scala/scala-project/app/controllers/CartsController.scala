package controllers

import models.cart._
import models.product.ProductModel
import play.api.libs.json.{JsError, JsSuccess, JsValue, Json}
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}
import repositories.{Repository, ResourceNotFound}
import validation._

import java.util.UUID
import javax.inject.Inject
import scala.collection.mutable.ListBuffer

class CartsController @Inject()(
		val controllerComponents: ControllerComponents,
		val cartRepository: Repository[CartModel, UUID],
		val productsRepository: Repository[ProductModel, UUID]
	) extends BaseController {

	def createCart(): Action[JsValue] = Action(parse.json) { implicit request =>
		request.body.validate[NewCartModel] match {
			case JsSuccess(newCartModel, _) =>
				val result = for {
					_ <- validateProductsExists(newCartModel.products.getOrElse(ListBuffer[UUID]()))
					createdCart <- cartRepository.insert(CartModel(
						UUID.randomUUID(),
						newCartModel.name,
						newCartModel.products.getOrElse(ListBuffer[UUID]()),
					))
				} yield createdCart
				result match {
					case Left(error) => error.getResult
					case Right(createdCart) => Ok(Json.toJson(createdCart))
				}

			case JsError(errors) => MalformedInput(errors).getResult
		}
	}

	def getAllCarts: Action[AnyContent] = Action {
		val result = for {
			carts <- cartRepository.find
		} yield carts
		result match {
			case Left(error) => error.getResult
			case Right(carts) => Ok(Json.toJson(carts))
		}
	}

	def getCartById(cartId: UUID): Action[AnyContent] = Action {
		val result = for {
			cart <- validateCartExists(cartId)
		} yield cart
		result match {
			case Left(error) => error.getResult
			case Right(cart) => Ok(Json.toJson(cart))
		}
	}

	def updateCart(cartId: UUID): Action[JsValue] = Action(parse.json) { implicit request =>
		request.body.validate[UpdateCartModel] match {
			case JsSuccess(updateCartModel, _) =>
				val result = for {
					cart <- validateCartExists(cartId)
					_ <- validateProductsExists(updateCartModel.products.getOrElse(ListBuffer[UUID]()))
					updatedCart <- cartRepository.update(cart.id, CartModel(
						cart.id,
						updateCartModel.name.getOrElse(cart.name),
						updateCartModel.products.getOrElse(cart.products)
					))
				} yield updatedCart
				result match {
					case Left(error) => error.getResult
					case Right(updatedCart) => Ok(Json.toJson(updatedCart))
				}

			case JsError(errors) => MalformedInput(errors).getResult
		}
	}

	def deleteCart(cartId: UUID): Action[AnyContent] = Action {
		val result = for {
			_ <- cartRepository.delete(cartId)
		} yield ()
		result match {
			case Left(error) => error match {
				case ResourceNotFound(_) => NoContent
				case error: Error => error.getResult
			}
			case Right(_) => NoContent
		}
	}

	def addProductToCart(cartId: UUID, productId: UUID): Action[AnyContent] = Action {
		val result = for {
			cart <- validateCartExists(cartId)
			product <- validateProductExists(productId)
			updatedCart <- cartRepository.update(cart.id, CartModel(
				cart.id,
				cart.name,
				cart.products.addOne(product.id)
			))
		} yield updatedCart
		result match {
			case Left(error) => error.getResult
			case Right(updatedCart) => Ok(Json.toJson(updatedCart))
		}
	}

	def getProductsInCart(cartId: UUID): Action[AnyContent] = Action {
		val validationResult = for {
			cart <- validateCartExists(cartId)
		} yield cart
		validationResult match {
			case Left(error) => error.getResult
			case Right(cart) => Ok(Json.toJson(cart.products))
		}
	}

	def updateProductsInCart(cartId: UUID): Action[JsValue] = Action(parse.json) { implicit request =>
		request.body.validate[ListBuffer[UUID]] match {
			case JsSuccess(productIds, _) =>
				val result = for {
					cart <- validateCartExists(cartId)
					productIds <- validateProductsExists(productIds)
					updatedCart <- cartRepository.update(cart.id, CartModel(
						cart.id,
						cart.name,
						productIds,
					))
				} yield updatedCart
				result match {
					case Left(error) => error.getResult
					case Right(updatedCart) => Ok(Json.toJson(updatedCart))
				}

			case JsError(errors) => MalformedInput(errors).getResult
		}
	}

	def deleteProductFromCart(cartId: UUID, productId: UUID): Action[AnyContent] = Action {
		val result = for {
			cart <- validateCartExists(cartId)
			updatedCart <- cartRepository.update(cart.id, CartModel(
				cart.id,
				cart.name,
				cart.products.filterNot(p => p == productId)
			))
		} yield updatedCart
		result match {
			case Left(error) => error.getResult
			case Right(_) => NoContent
		}
	}

	private def validateCartExists(cartId: UUID): Either[ValidationError, CartModel] = {
		cartRepository.findOne(cartId) match {
			case Right(x) => Right(x)
			case Left(_) => Left(CartDoesNotExist(cartId))
		}
	}

	private def validateProductsExists(productIds: ListBuffer[UUID]): Either[ValidationError, ListBuffer[UUID]] = {
		productIds.filter(productsRepository.findOne(_).isLeft).toList match {
			case Nil => Right(productIds)
			case x => Left(ProductsDoesNotExist(x))
		}
	}

	private def validateProductExists(productId: UUID): Either[ValidationError, ProductModel] = {
		productsRepository.findOne(productId) match {
			case Right(product) => Right(product)
			case Left(_) => Left(ProductDoesNotExist(productId))
		}
	}
}
