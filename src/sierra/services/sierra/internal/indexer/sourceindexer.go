package indexer

type sourceIndexer struct {
	IsIndexing bool
}

func newSourceIndexer() sourceIndexer {
	return sourceIndexer{}
}
