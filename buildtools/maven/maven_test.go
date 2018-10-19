package maven_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fossas/fossa-cli/buildtools/maven"
)

/*
	├── dep:one:1.0.0
	└─┬ dep:two:2.0.0
		├─┬ dep:three:3.0.0
		│ └── dep:four:4.0.0
		└── dep:five:5.0.0
*/
var depOne = maven.Dependency{Name: "dep:one", Version: "1.0.0", Failed: false}
var depTwo = maven.Dependency{Name: "dep:two", Version: "2.0.0", Failed: false}
var depThree = maven.Dependency{Name: "dep:three", Version: "3.0.0", Failed: false}
var depFour = maven.Dependency{Name: "dep:four", Version: "4.0.0", Failed: false}
var depFive = maven.Dependency{Name: "dep:five", Version: "5.0.0", Failed: false}

func TestParseDependencyTree(t *testing.T) {
	dat, err := ioutil.ReadFile("testdata/osx.out")
	assert.NoError(t, err)
	direct, transitive, err := maven.ParseDependencyTree(string(dat))

	assert.NoError(t, err)
	assert.Equal(t, 2, len(direct))
	depList := []maven.Dependency{depOne, depTwo}
	assert.Equal(t, depList, direct)

	expectedGraph := map[maven.Dependency][]maven.Dependency{
		depOne:   nil,
		depTwo:   []maven.Dependency{depThree, depFive},
		depThree: []maven.Dependency{depFour},
	}
	assert.Equal(t, expectedGraph, transitive)
	assert.Equal(t, 3, len(transitive))
}
