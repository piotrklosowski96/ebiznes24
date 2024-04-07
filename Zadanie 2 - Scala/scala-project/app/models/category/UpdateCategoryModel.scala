package models.category

import play.api.libs.functional.syntax.{toApplicativeOps, toFunctionalBuilderOps}
import play.api.libs.json.Reads._
import play.api.libs.json._

case class UpdateCategoryModel(var name: Option[String], var description: Option[String])

object UpdateCategoryModel {
	implicit def reads: Reads[UpdateCategoryModel] = (
		(JsPath \ "name").readNullable[String](minLength[String](1) keepAnd maxLength[String](255)) and
			(JsPath \ "description").readNullable[String]
		)(UpdateCategoryModel.apply _)
}