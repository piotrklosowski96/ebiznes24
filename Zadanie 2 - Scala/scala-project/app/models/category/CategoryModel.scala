package models.category

import play.api.libs.json.{Json, OFormat}
import repositories.WithIndex

import java.util.UUID

case class CategoryModel(var id: UUID, var name: String, var description: String) extends WithIndex[UUID]

object CategoryModel {
	implicit val categoryJson: OFormat[CategoryModel] = Json.format[CategoryModel]
}