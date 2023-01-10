package main

import (
	"github.com/razeli/go-vue-task-tracker/controller"
	"github.com/razeli/go-vue-task-tracker/model"
	
)

func main() {
	model.Init()
	controller.Start()
}
