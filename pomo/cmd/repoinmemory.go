// +build inmemory

package cmd

import (
  "pomo/pomodoro"
  "pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
  return repository.NewInMemoryRepo(), nil
}
