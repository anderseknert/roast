package encoding

import (
	"testing"
)

func TestRuleGeneratedBody(t *testing.T) {
	t.Parallel()

	policy := `package test

	import rego.v1

	foo contains "bar"
	`

	module := MustParseModule(policy)

	result, err := JSON().MarshalToString(module)
	if err != nil {
		t.Fatal(err)
	}

	//nolint:lll
	expect := `{"package":{"location":"1:1:cGFja2FnZQ==","path":[{"location":"1:9:dGVzdA==","type":"var","value":"data"},{"location":"1:9:dGVzdA==","type":"string","value":"test"}]},"imports":[{"location":"3:2:aW1wb3J0","path":{"location":"3:9:cmVnby52MQ==","type":"ref","value":[{"location":"3:9:cmVnbw==","type":"var","value":"rego"},{"location":"3:14:djE=","type":"string","value":"v1"}]}}],"rules":[{"location":"5:2:Zm9v","head":{"location":"5:2:Zm9vIGNvbnRhaW5zICJiYXIi","name":"foo","ref":[{"location":"5:2:Zm9v","type":"var","value":"foo"}],"key":{"location":"5:15:ImJhciI=","type":"string","value":"bar"}}}]}`

	if result != expect {
		t.Errorf("Expected %s but got %s", expect, result)
	}
}

func TestRuleElseGeneratedBody(t *testing.T) {
	t.Parallel()

	policy := `package test

	import rego.v1

	foo := "bar" if {
		input.baz
	} else := false
	`

	module := MustParseModule(policy)

	result, err := JSON().MarshalToString(module)
	if err != nil {
		t.Fatal(err)
	}

	//nolint:lll
	expect := `{"package":{"location":"1:1:cGFja2FnZQ==","path":[{"location":"1:9:dGVzdA==","type":"var","value":"data"},{"location":"1:9:dGVzdA==","type":"string","value":"test"}]},"imports":[{"location":"3:2:aW1wb3J0","path":{"location":"3:9:cmVnby52MQ==","type":"ref","value":[{"location":"3:9:cmVnbw==","type":"var","value":"rego"},{"location":"3:14:djE=","type":"string","value":"v1"}]}}],"rules":[{"location":"5:2:Zm9vIDo9ICJiYXIiIGlmIHsKCQlpbnB1dC5iYXoKCX0gZWxzZSA6PSBmYWxzZQ==","head":{"location":"5:2:Zm9vIDo9ICJiYXIi","name":"foo","ref":[{"location":"5:2:Zm9v","type":"var","value":"foo"}],"assign":true,"value":{"location":"5:9:ImJhciI=","type":"string","value":"bar"}},"body":[{"location":"6:3:aW5wdXQuYmF6","terms":{"location":"6:3:aW5wdXQuYmF6","type":"ref","value":[{"location":"6:3:aW5wdXQ=","type":"var","value":"input"},{"location":"6:9:YmF6","type":"string","value":"baz"}]}}],"else":{"location":"7:4:ZWxzZSA6PSBmYWxzZQ==","head":{"location":"7:4:ZWxzZSA6PSBmYWxzZQ==","name":"foo","ref":[{"location":"5:2:Zm9v","type":"var","value":"foo"}],"assign":true,"value":{"location":"7:12:ZmFsc2U=","type":"boolean","value":false}}}}]}`

	if result != expect {
		t.Errorf("Expected %s but got %s", expect, result)
	}
}
