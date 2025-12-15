+++
author = "Tuffy"
title = '建站指北'
date = 2024-05-26T11:20:57+08:00
comments = true
toc = true
description = "{{ .Summary }}"
+++

#### 原文见[如何用 GitHub Pages + Hugo 搭建个人博客 · 小绵尾巴 (cuttontail.blog)](https://cuttontail.blog/blog/create-a-wesite-using-github-pages-and-hugo/)

#### 整个过程基于 Windows11

----

## 0.部落格 

### 部落格是什么？
>部落格（博客）是一种在线日记型式的个人网站，借由张帖子章、图片或影片来记录生活、抒发情感或分享信息。博客上的文章通常根据张贴时间，以倒序方式由新到旧排列。许多博客作者专注评论特定的课题或新闻，其他则作为个人日记。一个典型的博客结合了文字、图像、其他博客或网站的超链接、及其它与主题相关的媒体。能够让读者以互动的方式留下意见，是许多博客的重要要素。大部分的博客内容以文字为主，也有一些博客专注艺术、摄影、视频、音乐、播客等各种主题。博客是社会媒体网络的一部分。——[维基百科](https://zh.wikipedia.org/wiki/%E7%B6%B2%E8%AA%8C)

一直认为一件事情如果你忘记了那么它就不存在。“记录”从古至今都是很有仪（niu）式（ben）感的事情：登高作赋，临泉赋诗，兰亭修序……甚至还要画张山水裱起来，聊以服人。如今人们爱用朋友圈（个人很鄙夷路上奔波之后的所谓“又是自由的一天”）记录生活也差不多是这个道理，只是依赖于微信总让人不舒服，因此自己搭一个部落格是一件很棒很酷的事情，否则怎么算会“上网”呢？本文便是出于这样的目的，争取让每个人都可以零基础搭建自己的部落格。



这篇教程假设你已经：

1. 安装了 [Git](https://git-scm.com/)，并且了解基本的 Git 知识；
2. 有一个 [GitHub](https://github.com/) 账号；
3. 有自己偏好的文本编辑器（比如 [Typora](https://typora.io/)）。

如果你满足以上条件**那就开始吧！**

------

## 1. 安装 Hugo

1. 点击 [Release v0.126.1 · gohugoio/hugo (github.com)]((https://github.com/gohugoio/hugo/releases/tag/v0.126.1)) 并下滑找到 hugo_0.126.1_windows-amd64.zip，点击下载。（本文写于 2024 年 5 月 26 号，你下载的时候版本可能已经更新。）
   ![版本选择](https://pic3.zhimg.com/80/v2-e0d9fd92a6ed78579e4dbaf223d20602_1440w.webp)

2. 把下载好的 .zip 文件解压，并放置在你想放的文件夹中（我的路径：D:\DevTools\Hugo，也就是说现在 Hugo 目录下有一个名为 hugo_0.126.1_windows-amd64 的文件夹）。打开设置，搜索“编辑系统环境变量”，在弹出的“系统属性”窗口中点击“环境变量”，再在弹出的“环境变量”窗口的“系统变量”中选中 Path 路径，并点击编辑。通过右侧的“新建”按钮，将 `D:\DevTools\Hugo\hugo_0.126.1_windows-amd64` 添加到环境变量（根据自己的路径修改）。

2. WIN+D 回到桌面，右键选择“在终端中打开”。输入：

   ```shell
   hugo version
   ```
   一切正常的话会显示你目前安装的版本号：
   
   ![hugo version](https://pic3.zhimg.com/80/v2-0da4527059ac8dee7484e9ee28be7446_1440w.webp)

------

## 2. 创建 GitHub 仓库

### 2.1 创建博客源仓库

1. 命名**博客源仓库**（我使用的是 `demo`）;
2.  勾选 **Public**，设置为公开仓库；
3.  勾选添加 **README** 文件。

![创建demo仓库](https://pic3.zhimg.com/80/v2-c718ed26fa965c2b5f9dcf16a59407a6_1440w.webp)

### 2.2 创建 GitHub Page 仓库

1. 命名 **GitHub Pages** 仓库，这个仓库必须使用特殊的命名格式 `<username.github.io>` ，`<username>` 是你自己的 GitHub 的用户名；
2.  勾选 **Public**，设置为公开仓库；
3.  勾选添加 **README** 文件，这会设置 `main` 分支为仓库的默认主分支，在后面提交推送博客内容时很重要。

![创建name.github.io仓库](https://pic3.zhimg.com/80/v2-7f6ec7526783a49527733d1b6dbdb576_1440w.webp)

------

## 3. 克隆博客源仓库到本地

1. 打开想要在本地储存项目的文件夹（比如我的项目的文件夹是 `Websites`）：

   ```shell
   cd D:\DevTools\Hugo\Websites
   ```

2. 克隆**博客源仓库**到项目文件夹，克隆时使用的 HTTPS 仓库链接在这里查看：

   ![克隆博客源仓库到本地](https://pic1.zhimg.com/80/v2-260acddc6e0acbb6187797cfcf726784_1440w.webp)

   

   

   ```shell
   git clone https://github.com/Liam-Zhong/demo.git
   ```


------

## 4. 使用 Hugo 创建网站

1. 进入刚刚克隆下来的**博客源仓库**文件夹（比如：我的博客源仓库文件夹名是 `demo`，则 `cd demo` ），在这个文件夹里用 Hugo 创建一个网站文件夹；

2. 用 Hugo 创建网站文件夹的命令是 `hugo new site 网站名字`（比如，我的命名是 `my-blog`）。

   ```shell
   cd demo
   hugo new site my-blog
   ```
   

------

## 5. 安装和配置 Hugo 主题

### 5.1 选择 Hugo 主题

可以从 [Hugo 社区提供的主题](https://themes.gohugo.io/)中选择一个喜欢的主题应用在自己的网站中。


### 5.2 安装 Hugo 主题

1. 一般在你选择的 Hugo 主题的文档中，都会给出「如何安装这个主题」的命令，比如我选用的 Paper 的文档中给出：

   ![如何安装主题](https://pic1.zhimg.com/80/v2-92bff148a599f592c3fbad4e454cbaa4_1440w.webp)

   

2. 打开刚刚用 Hugo 创建的网站文件夹（我的是 my-blog），在终端（可以 `cd my-blog` 也可以在这个目录下直接右键并选择“在终端中打开”）粘贴文档中给出的安装命令。

   ```shell
   git submodule add https://github.com/nanxiaobei/hugo-paper themes/paper
   ```

3. 这时可以看到在 `themes` 文件夹中，多出了刚刚安装的主题文件，代表主题安装成功。


### 5.3 配置 Hugo 主题

1. 一般安装的 Hugo 主题的文件结构中都会有 exampleSite 文件夹，也是你在选择主题时参考的网站 demo；

2. 把 exampleSite 的文件复制到站点目录，在此基础上进行基础配置。 非常推荐这么做，这样做能解决很多「为什么明明跟教程一步一步做下来，显示的结果却不一样？」的疑惑（这主要是因为不同的主题模版配置文件不同导致的）。

3. 在把 exampleSite 文件复制到站点目录时，请根据**对应**文件夹进行复制文件：

   - 比如 exampleSite 下有 content、layouts、static 3 个文件，就找到你自己的站点跟目录下这对应的三个文件。再把对应目录中的内容分别复制过去。

     ![hugo文件结构](https://pic2.zhimg.com/80/v2-8b33d3e652e6be714b1e2a0ad325f109_1440w.webp)
     
   
4. 其中在复制 config.toml 的内容时要注意：

   - baseURL：

     ```shell
     baseURL = "https://example.com/" #把 https://example.com/ 改成自己的域名
     ```
     
     如果你没有在 GitHub Pages 中设置自定义域名，这里的域名应该填：

     ```
     https://<username>.github.io/
     ```

     （⚠️ 注意：最后的/不要忘了加）
     

------

## 6. 用 Hugo 创建文章

用 Hugo 创建一篇文章的命令是：

```shell
hugo new xxx.md
```

用这个命令创建的 Markdown 文件会套用 archetypes 文件夹中的 front matter 模版，在空白处用 Markdown 写入内容。

![创建的文章](https://pic3.zhimg.com/80/v2-9c1f132b5e56f0fa698fb89f315f44ce_1440w.webp)

`draft: true` 代表这篇文章是一个草稿，Hugo 不会显示草稿，要在主页显示添加的文章，可以设置 `draft: false`；或者直接删掉这行。

------

## 7. 本地调试和预览

1. 在发布到网站前可以在本地预览网站或内容的效果，运行命令：

   ```shell
   hugo server
   ```
   
   ![hugo server命令](https://pic3.zhimg.com/80/v2-5fede2c71c8b1e16d91230f7063ac6c2_1440w.webp)

2. 也可以在本地编辑 Markdown 文件时，通过 `hugo server` 来实时预览显示效果：

   ```shell
   hugo server
   ```

    运行成功后，可以在 http://localhost:1313/中预览网站

   ![预览页面](https://pic2.zhimg.com/80/v2-b4d43a48f7e89b455c5be36cbc074891_1440w.webp)

------

## 8. 发布内容

1. hugo 命令可以将你写的 Markdown 文件生成静态 HTML 网页，生成的 HTML 文件默认存放在 public 中；

   ```shell
   hugo
   ```
   
   ![hugo命令](https://pic3.zhimg.com/80/v2-f714d4d90fe506faa8de023546366626_1440w.webp)

2. 因为 hugo 生成的静态 HTML 网页文件默认存放在 public 文件中，所以推送网页内容只需要把 public 中的 HTML 网页文件发布到 GitHub Pages 仓库中；

3. 将 public 文件夹初始化为 Git 仓库，并设置默认主分支名为 main。这么做的原因是：

   - GitHub 创建仓库时生成的默认主分支名是 main；
   - 用 git init 初始化 Git 仓库时创建的默认主分支名是 master；
   - 将 git init 创建的 master 修改成 main，再推送给远端仓库 <username>.github.io ，这样才不会报错。

   ```shell
   cd public
   git init -b main
   ```
   
   ![初始化git命令](https://pic3.zhimg.com/80/v2-dc64e179de67231c963137d6ca2fb92e_1440w.webp)

4. 将 public 文件夹关联远程 GitHub Pages 仓库，使用 GitHub Pages 仓库的 SSH 链接；

   - （ ⚠️ 注意：要让 SSH 链接起作用，需要你添加过 SSH Key。如果你没有设置过 SSH Key，请参考[新增 SSH 密钥到 GitHub 帐户 - GitHub 文档](https://docs.github.com/zh/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)）

   - GitHub Pages 仓库的 SSH 链接可以在这里查看：

     ![SSH链接](https://pic3.zhimg.com/80/v2-a23dd95817f28f6f662c6174db505c86_1440w.webp)

   
   ```shell
   git remote add origin git@github.com:Liam-Zhong/Liam-Zhong.github.io.git
   ```
   
5. 推送博客源仓库的 public 文件夹中的 HTML 网页文件到 GitHub Pages 仓库中，在推送仓库内容前要先用 `git pull --rebase origin main` 和远端仓库同步，否则会报错；

   ```shell
   git pull --rebase origin main
   git add .
   git commit -m "...(changes)"
   git push origin main
   ```

6. 推送完后，可以在浏览器中输入 `https://<username>.github.io`，访问自己搭建的博客。

8. 后续的更新步骤：

   1. 创建你的文章 `xxx.md`；
   2. 用 `hugo server` 在本地预览，满意后准备发布；
   3. 运行 `hugo` 命令将 Markdown 文件生成 HTML 文件。
   4. 将修改先提交至博客源仓库

   ```shell
   git add .
   git commit -m "...(changes)"
   git push
   ```
   
   5. 打开 `public` 文件
   
   ```shell
   cd public
   ```
   
   6. 运行：
   
   ```shell
   git add .
   git commit -m "...(changes)"
   git push origin main
   ```
   
   以上。
