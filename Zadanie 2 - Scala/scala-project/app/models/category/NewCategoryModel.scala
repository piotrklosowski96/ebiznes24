package models.category

import play.api.libs.functional.syntax.{toApplicativeOps, toFunctionalBuilderOps}
import play.api.libs.json.Reads._
import play.api.libs.json._

case class NewCategoryModel(var name: String, var description: Option[String])

object NewCategoryModel {
	implicit def reads: Reads[NewCategoryModel] = (
		(__ \ "name").read[String](minLength[String](1) keepAnd maxLength[String](255)) and
			(__ \ "description").readNullableWithDefault[String](Option[String](""))
		)(NewCategoryModel.apply _)
}