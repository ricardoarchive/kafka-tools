// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import "github.com/stretchr/testify/mock"
import "github.com/Shopify/sarama"

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Consumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsumePartition provides a mock function with given fields: topic, partition, offset
func (_m *Consumer) ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	ret := _m.Called(topic, partition, offset)

	var r0 sarama.PartitionConsumer
	if rf, ok := ret.Get(0).(func(string, int32, int64) sarama.PartitionConsumer); ok {
		r0 = rf(topic, partition, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sarama.PartitionConsumer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int32, int64) error); ok {
		r1 = rf(topic, partition, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HighWaterMarks provides a mock function with given fields:
func (_m *Consumer) HighWaterMarks() map[string]map[int32]int64 {
	ret := _m.Called()

	var r0 map[string]map[int32]int64
	if rf, ok := ret.Get(0).(func() map[string]map[int32]int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[int32]int64)
		}
	}

	return r0
}

// Partitions provides a mock function with given fields: topic
func (_m *Consumer) Partitions(topic string) ([]int32, error) {
	ret := _m.Called(topic)

	var r0 []int32
	if rf, ok := ret.Get(0).(func(string) []int32); ok {
		r0 = rf(topic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(topic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Topics provides a mock function with given fields:
func (_m *Consumer) Topics() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}