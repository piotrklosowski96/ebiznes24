package validation

import play.api.libs.json.{JsPath, Json, JsonValidationError}
import play.api.mvc.Result
import play.api.mvc.Results.BadRequest

import scala.collection.Seq

case class MalformedInput(errors: Seq[(JsPath, Seq[JsonValidationError])]) extends ValidationError {
	override def getResult: Result = BadRequest(Json.toJson(
		errors.groupMap(
			k => k._1.toString()
		)(
			v => v._2.collectFirst(p => Json.obj(p.message -> p.args.toString()))
		)
	))
}
