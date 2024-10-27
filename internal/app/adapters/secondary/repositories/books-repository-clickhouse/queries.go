package books_repository_clickhouse

func (repo *BooksRepositoryClickhouse) getGetSomethingQuery(shkIDs []int64) string {
	query := `SELECT * FROM table`

	return query
}
