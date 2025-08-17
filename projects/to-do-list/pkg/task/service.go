package task

type Service interface {
	GetListTasks() []Task
	AddTask(description string) Task
	DeleteTask(id uint)
	UpdateTask(id uint, description string, completed bool) Task
	GetTask(id uint) Task
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

func (s *service) DeleteTask(id uint) {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return
		}
	}
}

func (s *service) UpdateTask(id uint, description string, completed bool) Task {

	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Description = description
			s.tasks[i].Completed = completed
			return s.GetTask(id)
		}
	}
	return s.GetTask(id)
}

func (s *service) GetTask(id uint) Task {
	for _, task := range s.tasks {
		if task.ID == id {
			return task
		}
	}
	return Task{}
}