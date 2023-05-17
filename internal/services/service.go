package services

import (
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

var Container di.Container

func SetupServices(services ...*di.Def) {
	builder, _ := di.NewBuilder()

	for _, service := range services {
		err := builder.Add(*service)
		if err != nil {
			log.Fatal("Failed to register service: ", service.Name, "")
		}
	}
	Container = builder.Build()
	for _, service := range services {
		_, err := Container.SafeGet(service.Name)
		if err != nil {
			log.Fatalf("Error Loading Service %v Failed, Err: %v: ",
				service.Name, err.Error())
		}
	}

	log.Infof("Loaded %v services", len(Container.Definitions()))
}

func GetService[T interface{}](name string) T {
	service, err := GetServiceSafe[T](name)
	if err != nil {
		log.Fatal("Failed to get service: ", name, "")
		return service
	}
	return service
}

func GetServiceSafe[T interface{}](name string) (T, error) {
	service, err := Container.SafeGet(name)
	return service.(T), err
}

func HasService(name string) bool {
	service, err := Container.SafeGet(name)
	return err == nil && service != nil
}
