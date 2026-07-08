package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// 极其强悍的正则：匹配连续的西文、数字，并允许中间带有空格和特殊符号（如 C++, iPhone 17 Pro Max）
var latinRegex = regexp.MustCompile(`[a-zA-Z0-9\.\+\-]+(?:\s+[a-zA-Z0-9\.\+\-]+)*`)

var skipTags = map[string]bool{
	"code":     true,
	"pre":      true,
	"script":   true,
	"style":    true,
	"svg":      true,
	"math":     true,
	"kbd":      true,
	"title":    true, // 🚨 核心修复：网页标题绝对不能包裹 span
	"textarea": true, // 补充：输入框里的默认文本也不能碰
	"option":   true, // 补充：下拉菜单里的文本也不能碰
}

func main() {
	// 设定扫描目标为 public 目录（Hugo 编译后的静态文件输出目录）
	targetDir := "../../public"

	fmt.Println("🚀 正在启动 Typography Post Processor...")

	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只处理 .html 文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			processHTMLFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ 处理失败:", err)
	} else {
		fmt.Println("✨ 全站 HTML 中西文排版处理完毕！")
	}
}

func processHTMLFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return
	}

	// 开启 DFS 深度优先遍历
	traverse(doc, false)

	// 将处理后的 DOM 树重新渲染并覆写原文件
	var buf bytes.Buffer
	html.Render(&buf, doc)
	os.WriteFile(path, buf.Bytes(), 0644)
}

func traverse(n *html.Node, skip bool) {
	// 如果遇到黑名单标签，将其及其子节点标记为 skip
	if n.Type == html.ElementNode && skipTags[n.Data] {
		skip = true
	}

	// 发现普通文本节点，且当前处于安全可操作状态
	if n.Type == html.TextNode && !skip {
		// 检查是否包含西文字符或数字串
		if latinRegex.MatchString(n.Data) {
			wrapLatinNodes(n)
		}
	}

	// 继续递归遍历。必须预存 NextSibling，因为我们在处理中可能会动态插入/删除同级节点
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		traverse(c, skip)
		c = next
	}
}

// wrapLatinNodes 负责将 TextNode 劈开，抽离出西文片段并包裹 <span class="latin">
func wrapLatinNodes(n *html.Node) {
	text := n.Data
	matches := latinRegex.FindAllStringIndex(text, -1)
	if len(matches) == 0 {
		return
	}

	parent := n.Parent
	if parent == nil {
		return
	}

	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]

		// 1. 插入英文前面剩下的中文
		if start > lastIndex {
			prevText := text[lastIndex:start]
			parent.InsertBefore(&html.Node{Type: html.TextNode, Data: prevText}, n)
		}

		// 2. 将英文片段包裹进 <span class="latin">
		latinStr := text[start:end]
		span := &html.Node{
			Type: html.ElementNode,
			Data: "span",
			Attr: []html.Attribute{{Key: "class", Val: "latin"}},
		}
		span.AppendChild(&html.Node{Type: html.TextNode, Data: latinStr})
		parent.InsertBefore(span, n)

		lastIndex = end
	}

	// 3. 插入最后一段尾巴（如果有）
	if lastIndex < len(text) {
		nextText := text[lastIndex:]
		parent.InsertBefore(&html.Node{Type: html.TextNode, Data: nextText}, n)
	}

	// 4. 过河拆桥：移除原本那个未经处理的超大 TextNode
	parent.RemoveChild(n)
}