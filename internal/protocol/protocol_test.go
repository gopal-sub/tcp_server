package protocol

import (
	"errors"
	"testing"
)


type ParserTestCase struct{
	Name string
	Input string
	Output Command
	err error
	
}



func TestParser(t *testing.T){
	testCase := []ParserTestCase{
		{Name: "Valid Join", Input: "JOIN roomx", Output: Command{Type: JOIN, Arg: "roomx"}, err: nil},
		{Name: "Valid Message", Input: "MESSAGE hi there", Output: Command{Type: MESSAGE, Arg: "hi there"}, err: nil},
		{Name: "Valid Quit", Input: "QUIT", Output: Command{Type: QUIT, Arg: ""},  err: nil},
		{Name: "Invalid commmand", Input: "Message hi there", Output: Command{},  err: InvalidCommand},
		{Name: "Join room 2 args", Input: "JOIN room1 room2", Output: Command{},  err: InvalidCommand},
		{Name: "Quit with arg", Input: "QUIT room1", Output: Command{},  err: InvalidCommand},
	}

	for _ , test := range testCase{
		t.Run(test.Name, func(t *testing.T) {
			cmd, err := Parser(test.Input)
			if err != nil {
				if !errors.Is(err, test.err) {
					t.Fatalf("expected error %v, got %v", test.err, err)
				}
				return
			}

			if test.err != nil {
				t.Fatalf("expected error for this test case got nil")
			}

			if cmd == nil {
				t.Fatalf("got cmd nil and err nil")
			}
			
			// cmd can be nil so below can deref null

			if *cmd != test.Output {
				t.Fatalf("Test could not pass for testName=:%v, wrong output %v", test.Name, cmd)
			}

		})
	}
}