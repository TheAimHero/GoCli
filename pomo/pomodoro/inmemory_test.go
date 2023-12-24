// +build inmemory

package pomodoro_test

import (
  "testing"

  "pomo/pomodoro"
  "pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
  t.Helper()

  return repository.NewInMemoryRepo(), func() {}
}
