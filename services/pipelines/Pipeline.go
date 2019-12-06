package pipelines

// Pipeline интерфейс для любого папйлайна.
type Pipeline interface {
	// Стартует пийплайн.
	InitPipeline() (bool, error)
}
