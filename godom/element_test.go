package godom

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

var testAssetPath = "../test_asset/test.html"

func getDocument(r *require.Assertions) *Element {
	bytesRead, err := ioutil.ReadFile(testAssetPath)
	r.NoError(err, "Expect no error when reading from", testAssetPath)

	rTest := bytes.NewReader(bytesRead)

	return ParseDocument(rTest)
}

func TestParseDocument(t *testing.T) {
	require := require.New(t)

	bytesRead, err := ioutil.ReadFile(testAssetPath)
	require.NoError(err, "Expected no error when reading from", testAssetPath)

	rTest := bytes.NewReader(bytesRead)
	rShould := bytes.NewReader(bytesRead)

	elTest := ParseDocument(rTest)

	doc, err := html.Parse(rShould)
	var elShould *Element
	if err != nil {
		elShould = &Element{}
	}
	elShould = &Element{
		node: doc,
	}

	require.Equal(elShould, elTest, "Expected the created Document to be equal to", elShould)

	var node *html.Node
	elShouldNot := &Element{
		node: node,
	}

	require.NotEqual(elShouldNot, elTest, "Expected the created Document to be equal to", elShould)
}

func TestAnchors(t *testing.T) {
	require := require.New(t)

	bytesRead, err := ioutil.ReadFile(testAssetPath)
	require.NoError(err, "Expected no error when reading from", testAssetPath)

	rTest := bytes.NewReader(bytesRead)

	docTest := ParseDocument(rTest)
	anchors := docTest.Anchors()

	require.Len(anchors, 548, "Expected the anchors to have length of 548 elements")
}

func TestChildren(t *testing.T) {
	require := require.New(t)

	docTest := getDocument(require)

	require.NotEmpty(docTest.Children, "Expected to have more than one child")
}
