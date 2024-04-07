package models.product

import play.api.libs.functional.syntax.{toApplicativeOps, toFunctionalBuilderOps}
import play.api.libs.json.{JsPath, Reads, __}
import play.api.libs.json.Reads.{maxLength, minLength}

case class UpdateProductModel(var name: Option[String], var description: Option[String])

object UpdateProductModel {
	implicit def reads: Reads[UpdateProductModel] = (
		(JsPath \ "name").readNullable[String](minLength[String](1) keepAnd maxLength[String](255)) and
			(JsPath \ "description").readNullable[String]
		)(UpdateProductModel.apply _)
}