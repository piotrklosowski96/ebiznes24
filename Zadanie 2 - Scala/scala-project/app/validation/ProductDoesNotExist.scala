package validation

import play.api.libs.json.Json
import play.api.mvc.Result
import play.api.mvc.Results.NotFound

import java.util.UUID

case class ProductDoesNotExist(productId: UUID) extends ValidationError {
	override def getResult: Result = NotFound(Json.obj(
		"message" -> s"product (id: '$productId') does not exist",
	))
}