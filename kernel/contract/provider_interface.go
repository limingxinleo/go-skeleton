package contract

type ProviderInterface interface {
	Invoke() DefinitionHandler
}
