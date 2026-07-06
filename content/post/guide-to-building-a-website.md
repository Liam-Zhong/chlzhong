+++
author = "Tuffy"
title = '建站指北'
date = 2024-05-26T11:20:57+08:00
comments = true
toc = true
description = "{{ .Summary }}"
+++

#### 原文见[如何用 GitHub Pages + Hugo 搭建个人博客 · 小綿尾巴 (cuttontail.blog)](https://cuttontail.blog/blog/create-a-wesite-using-github-pages-and-hugo/)

#### 整个過程基于 Windows11

----

## 0.部落格 

### 部落格是什么？
>部落格（博客）是一种在線日記型式的个人网站，借由張帖子章、圖片或影片來記錄生活、抒发情感或分享信息。博客上的文章通常根据張貼時间，以倒序方式由新到舊排列。許多博客作者專注評論特定的課題或新聞，其他則作爲个人日記。一个典型的博客結合了文字、圖像、其他博客或网站的超鏈接、及其它与主題相關的媒体。能够讓讀者以互动的方式留下意见，是許多博客的重要要素。大部分的博客內容以文字爲主，也有一些博客專注艺術、攝影、视頻、音樂、播客等各种主題。博客是社会媒体网絡的一部分。——[維基百科](https://zh.wikipedia.org/wiki/%E7%B6%B2%E8%AA%8C)

一直認爲一件事情如果你忘記了那么它就不存在。「記錄」从古至今都是很有仪（niu）式（ben）感的事情：登高作賦，臨泉賦詩，蘭亭修序……甚至還要画張山水裱起來，聊以服人。如今人們愛用朋友圈（个人很鄙夷路上奔波之后的所謂「又是自由的一天」）記錄生活也差不多是這个道理，只是依賴于微信总讓人不舒服，因此自己搭一个部落格是一件很棒很酷的事情，否則怎么算会「上网」呢？本文便是出于這样的目的，争取讓每个人都可以零基礎搭建自己的部落格。



這篇教程假设你已经：

1. 安装了 [Git](https://git-scm.com/)，并且了解基本的 Git 知識；
2. 有一个 [GitHub](https://github.com/) 賬号；
3. 有自己偏好的文本編輯器（比如 [Typora](https://typora.io/)）。

如果你滿足以上条件**那就開始吧！**

------

## 1. 安装 Hugo

1. 点擊 [Release v0.126.1 · gohugoio/hugo (github.com)]((https://github.com/gohugoio/hugo/releases/tag/v0.126.1)) 并下滑找到 hugo_0.126.1_windows-amd64.zip，点擊下載。（本文寫于 2024 年 5 月 26 号，你下載的時候版本可能已经更新。）
   ![版本選擇](https://pic3.zhimg.com/80/v2-e0d9fd92a6ed78579e4dbaf223d20602_1440w.webp)

2. 把下載好的 .zip 文件解壓，并放置在你想放的文件夾中（我的路徑：D:\DevTools\Hugo，也就是說現在 Hugo 目錄下有一个名爲 hugo_0.126.1_windows-amd64 的文件夾）。打開设置，搜索「編輯系统環境变量」，在彈出的「系统属性」窗口中点擊「環境变量」，再在彈出的「環境变量」窗口的「系统变量」中選中 Path 路徑，并点擊編輯。通過右側的「新建」按鈕，將 `D:\DevTools\Hugo\hugo_0.126.1_windows-amd64` 添加到環境变量（根据自己的路徑修改）。

2. WIN+D 回到桌面，右鍵選擇「在終端中打開」。输入：

   ```shell
   hugo version
   ```
   一切正常的話会顯示你目前安装的版本号：
   
   ![hugo version](https://pic3.zhimg.com/80/v2-0da4527059ac8dee7484e9ee28be7446_1440w.webp)

------

## 2. 創建 GitHub 倉庫

### 2.1 創建博客源倉庫

1. 命名**博客源倉庫**（我使用的是 `demo`）;
2.  勾選 **Public**，设置爲公開倉庫；
3.  勾選添加 **README** 文件。

![創建demo倉庫](https://pic3.zhimg.com/80/v2-c718ed26fa965c2b5f9dcf16a59407a6_1440w.webp)

### 2.2 創建 GitHub Page 倉庫

1. 命名 **GitHub Pages** 倉庫，這个倉庫必須使用特殊的命名格式 `<username.github.io>` ，`<username>` 是你自己的 GitHub 的用户名；
2.  勾選 **Public**，设置爲公開倉庫；
3.  勾選添加 **README** 文件，這会设置 `main` 分支爲倉庫的默認主分支，在后面提交推送博客內容時很重要。

![創建name.github.io倉庫](https://pic3.zhimg.com/80/v2-7f6ec7526783a49527733d1b6dbdb576_1440w.webp)

------

## 3. 克隆博客源倉庫到本地

1. 打開想要在本地儲存項目的文件夾（比如我的項目的文件夾是 `Websites`）：

   ```shell
   cd D:\DevTools\Hugo\Websites
   ```

2. 克隆**博客源倉庫**到項目文件夾，克隆時使用的 HTTPS 倉庫鏈接在這里查看：

   ![克隆博客源倉庫到本地](https://pic1.zhimg.com/80/v2-260acddc6e0acbb6187797cfcf726784_1440w.webp)

   

   

   ```shell
   git clone https://github.com/Liam-Zhong/demo.git
   ```


------

## 4. 使用 Hugo 創建网站

1. 進入剛剛克隆下來的**博客源倉庫**文件夾（比如：我的博客源倉庫文件夾名是 `demo`，則 `cd demo` ），在這个文件夾里用 Hugo 創建一个网站文件夾；

2. 用 Hugo 創建网站文件夾的命令是 `hugo new site 网站名字`（比如，我的命名是 `my-blog`）。

   ```shell
   cd demo
   hugo new site my-blog
   ```
   

------

## 5. 安装和配置 Hugo 主題

### 5.1 選擇 Hugo 主題

可以从 [Hugo 社區提供的主題](https://themes.gohugo.io/)中選擇一个喜歡的主題應用在自己的网站中。


### 5.2 安装 Hugo 主題

1. 一般在你選擇的 Hugo 主題的文檔中，都会給出「如何安装這个主題」的命令，比如我選用的 Paper 的文檔中給出：

   ![如何安装主題](https://pic1.zhimg.com/80/v2-92bff148a599f592c3fbad4e454cbaa4_1440w.webp)

   

2. 打開剛剛用 Hugo 創建的网站文件夾（我的是 my-blog），在終端（可以 `cd my-blog` 也可以在這个目錄下直接右鍵并選擇「在終端中打開」）粘貼文檔中給出的安装命令。

   ```shell
   git submodule add https://github.com/nanxiaobei/hugo-paper themes/paper
   ```

3. 這時可以看到在 `themes` 文件夾中，多出了剛剛安装的主題文件，代表主題安装成功。


### 5.3 配置 Hugo 主題

1. 一般安装的 Hugo 主題的文件結构中都会有 exampleSite 文件夾，也是你在選擇主題時参考的网站 demo；

2. 把 exampleSite 的文件复制到站点目錄，在此基礎上進行基礎配置。 非常推荐這么做，這样做能解决很多「爲什么明明跟教程一步一步做下來，顯示的結果却不一样？」的疑惑（這主要是因爲不同的主題模版配置文件不同導致的）。

3. 在把 exampleSite 文件复制到站点目錄時，請根据**對應**文件夾進行复制文件：

   - 比如 exampleSite 下有 content、layouts、static 3 个文件，就找到你自己的站点跟目錄下這對應的三个文件。再把對應目錄中的內容分別复制過去。

     ![hugo文件結构](https://pic2.zhimg.com/80/v2-8b33d3e652e6be714b1e2a0ad325f109_1440w.webp)
     
   
4. 其中在复制 config.toml 的內容時要注意：

   - baseURL：

     ```shell
     baseURL = "https://example.com/" #把 https://example.com/ 改成自己的域名
     ```
     
     如果你沒有在 GitHub Pages 中设置自定义域名，這里的域名應該填：

     ```
     https://<username>.github.io/
     ```

     （⚠️ 注意：最后的/不要忘了加）
     

------

## 6. 用 Hugo 創建文章

用 Hugo 創建一篇文章的命令是：

```shell
hugo new xxx.md
```

用這个命令創建的 Markdown 文件会套用 archetypes 文件夾中的 front matter 模版，在空白处用 Markdown 寫入內容。

![創建的文章](https://pic3.zhimg.com/80/v2-9c1f132b5e56f0fa698fb89f315f44ce_1440w.webp)

`draft: true` 代表這篇文章是一个草稿，Hugo 不会顯示草稿，要在主页顯示添加的文章，可以设置 `draft: false`；或者直接删掉這行。

------

## 7. 本地調試和預覽

1. 在发布到网站前可以在本地預覽网站或內容的效果，运行命令：

   ```shell
   hugo server
   ```
   
   ![hugo server命令](https://pic3.zhimg.com/80/v2-5fede2c71c8b1e16d91230f7063ac6c2_1440w.webp)

2. 也可以在本地編輯 Markdown 文件時，通過 `hugo server` 來实時預覽顯示效果：

   ```shell
   hugo server
   ```

    运行成功后，可以在 http://localhost:1313/中預覽网站

   ![預覽页面](https://pic2.zhimg.com/80/v2-b4d43a48f7e89b455c5be36cbc074891_1440w.webp)

------

## 8. 发布內容

1. hugo 命令可以將你寫的 Markdown 文件生成靜態 HTML 网页，生成的 HTML 文件默認存放在 public 中；

   ```shell
   hugo
   ```
   
   ![hugo命令](https://pic3.zhimg.com/80/v2-f714d4d90fe506faa8de023546366626_1440w.webp)

2. 因爲 hugo 生成的靜態 HTML 网页文件默認存放在 public 文件中，所以推送网页內容只需要把 public 中的 HTML 网页文件发布到 GitHub Pages 倉庫中；

3. 將 public 文件夾初始化爲 Git 倉庫，并设置默認主分支名爲 main。這么做的原因是：

   - GitHub 創建倉庫時生成的默認主分支名是 main；
   - 用 git init 初始化 Git 倉庫時創建的默認主分支名是 master；
   - 將 git init 創建的 master 修改成 main，再推送給遠端倉庫 <username>.github.io ，這样才不会報錯。

   ```shell
   cd public
   git init -b main
   ```
   
   ![初始化git命令](https://pic3.zhimg.com/80/v2-dc64e179de67231c963137d6ca2fb92e_1440w.webp)

4. 將 public 文件夾關联遠程 GitHub Pages 倉庫，使用 GitHub Pages 倉庫的 SSH 鏈接；

   - （ ⚠️ 注意：要讓 SSH 鏈接起作用，需要你添加過 SSH Key。如果你沒有设置過 SSH Key，請参考[新增 SSH 密鑰到 GitHub 帳户 - GitHub 文檔](https://docs.github.com/zh/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)）

   - GitHub Pages 倉庫的 SSH 鏈接可以在這里查看：

     ![SSH鏈接](https://pic3.zhimg.com/80/v2-a23dd95817f28f6f662c6174db505c86_1440w.webp)

   
   ```shell
   git remote add origin git@github.com:Liam-Zhong/Liam-Zhong.github.io.git
   ```
   
5. 推送博客源倉庫的 public 文件夾中的 HTML 网页文件到 GitHub Pages 倉庫中，在推送倉庫內容前要先用 `git pull --rebase origin main` 和遠端倉庫同步，否則会報錯；

   ```shell
   git pull --rebase origin main
   git add .
   git commit -m "...(changes)"
   git push origin main
   ```

6. 推送完后，可以在瀏覽器中输入 `https://<username>.github.io`，訪問自己搭建的博客。

8. 后續的更新步驟：

   1. 創建你的文章 `xxx.md`；
   2. 用 `hugo server` 在本地預覽，滿意后准备发布；
   3. 运行 `hugo` 命令將 Markdown 文件生成 HTML 文件。
   4. 將修改先提交至博客源倉庫

   ```shell
   git add .
   git commit -m "...(changes)"
   git push
   ```
   
   5. 打開 `public` 文件
   
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
