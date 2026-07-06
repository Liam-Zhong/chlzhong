+++
author = "Tuffy"
title = '報告格式——通往知識的路上給你絆一跤'
date = 2024-11-23T22:25:19+08:00
math = true 
draft = false
comments = true
toc = false
description = "pandoc 使用指南"

+++

早在今年 [6 月 3 号](https://chlzhong.org/post/crazy-life_1/)，我就已经被格式這东西折磨過一遍了。如果說大学的知識還能勉强算是人类進步的階梯，那么格式就是這个階梯的絆脚石，放着好好的 Markdown 不用，給一个 95 年的 `.doc` 修修改改（各种格式刷字号间距什么的），收齐給老師之后除了收到是什么都沒得的一个东西。实在是反人类。

幸好，還有 [pandoc](https://github.com/jgm/pandoc/releases/tag/3.5) 。盡管博主只是第一次接触這个东西，但已经控制不住安利給同学們的冲动，因爲你甚至不需要在餅子上搜索学習它的种种用法，只需要准备一个完成的`样例.docx`、你寫好的 `name.md`，放在同一个目錄下，終端敲一句簡單的：

```shell
pandoc name.md -o output.docx --reference-doc=样例.docx
```

OK，去交作业吧  :P
