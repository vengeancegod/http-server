package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"http-server/internal/entities"
	"http-server/internal/infrastructure/beanstalk"
	"http-server/internal/infrastructure/database/sql"
	"http-server/internal/repository/account"
	"http-server/internal/repository/contacts"
	"http-server/internal/repository/unisender_integration"
	"http-server/internal/service"
	cService "http-server/internal/service/contacts"
)

type Worker struct {
	beanstalkService *beanstalk.Service
	contactsService  service.ContactsService
}

func NewWorker(beanstalkService *beanstalk.Service, contactsService service.ContactsService) *Worker {
	return &Worker{
		beanstalkService: beanstalkService,
		contactsService:  contactsService,
	}
}

func (w *Worker) Start() {
	for {
		taskID, taskData, err := w.beanstalkService.GetTask()
		if err != nil {
			log.Println("Error getting task:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		var task entities.ContactsTask
		if err := json.Unmarshal(taskData, &task); err != nil {
			log.Println("Error unmarshalling task:", err)
			continue
		}

		switch task.Action {
		case "add":
			w.handleAddContact(task)
		case "update":
			w.handleUpdateContact(task)
		case "delete":
			w.handleDeleteContact(task)
		default:
			log.Println("Unknown action:", task.Action)
		}

		if err := w.beanstalkService.DeleteTask(taskID); err != nil {
			log.Printf("Error delete task with ID %d: %v", taskID, err)
		} else {
			log.Printf("Task with ID %d successfull delete from queue", taskID)
		}
	}
}

func (w *Worker) handleAddContact(task entities.ContactsTask) {
	contact := entities.Contacts{
		ID:        task.ContactID,
		Name:      task.Name,
		Email:     task.Email,
		AccountID: task.AccountID,
	}

	if err := w.contactsService.CreateContact(contact); err != nil {
		log.Printf("Error creating contact: %v", err)
		return
	}

	log.Printf("Contact with ID %d created", contact.ID)
}

func (w *Worker) handleUpdateContact(task entities.ContactsTask) {
	contact := entities.Contacts{
		ID:    task.ContactID,
		Name:  task.Name,
		Email: task.Email,
	}

	if err := w.contactsService.UpdateContact(contact); err != nil {
		log.Printf("Error update contact: %v", err)
		return
	}

	log.Printf("Contact with ID %d updated", contact.ID)
}

func (w *Worker) handleDeleteContact(task entities.ContactsTask) {
	if err := w.contactsService.DeleteContact(task.ContactID); err != nil {
		log.Printf("Error delete contact: %v", err)
		return
	}

	log.Printf("Contact with ID %d deleted", task.ContactID)
}

func main() {
	workerCountStr := os.Getenv("WORKER_COUNT")
	if workerCountStr == "" {
		workerCountStr = "5"
	}

	numWorkers, err := strconv.Atoi(workerCountStr)
	if err != nil || numWorkers <= 0 {
		log.Fatal("Invalid workers number")
	}

	db, err := sql.InitDB()
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	accountRepo, err := account.NewRepository(db)
	if err != nil {
		log.Fatal("Error creating account repository:", err)
	}

	contactsRepo, err := contacts.NewRepository(db)
	if err != nil {
		log.Fatal("Error creating contacts repository:", err)
	}

	unisenderRepo, err := unisender_integration.NewRepository(db)
	if err != nil {
		log.Fatal("Error creating unisender repository:", err)
	}

	beanstalkService, err := beanstalk.NewService()
	if err != nil {
		log.Fatal("Error initializing beanstalk service:", err)
	}

	contactsService := cService.NewService(contactsRepo, accountRepo, unisenderRepo)

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(beanstalkService, contactsService)
		go worker.Start()
		log.Printf("Start worker %d", i)
	}

	select {}
}
