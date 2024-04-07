package models.product

import play.api.libs.json.{Json, OFormat}
import repositories.WithIndex

import java.util.UUID

case class ProductModel(var id: UUID, var name: String, var description: String) extends WithIndex[UUID]

object ProductModel {
	implicit val productJson: OFormat[ProductModel] = Json.format[ProductModel]
}