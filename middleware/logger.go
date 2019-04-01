package middleware

import (
	"context"
	"log"

	eh "github.com/looplab/eventhorizon"
)

// LoggingMiddleware will log any commands
func LoggingMiddleware(h eh.CommandHandler) eh.CommandHandler {
	return eh.CommandHandlerFunc(func(ctx context.Context, cmd eh.Command) error {
		log.Printf("command: %#v", cmd)
		return h.HandleCommand(ctx, cmd)
	})
}

// Logger is an observer for events
type Logger struct{}

// HandlerType implements the handler type
func (l *Logger) HandlerType() eh.EventHandlerType {
	return "logger"
}

// HandleEvent will log events
func (l *Logger) HandleEvent(ctx context.Context, event eh.Event) error {
	log.Printf("event:  %#v", event)
	return nil
}
