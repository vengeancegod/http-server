package beanstalk

import (
	"errors"
	"http-server/internal/entities"
	"log"

	"github.com/beanstalkd/go-beanstalk"
)

type Service struct {
	conn *beanstalk.Conn
}

type BeanstalkService interface {
	PutTask(task []byte) error
	GetTask() (uint64, []byte, error)
	DeleteTask(taskID uint64) error
}

func NewService() (*Service, error) {
	conn, err := beanstalk.Dial("tcp", "ddev-beanstalkd:11300")
	if err != nil {
		return nil, errors.New(entities.ErrConnectBeanstalk)
	}
	log.Println("Successfull connected to beanstalk")
	return &Service{
		conn: conn,
	}, nil
}
