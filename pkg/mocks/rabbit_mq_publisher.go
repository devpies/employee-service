// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import rabbitmq "github.com/wagslane/go-rabbitmq"

// RabbitMQPublisher is an autogenerated mock type for the rabbitMQPublisher type
type RabbitMQPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: message, routingKey, options
func (_m *RabbitMQPublisher) Publish(message []byte, routingKey []string, options ...func(*rabbitmq.PublishOptions)) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message, routingKey)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}