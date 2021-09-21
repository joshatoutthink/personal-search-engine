package lib

import (
	"testing"
)

func TestTest(t *testing.T) {
	if 1 == 2 {
		t.Fatalf(`Hi() failed`)
	}
}

func TestWrite(t *testing.T) {
	content := "Content"
	Write("./test", &content)
	if 1 == 2 {
		t.Fatalf("err")
	}
}

func TestMdTitle(t *testing.T) {
	input := "title: Hi Im a dog"
	t.Log("hi")
	if MdTitle.MatchString(input) != true {
		t.Fatalf("MdTitle is a bust duude")
	}

	input2 := "not a title:"
	if MdTitle.MatchString(input2) == true {
		t.Fatalf("MdTitle says everything is a title	")
	}
}

//TODO: Write more test cases
func TestTokenize(t *testing.T) {
	input := "title:Hi Im a dog poop\n"
	tokens := Tokenize(input)

	for token, count := range tokens {
		t.Log(token, count)
	}
	if tokens["title"] == float64(1) {
		t.Fatalf("Tokenizer did not tokenize the word \"title\"")
	}
}
