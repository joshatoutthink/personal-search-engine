package lib

type ModuleInterface interface {
	CollectContent() map[string]string
	IndexContent() map[string]Doc
}
