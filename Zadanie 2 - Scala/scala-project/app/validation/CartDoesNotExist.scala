package validation

import play.api.libs.json.Json
import play.api.mvc.Result
import play.api.mvc.Results.NotFound

import java.util.UUID

case class CartDoesNotExist(cartId: UUID) extends ValidationError {
	override def getResult: Result = NotFound(Json.obj(
		"message" -> s"cart (id: '$cartId') does not exist"
	))
}
