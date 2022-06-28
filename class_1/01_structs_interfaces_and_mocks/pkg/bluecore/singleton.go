package bluecore

type Singleton struct{}

var globalSingleton = Singleton{}

func GetSingleton() *Singleton {
	return &globalSingleton
}
