package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
            // Method under test
            func (a Awesomo) SuggestAFilmIdea(store FilmIdeaStore) string {
              baseIdea := store.RandomIdea()
              return a.adamSandlerify(baseIdea)
            }

            func (a Awesomo) adamSandlerify(idea string) string {
              return strings.ReplaceAll(idea, "Someone", "Adam Sandler")
            }

            // Real usage
            a := Awesomo{}
            store := RealStore{}
            idea := a.SuggestAFilmIdea(store)
            ...

            // Test
            func TestCountdown(t *testing.T) {
              a := Awesomo{}
              store := FakeStore{}
              got := a.SuggestAFilmIdea(store)
              ...
            }
