+++
author = "Tuffy"
title = '报告格式——通往知识的路上给你绊一跤'
date = 2024-11-23T22:25:19+08:00
math = true 
draft = false
comments = true
toc = false
description = "pandoc 使用指南"

+++

早在今年 [6 月 3 号](https://chlzhong.org/post/crazy-life_1/)，我就已经被格式这东西折磨过一遍了。如果说大学的知识还能勉强算是人类进步的阶梯，那么格式就是这个阶梯的绊脚石，放着好好的 Markdown 不用，给一个 95 年的 `.doc` 修修改改（各种格式刷字号间距什么的），收齐给老师之后除了收到是什么都没得的一个东西。实在是反人类。

幸好，还有 [pandoc](https://github.com/jgm/pandoc/releases/tag/3.5) 。尽管博主只是第一次接触这个东西，但已经控制不住安利给同学们的冲动，因为你甚至不需要在饼子上搜索学习它的种种用法，只需要准备一个完成的`样例.docx`、你写好的 `name.md`，放在同一个目录下，终端敲一句简单的：

```shell
pandoc name.md -o output.docx --reference-doc=样例.docx
```

OK，去交作业吧  :P
