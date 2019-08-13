package main

import "testing"

func TestURLDecode(t *testing.T) {

	urls := map[string]string{
		"?user=Fran%C3%A7ois":        "?user=François",
		"http://?user=Fran%C3%A7ois": "http://?user=François",
		"https%3A%2F%2Fwww.google.com%2Fsearch%3Fclient%3Dubuntu%26channel%3Dfs%26q%3Durls+to+decode+examples%26ie%3Dutf-8%26oe%3Dutf-8": "https://www.google.com/search?client=ubuntu&channel=fs&q=urls+to+decode+examples&ie=utf-8&oe=utf-8",
	}

	for encoded, expected := range urls {
		if decoded, _ := decode(encoded); decoded != expected {
			t.Fatalf("Got: %s, Expected: %s", decoded, expected)
		}
	}
}

func TestURLEncode(t *testing.T) {

	urls := map[string]string{
		"?user=François":        "%3Fuser%3DFran%C3%A7ois",
		"http://?user=François": "http%3A%2F%2F%3Fuser%3DFran%C3%A7ois",
		"https://www.google.com/search?client=ubuntu&channel=fs&q=urls+to+decode+examples&ie=utf-8&oe=utf-8": "https%3A%2F%2Fwww.google.com%2Fsearch%3Fclient%3Dubuntu%26channel%3Dfs%26q%3Durls%2Bto%2Bdecode%2Bexamples%26ie%3Dutf-8%26oe%3Dutf-8",
	}

	for decoded, expected := range urls {
		if encoded := encode(decoded); encoded != expected {
			t.Fatalf("Got: %s, Expected: %s", encoded, expected)
		}
	}
}
