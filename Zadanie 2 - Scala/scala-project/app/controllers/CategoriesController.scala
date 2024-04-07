package controllers

import models.category._
import play.api.libs.json.{JsError, JsSuccess, JsValue, Json}
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}
import repositories.{Repository, ResourceNotFound}
import validation.{CategoryDoesNotExist, MalformedInput, ValidationError}

import java.util.UUID
import javax.inject.Inject

class CategoriesController @Inject()(val controllerComponents: ControllerComponents, val categoryRepository: Repository[CategoryModel, UUID]) extends BaseController {
	def addCategory(): Action[JsValue] = Action(parse.json) { implicit request =>
		request.body.validate[NewCategoryModel] match {
			case JsSuccess(newCategoryModel, _) =>
				val result = for {
					newCategory <- categoryRepository.insert(CategoryModel(
						UUID.randomUUID(),
						newCategoryModel.name,
						newCategoryModel.description.getOrElse("")
					))
				} yield newCategory
				result match {
					case Left(error) => error.getResult
					case Right(newCategory) => Ok(Json.toJson(newCategory))
				}

			case JsError(errors) => MalformedInput(errors).getResult
		}
	}

	def getCategories: Action[AnyContent] = Action {
		val result = for {
			categories <- categoryRepository.find
		} yield categories
		result match {
			case Left(error) => error.getResult
			case Right(categories) => Ok(Json.toJson(categories))
		}
	}

	def getCategoryById(categoryId: UUID): Action[AnyContent] = Action {
		val result = for {
			category <- validateCategoryExists(categoryId)
		} yield category
		result match {
			case Left(error) => error.getResult
			case Right(category) => Ok(Json.toJson(category))
		}
	}

	def updateCategory(categoryId: UUID): Action[JsValue] = Action(parse.json) { implicit request =>
		request.body.validate[UpdateCategoryModel] match {
			case JsSuccess(updateCategoryModel, _) =>
				val result = for {
					category <- validateCategoryExists(categoryId)
					updatedCategory <- categoryRepository.update(categoryId, CategoryModel(
						categoryId,
						updateCategoryModel.name.getOrElse(category.name),
						updateCategoryModel.description.getOrElse(category.description)
					))
				} yield updatedCategory
				result match {
					case Left(error) => error.getResult
					case Right(updatedCategory) => Ok(Json.toJson(updatedCategory))
				}

			case JsError(errors) => MalformedInput(errors).getResult
		}
	}

	def deleteCategory(categoryId: UUID): Action[AnyContent] = Action {
		val result = for {
			_ <- categoryRepository.delete(categoryId)
		} yield ()
		result match {
			case Left(error) => error match {
				case ResourceNotFound(_) => NoContent
				case error: Error => error.getResult
			}
			case Right(_) => NoContent
		}
	}

	private def validateCategoryExists(cartId: UUID): Either[ValidationError, CategoryModel] = {
		categoryRepository.findOne(cartId) match {
			case Right(x) => Right(x)
			case Left(_) => Left(CategoryDoesNotExist(cartId))
		}
	}
}