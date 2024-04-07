package repositories

case class ResourceNotFound[IndexType](index: IndexType) extends RepositoryError {}

