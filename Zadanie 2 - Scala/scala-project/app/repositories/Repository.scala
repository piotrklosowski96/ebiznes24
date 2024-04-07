package repositories

import javax.inject.Singleton
import scala.collection.mutable

trait WithIndex[IndexType] {
	var id: IndexType
}

@Singleton
class Repository[ResourceType <: WithIndex[IndexType], IndexType] {
	private val resources = new mutable.ListBuffer[ResourceType]()

	def insert(resource: ResourceType): Either[RepositoryError, ResourceType] = {
		resources.find(_.id == resource.id) match {
			case Some(resource) => Left(DuplicatedIndexError(resource.id))
			case None =>
				resources.addOne(resource)
				Right(resource)
		}
	}

	def find: Either[RepositoryError, Iterable[ResourceType]] = {
		Right(resources.toList)
	}

	def findOne(id: IndexType): Either[RepositoryError, ResourceType] = {
		resources.find(_.id == id) match {
			case Some(resource) => Right(resource)
			case None => Left(ResourceNotFound(id))
		}
	}

	def update(id: IndexType, resource: ResourceType): Either[RepositoryError, ResourceType] = {
		resources.zipWithIndex.find(((resource: ResourceType, _: Int) => resource.id == id).tupled) match {
			case Some((_, index)) =>
				resources.update(index, resource)
				Right(resource)
			case None => Left(ResourceNotFound(id))
		}
	}

	def delete(id: IndexType): Either[RepositoryError, Boolean] = {
		resources.zipWithIndex.find(((resource: ResourceType, _: Int) => resource.id == id).tupled) match {
			case Some((_, index)) =>
				resources.remove(index)
				Right(true)
			case None =>
				Left(ResourceNotFound(id))
		}
	}
}