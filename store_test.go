package main

type MockStore struct{}

func (m *MockStore) CreateUser(u *User) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) GetUserById(id string) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) GetUserByEmail(email string) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}
