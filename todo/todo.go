package todo

import (
	"fmt"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"time"
)

type item struct {
	Task        string	`json:task`
	Done        bool	`json: done`
	CreatedAt   time.Time	`json:create_at`
	CompletedAt time.Time	`json:completed_at`
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item {
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index >= len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index >= len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index], ls[index+1:]...)
	return nil
}

func (t *Todos) Load(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(fileName string) error {
	file, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, file, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}