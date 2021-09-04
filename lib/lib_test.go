package lib

import "testing"

func TestTest(t *testing.T){
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
 
func testMdTitle(t *testing.T){
	input:= "title: Hi Im a dog"
	if MdTitle.MatchString(input) != true {
		t.Fatalf("MdTitle is a bust duude")
	}

	input2 := "not a title:"
	if MdTitle.MatchString(input2) == true {
		t.Fatalf("MdTitle says everything is a title	")
	}
}

