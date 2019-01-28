package event

import "go-starter-commandhandler/src/domain/model"

//DomainEvent : Event to define a changing of business object state
type DomainEvent struct {
	DomainModel             model.DomainModel
	PublicationDateTimeUnix string
}
