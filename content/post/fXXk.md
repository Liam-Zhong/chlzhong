+++
author = "Tuffy"
title = '奶奶的'
date = 2024-06-21T11:30:27+08:00
math = true                                
draft = false
comments = true
description = "{{ .Summary }}"
+++

這是習思也是奶奶的：<br>

這个 P 事 沒做過 再早一点 就爆破<br>
我想我沒听錯 你說你 bother 我<br>
問了問 你比我老很多<br>
有的聊很多 有的沒有說<br>
或者 心里想的 指沒照做<br>
收到 我的嘴巴在隱瞞什么<br>
目前對自己 不是太滿意<br>
一个謝謝 像在付時薪<br>
如果丢个不 怕你沒反應<br>
是否有得等了 是否空歡喜<br>
有時太軟弱 我有時太懶惰<br>
有時太敢說 有時又太晚說<br>
我說 回收到的那个不是我<br>
怎麽說呢 因爲他不愿承認<br>

奶奶的奶奶的奶奶的奶奶的！<br>
奶奶的奶奶的奶奶的奶奶的！<br>
奶奶的奶奶的奶奶的奶奶的！<br>

-----


<div style="display: flex; gap: 16px; justify-content: center; align-items: flex-start;">
  <img src="https://picx.zhimg.com/80/v2-168f302dbfdbabccac3077a5d391c773_1440w.png" alt="左圖" class="img-apple" style="flex: 1; min-width: 0;">
  <img src="https://picx.zhimg.com/80/v2-ea9d4879b9444b96468aeb61d12dd41f_1440w.png" alt="右圖" class="img-apple" style="flex: 1; min-width: 0;">
</div>


-----

附上解决方案（不知道 Excel 能不能直接做）

```python
import pandas as pd

df1 = pd.read_excel('data/A04232A1100071011-+成績上報 Excel 模板。xls')
df2 = pd.read_excel('data/雨課堂作业和期中成績。xlsx')

# 選擇需要列
df2 = df2[['姓名', '雨課堂作业折合分=总分/18*0.2', '期中成績']]
# 使用姓名列來合并两个 DF
df_merged = pd.merge(df1, df2, on='姓名', how='left')
#之所以 head 是要查看合并后的名字
df_merged.head()
```

```python
# 更新成績列
df1['雨課堂作业折合分=总分/18*0.2'] = df_merged['雨課堂作业折合分=总分/18*0.2_y']
df1['期中成績'] = df_merged['期中成績_y']

# 保存更新后的 Excel 文件
df1.to_excel('更新后的成績表。xlsx', index=False)

print("奶奶的老登")
```

