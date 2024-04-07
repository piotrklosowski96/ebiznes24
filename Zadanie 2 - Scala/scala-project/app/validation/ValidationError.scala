package validation

import play.api.mvc.Result
import play.api.mvc.Results.InternalServerError

trait Error {
	def getResult: Result = InternalServerError
}

trait ValidationError extends Error {}