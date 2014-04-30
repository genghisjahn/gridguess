package main

type GuessResult struct {
	Message    string
	Found      bool
	GuessCount int
}

func (gr GuessResult) String() string {
	return gr.Message
}
