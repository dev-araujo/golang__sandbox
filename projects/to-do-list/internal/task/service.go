package task

type Service interface {
	GetListTasks() []Task
	AddTask(description string) Task
	DeleteTask(id uint)
	UpdateTask(id uint, description string, completed bool)
	GetTask(id uint) Task
	CheckTask(id uint)
}

type Task struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
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

func (s *service) UpdateTask(id uint, description string, completed bool) {
}

func (s *service) GetTask(id uint) Task {
	for _, task := range s.tasks {
		if task.ID == id {
			return task
		}
	}
	return Task{}
}

func (s *service) CheckTask(id uint) {
}
