package main

import (
	"cli/storage"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/sirupsen/logrus"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func commandHandler(s storage.PassageStorage) func(string) {
	return func(input string) {
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "get":
			e, err := s.Get(cmds[1])
			if err != nil {
				logrus.Errorln(err)
				return
			}
			logrus.Infof("%s: %s - %s", e.Name, e.Username, e.Password)
		case "set":
			err := s.Set(&storage.PassageEntry{
				Name:     cmds[1],
				Username: cmds[2],
				Password: cmds[3],
			})
			if err != nil {
				logrus.Errorln(err)
			}
		}
	}
}

func main() {
	s, err := storage.NewPassageBBoltStorage("", "")
	if err != nil {
		logrus.Fatalln(err)
	}
	defer s.Close()
	p := prompt.New(commandHandler(s), completer)

	p.Run()
}
