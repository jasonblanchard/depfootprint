package npmjs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type DistTags struct {
	Latest string `json:"latest"`
}

type Version struct {
	Dependencies map[string]string `json:"dependencies"`
}

type Pkg struct {
	DistTags DistTags           `json:"dist-tags"`
	Versions map[string]Version `json:"versions"`
}

type NpmJS struct {
}

func (n *NpmJS) Get(pkgName string) (*Pkg, error) {
	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://registry.npmjs.org/%s", pkgName), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error, non-200 status code for %s%s: %v", req.URL.Host, req.URL.Path, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	pkg := &Pkg{}

	err = json.Unmarshal(body, pkg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return pkg, nil
}

func GetDependencyListByVersion(pkg *Pkg, version *string) (dependencies map[string]string, ok bool) {
	selectedVersion := pkg.DistTags.Latest
	if version != nil {
		// TODO: Find by matching semver
		selectedVersion = *version
	}
	packageVersion, ok := pkg.Versions[selectedVersion]
	if !ok {
		return map[string]string{}, ok
	}

	return packageVersion.Dependencies, true
}

type DepNode struct {
	Name     string
	Version  string
	Children []*DepNode
}

func (n *NpmJS) WalkDependenciesSync(rootNode *DepNode, level uint) (*DepNode, error) {
	fmt.Printf("Getting dep %s at level %v\n", rootNode.Name, level)
	root, err := n.Get(rootNode.Name)

	if err != nil {
		return rootNode, fmt.Errorf("error getting package %v: %w", rootNode.Name, err)
	}

	deps, ok := GetDependencyListByVersion(root, nil)

	if !ok {
		return rootNode, nil
	}

	if len(deps) == 0 {
		return rootNode, nil
	}

	for pkg := range deps {
		intermediateNode := &DepNode{
			Name:     pkg,
			Version:  root.DistTags.Latest,
			Children: make([]*DepNode, 0),
		}

		childNode, err := n.WalkDependenciesSync(intermediateNode, level+1)
		if err != nil {
			return rootNode, fmt.Errorf("error walking dep: %w", err)
		}
		rootNode.Children = append(rootNode.Children, childNode)
	}

	return rootNode, nil
}

func (n *NpmJS) WalkDependenciesAsync(rootNode *DepNode, level uint) (*DepNode, error) {
	root, err := n.Get(rootNode.Name)

	if err != nil {
		return rootNode, fmt.Errorf("error getting package: %w", err)
	}

	rootNode.Version = root.DistTags.Latest

	deps, ok := GetDependencyListByVersion(root, nil)
	if !ok {
		return rootNode, nil
	}

	if len(deps) == 0 {
		return rootNode, nil
	}

	depListStream := make(chan string)

	// Push onto deps stream
	go func() {
		for pkg := range deps {
			depListStream <- pkg
		}
		close(depListStream)
	}()

	resultStreams := []chan *DepNode{}
	errStreams := []chan error{}

	// make workers to work on pkgs
	for i := 0; i < 8; i++ {
		resultStream := make(chan *DepNode)
		resultStreams = append(resultStreams, resultStream)
		errStream := make(chan error)
		errStreams = append(errStreams, errStream)
		// j := i
		go func() {
			for depName := range depListStream {
				intermediateNode := &DepNode{
					Name:     depName,
					Children: make([]*DepNode, 0),
				}

				childNode, err := n.WalkDependenciesAsync(intermediateNode, level+1)
				if err != nil {
					errStream <- err
				}

				resultStream <- childNode
				// fmt.Printf("handled %v in worker %v at level %v\n", depName, j, level)
			}
			close(resultStream)
			close(errStream)
		}()
	}

	// merge result of workers back and append onto rootNode.Children
	var wg sync.WaitGroup
	resultDeps := make(chan *DepNode)
	mergeResults := func(c <-chan *DepNode) {
		defer wg.Done()
		for i := range c {
			resultDeps <- i
		}
	}

	for _, c := range resultStreams {
		wg.Add(1)
		go mergeResults(c)
	}

	resultErrs := make(chan error)
	mergeResultErrs := func(c <-chan error) {
		defer wg.Done()
		for i := range c {
			resultErrs <- i
		}
	}

	for _, ec := range errStreams {
		wg.Add(1)
		go mergeResultErrs(ec)
	}

	go func() {
		for d := range resultDeps {
			rootNode.Children = append(rootNode.Children, d)
		}
	}()

	finalErrs := []error{}

	go func() {
		for e := range resultErrs {
			finalErrs = append(finalErrs, e)
		}
	}()

	wg.Wait()

	var finalErr error

	if len(finalErrs) > 0 {
		finalErr = finalErrs[0]
	}

	return rootNode, finalErr
}

func (n *NpmJS) Tree(pkg string) (*DepNode, error) {
	node := &DepNode{
		Name:     "express",
		Children: make([]*DepNode, 0),
	}

	deps, err := n.WalkDependenciesAsync(node, 0)
	if err != nil {
		return deps, fmt.Errorf("error walking dependencies: %w", err)
	}

	return deps, nil
}
