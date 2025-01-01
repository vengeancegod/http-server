package beanstalk

import (
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func (s *Service) PutTask(task []byte) error {
	tube := beanstalk.NewTube(s.conn, "job_contacts")

	_, err := tube.Put(task, 1, 0, 60)
	if err != nil {
		log.Printf("Error add task to queue: %v", err)
		return err
	}

	log.Printf("Task successfull add to queue: %s", string(task))
	return nil
}

func (s *Service) GetTask() (uint64, []byte, error) {
	tubeSet := beanstalk.NewTubeSet(s.conn, "job_contacts")

	log.Printf("Trying to reserve task for %d seconds", 60)
	id, body, err := tubeSet.Reserve(60 * time.Second)
	if err != nil {
		log.Printf("Error reserving task: %v", err)
		return 0, nil, err
	}

	log.Printf("Task successfull reserved: ID=%d, Body=%s", id, string(body))

	return id, body, nil
}

func (s *Service) DeleteTask(taskID uint64) error {
	err := s.conn.Delete(taskID)
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", taskID, err)
		return err
	}

	log.Printf("Task with ID %d successfull delete from queue", taskID)
	return nil
}
