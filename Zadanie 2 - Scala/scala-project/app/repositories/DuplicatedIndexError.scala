package repositories

case class DuplicatedIndexError[IndexType](index: IndexType) extends RepositoryError {}

