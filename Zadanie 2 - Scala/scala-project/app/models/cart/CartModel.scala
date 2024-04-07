package models.cart

import play.api.libs.json.{Json, OFormat}
import repositories.WithIndex

import java.util.UUID
import scala.collection.mutable

case class CartModel(var id: UUID, var name: String, var products: mutable.ListBuffer[UUID]) extends WithIndex[UUID] {}

object CartModel {
	implicit val cartJson: OFormat[CartModel] = Json.format[CartModel]
}