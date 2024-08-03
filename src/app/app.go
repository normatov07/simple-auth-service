package app

type ApplicationContext struct {
}

func GetApplicationContext() *ApplicationContext {
	return &ApplicationContext{}
}
