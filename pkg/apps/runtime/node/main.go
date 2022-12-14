package node

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Node struct {
	Name      string
	Port      int64
	Directory string
	Language  string
	Framework string
}

type NodeSettings struct {
	Version string
}

var settings NodeSettings = NodeSettings{
	Version: "18",
}

func (n *Node) CreateNodeApp() error {
	// all node apps should have a src directory
	err := os.MkdirAll(n.Directory+"/src", 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// all node apps will be on node 18

	// framework
	if err := n.createFramework(); err != nil {
		return err
	}

	//
	return nil
}

func (n *Node) npmInstall() error {
	cmd := exec.Command("npm", "install")
	cmd.Dir = n.Directory
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	n.Log(fmt.Sprintf("%s\n", out))
	return nil
}

func (n *Node) createFramework() error {
	switch strings.ToLower(n.Framework) {
	case "express":
		if err := n.setupExpress(); err != nil {
			return err
		}
	case "koa":
		if err := n.setupKoa(); err != nil {
			return err
		}
	case "react":
		if err := n.setupReact(); err != nil {
			return err
		}
	}
	return nil
}