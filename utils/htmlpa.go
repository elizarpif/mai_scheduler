package utils

import (
	"errors"
	"fmt"
	stack "github.com/golang-collections/collections/stack"
	"strings"
)

type St struct {
	name   string
	values []string
	parent *St
	nodes  []*St
}

func NewSt(name string) *St {
	s := new(St)
	s.name = name
	return s
}
func (s *St) AddNode(v *St) *St {
	s.nodes = append(s.nodes, v)
	if v.parent == nil {
		v.parent = s
	}
	return v
}
func (s *St) AddValue(v string) {
	s.values = append(s.values, v)
}

func (s *St) ToParent() *St {
	return s.parent
}

const delimiter = "--"

func (s *St) Print() {
	s.printNode(delimiter)
}

func (s *St) printNode(nodeTab string) {
	fmt.Println(nodeTab+"name = ", s.name)
	if len(s.values) > 0 {
		fmt.Println(nodeTab+"values: ", s.values)
	}
	for _, node := range s.nodes {
		node.printNode(nodeTab + delimiter)
	}
}

func (s *St) FindSt(name string) *St {
	if s.name == name {
		return s
	}
	for _, node := range s.nodes {
		res := node.FindSt(name)
		if res != nil {
			return res
		}

	}
	return nil
}

func (s *St) FindAllSt(name string) []*St {
	if s.name == name {
		return []*St{s}
	}

	values := []*St{}

	for _, node := range s.nodes {
		res := node.FindAllSt(name)
		values = append(values, res...)
	}
	return values
}

func (s *St) FindAllValues(name string) []string {
	values := []string{}
	if s.name == name { // не проваливаемся
		return s.values
	}
	for _, node := range s.nodes {
		vals := node.FindAllValues(name)
		values = append(values, vals...)
	}
	return values
}

func DeleteComments(html string) (*string, error) {
	for strings.Contains(html, "<!--") {
		ind := strings.Index(html, "<!--")
		last := strings.Index(html, "-->") + 3
		if last < ind {
			return nil, errors.New("error comments parse")
		}
		html = html[:ind] + html[last:]
	}
	return &html, nil
}

func ParseHtml(html string) (*St, error) {
	gethtml, err := DeleteComments(html)
	if err != nil {
		return nil, err
	}
	html = *gethtml

	tags := stack.New()
	values := &St{}

	for strings.Contains(html, "<") {

		ind := strings.Index(html, "<")
		last := strings.Index(html, ">")

		if last < ind {
			fmt.Println(html)
			break
		}
		tag := html[ind+1 : last]
		if ind > 0 {
			slice := html[:ind]
			slice = strings.Trim(slice, "\t")
			slice = strings.Trim(slice, "\n")
			slice = strings.TrimSpace(slice)
			if len(slice) > 0 {
				values.AddValue(slice)
			}
		}

		html = html[last+1:]

		if len(tag) > 0 {
			if strings.Contains(tag, "/") {
				nameTag := tag[strings.Index(tag, "/")+1:]
				if tags.Peek() == nameTag {
					tags.Pop()
					if v := values.ToParent(); v != nil {
						values = v
					}
				}
				//<!--<li class="separator">&nbsp;</li>-->
			} else {
				nameTag := tag
				if strings.Contains(tag, " ") {
					nameTag = tag[:strings.Index(tag, " ")]
				}
				if strings.EqualFold(nameTag, "input") || strings.EqualFold(nameTag, "br") {
					continue
				}
				tags.Push(nameTag)

				if ind := strings.Index(tag, "class="); ind > 0 {
					tag = tag[ind:]
					first := strings.Index(tag, `"`)
					tag = tag[first+1:]
					last := strings.Index(tag, `"`)

					name := tag[:last]
					s := NewSt(name)
					values = values.AddNode(s)
				}
			}
		}

	}

	if tags.Len() > 0 {
		for tags.Len() > 0 {
			fmt.Println(tags.Peek())
			tags.Pop()
		}
		return nil, errors.New("error parsing")
	}
	return values, nil
}
