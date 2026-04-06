+++
author = "Tuffy"
title = '奶奶的'
date = 2024-06-21T11:30:27+08:00
math = true                                
draft = false
comments = true
description = "{{ .Summary }}"
+++

这是习思也是奶奶的：<br>

这个 P 事 没做过 再早一点 就爆破<br>
我想我没听错 你说你 bother 我<br>
问了问 你比我老很多<br>
有的聊很多 有的没有说<br>
或者 心里想的 指没照做<br>
收到 我的嘴巴在隐瞒什么<br>
目前对自己 不是太满意<br>
一个谢谢 像在付时薪<br>
如果丢个不 怕你没反应<br>
是否有得等了 是否空欢喜<br>
有时太软弱 我有时太懒惰<br>
有时太敢说 有时又太晚说<br>
我说 回收到的那个不是我<br>
怎麽说呢 因为他不愿承认<br>

奶奶的奶奶的奶奶的奶奶的！<br>
奶奶的奶奶的奶奶的奶奶的！<br>
奶奶的奶奶的奶奶的奶奶的！<br>

-----


<div align="center">    
<img src="https://picx.zhimg.com/80/v2-168f302dbfdbabccac3077a5d391c773_1440w.png" alt="50%" width="50%" height="auto">
<img src="https://picx.zhimg.com/80/v2-ea9d4879b9444b96468aeb61d12dd41f_1440w.png" alt="50%" width="50%" height="auto">
</div>



-----

附上解决方案（不知道 Excel 能不能直接做）

```python
import pandas as pd

df1 = pd.read_excel('data/A04232A1100071011-+成绩上报 Excel 模板。xls')
df2 = pd.read_excel('data/雨课堂作业和期中成绩。xlsx')

# 选择需要列
df2 = df2[['姓名', '雨课堂作业折合分=总分/18*0.2', '期中成绩']]
# 使用姓名列来合并两个 DF
df_merged = pd.merge(df1, df2, on='姓名', how='left')
#之所以 head 是要查看合并后的名字
df_merged.head()
```

```python
# 更新成绩列
df1['雨课堂作业折合分=总分/18*0.2'] = df_merged['雨课堂作业折合分=总分/18*0.2_y']
df1['期中成绩'] = df_merged['期中成绩_y']

# 保存更新后的 Excel 文件
df1.to_excel('更新后的成绩表。xlsx', index=False)

print("奶奶的老登")
```

