package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	// pattern: /programs/golang
	// return: parts: [programs, golang]
	partList := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range partList {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' { // ?
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method] // The key of roots is method name. Such as GET or POST. The value is a trie.
	if !ok {
		r.roots[method] = &node{} // eg: If the "GET" method(key) not exists, making a trie root node for "GET" key.
	}
	r.roots[method].insert(pattern, parts, 0) // If the key(method) exists, the value(pattern) will be inserted.
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) // path == pattern
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {  // If a path includes a param, the param will include ":" or "*".
				params[part[1:]] = searchParts[index]  // Real param needs to remove ":".
			}
			if part[0] == '*' && len(part) > 1 {  // "*" express arbitrary number of params.
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}

	c.Params = params
	key := c.Method + "-" + n.pattern
	r.handlers[key](c)
}
