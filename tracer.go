package gutils

import (
	"github.com/opentracing/opentracing-go"
	"github.com/streadway/amqp"
)

// StartRootSpan : Start root span with request URL
func StartRootSpan(tracer opentracing.Tracer, requestURL string) opentracing.Span {
	return tracer.StartSpan(requestURL)
}

// StartSpanFromSpanContext : Start span from span context
func StartSpanFromSpanContext(tracer opentracing.Tracer, operationName string, spanContext opentracing.SpanContext) opentracing.Span {
	return tracer.StartSpan(operationName, opentracing.FollowsFrom(spanContext))
}

// Inject : Inject span context
func Inject(tracer opentracing.Tracer, spanContext opentracing.SpanContext, header amqp.Table) {
	carrier := amqpHeadersCarrier(header)
	tracer.Inject(spanContext, opentracing.TextMap, carrier)
}

// StartSpanFromAMQPHeader : Start span from amqp header
func StartSpanFromAMQPHeader(tracer opentracing.Tracer, operationName string, header amqp.Table) opentracing.Span {
	carrier := amqpHeadersCarrier(header)
	spanContext, _ := tracer.Extract(opentracing.TextMap, carrier)

	return tracer.StartSpan(operationName, opentracing.FollowsFrom(spanContext))
}

type amqpHeadersCarrier map[string]interface{}

// ForeachKey conforms to the TextMapReader interface.
func (c amqpHeadersCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, val := range c {
		v, ok := val.(string)
		if !ok {
			continue
		}
		if err := handler(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Set implements Set() of opentracing.TextMapWriter.
func (c amqpHeadersCarrier) Set(key, val string) {
	c[key] = val
}
