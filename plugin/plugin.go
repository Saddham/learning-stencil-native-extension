package main

import (
	"fmt"

	"github.com/getoutreach/stencil/pkg/extensions/apiv1"
	"github.com/sirupsen/logrus"
)

// _ is a compile time assertion to ensure we implement
// the Implementation interface
var _ apiv1.Implementation = &GreetingMessagePlugin{}

type GreetingMessagePlugin struct{}

func (tp *GreetingMessagePlugin) GetConfig() (*apiv1.Config, error) {
	return &apiv1.Config{}, nil
}

func (tp *GreetingMessagePlugin) ExecuteTemplateFunction(t *apiv1.TemplateFunctionExec) (interface{}, error) {
	if t.Name == "WelcomeMessage" {
		return "WelcomeMessage", nil
	}

	if t.Name == "GoodbyeMessage" {
		return "GoodbyeMessage", nil
	}

	return nil, fmt.Errorf("No such method: %v", t.Name)
}

func (tp *GreetingMessagePlugin) GetTemplateFunctions() ([]*apiv1.TemplateFunction, error) {
	return []*apiv1.TemplateFunction{
		{
			Name: "WelcomeMessage",
		},
		{
			Name: "GoodbyeMessage",
		},
	}, nil
}

func WelcomeMessage() (interface{}, error) {
	fmt.Println("👋 from the greeting message plugin")
	return "Have a good day, ahead!", nil
}

func GoodbyeMessage() (interface{}, error) {
	fmt.Println("👋 from the greeting message plugin")
	return "See you tomorrow!", nil
}

func main() {
	err := apiv1.NewExtensionImplementation(&GreetingMessagePlugin{}, logrus.New())
	if err != nil {
		logrus.WithError(err).Fatal("failed to start extension")
	}
}
