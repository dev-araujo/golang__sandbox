package task

import "fmt"

type Service interface {
	GetListTasks() []Task
	AddTask(description string) Task
	DeleteTask(id uint) error
	UpdateTask(id uint, description string, completed bool) (Task, error)
	GetTask(id uint) (Task, error)
}

type service struct {
	tasks  []Task
	nextID uint
}

func NewService() Service {
	return &service{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (s *service) GetListTasks() []Task {
	return s.tasks
}

func (s *service) AddTask(description string) Task {
	newTask := Task{
		ID:          s.nextID,
		Description: description,
		Completed:   false,
	}

	s.tasks = append(s.tasks, newTask)
	s.nextID++

	return newTask
}

func (s *service) DeleteTask(id uint) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Tarefa não encontrada")
}

func (s *service) UpdateTask(id uint, description string, completed bool) (Task, error) {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Description = description
			s.tasks[i].Completed = completed
			return s.GetTask(id)
		}
	}
	return Task{}, fmt.Errorf("Tarefa não encontrada")
}

func (s *service) GetTask(id uint) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("Tarefa não encontrada")
}
