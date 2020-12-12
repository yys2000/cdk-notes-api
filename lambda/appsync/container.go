package main

import (
	"fmt"
	"github.com/geoffLondon/aws-appsync-resolvers"
	service_resolver "github.com/geoffLondon/cdk-notes-api/resolver/service"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Resolver() resolvers.Repository
}

type DefaultContainer struct {
	CreateNoteResolver service_resolver.CreateNoteResolver
}

func (container DefaultContainer) Resolver() resolvers.Repository {
	repository := resolvers.New()

	/*	if err := repository.Add("mutation.createNote", container.CreateNoteResolver.Handle); err != nil {
			log.WithField("err", err).Warn("error adding resolver to repository")
		}
	*/
	resolversMap := map[string]interface{}{
		"mutation.createNote": container.CreateNoteResolver.Handle,
	}

	for resolver, handler := range resolversMap {
		if err := repository.Add(resolver, handler); err != nil {
			log.WithField("err", err).Warn("error adding resolver to repository")
		}
		fmt.Println("******* resolver *******", resolver)
		fmt.Println("******* handler *******", handler)
		log.WithField("resolver", resolver).Warn("resolver added to repository")
	}

	return repository
}
