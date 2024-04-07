package models.cart

import play.api.libs.functional.syntax.{toApplicativeOps, toFunctionalBuilderOps}
import play.api.libs.json.Reads._
import play.api.libs.json._

import java.util.UUID
import scala.collection.mutable

case class NewCartModel(var name: String, var products: Option[mutable.ListBuffer[UUID]]) {}

object NewCartModel {
	implicit def newCartModelReads: Reads[NewCartModel] = (
		(JsPath \ "name").read[String](minLength[String](1) keepAnd maxLength[String](255)) and
		(JsPath \ "products").readNullableWithDefault[mutable.ListBuffer[UUID]](None)
	)(NewCartModel.apply _)
}