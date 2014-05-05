package main

type GuessResult struct {
	Message    string
	Found      bool
	GuessCount int
	Value      int //This should return -1 if target is lower, 1 if target is higher and 0 if correct.
}

func (gr GuessResult) String() string {
	return gr.Message
}
