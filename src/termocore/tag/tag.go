package tag

type Tag interface {
	//Constructors
	NewTag(key string) Tag
	NewValuedTag(key string, values []string) Tag

	//Properties
	GetKey() string
	GetPlugin() string
	GetValue() []string
	HasValue() bool
	GetImportance() int
	AddValue(value string)
	RemoveValue(value string)
}
