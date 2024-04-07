package models.product

import play.api.libs.functional.syntax.{toApplicativeOps, toFunctionalBuilderOps}
import play.api.libs.json.Reads._
import play.api.libs.json._

case class NewProductModel(var name: String, var description: Option[String])

object NewProductModel {
	implicit def reads: Reads[NewProductModel] = (
		(__ \ "name").read[String](minLength[String](1) keepAnd maxLength[String](255)) and
		(__ \ "description").readNullableWithDefault[String](Option[String](""))
	)(NewProductModel.apply _)
}