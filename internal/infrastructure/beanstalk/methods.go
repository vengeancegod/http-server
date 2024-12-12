package beanstalk

import (
	"log"

	"github.com/beanstalkd/go-beanstalk"
)

func (s *Service) PutTask(task []byte) error {
	tube := beanstalk.NewTube(s.conn, "sync_contacts")

	_, err := tube.Put(task, 1, 0, 60)
	if err != nil {
		log.Printf("Error add task to queue: %v", err)
		return err
	}

	log.Printf("Task successfully add to queue: %s", string(task))
	return nil
}
