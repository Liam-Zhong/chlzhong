+++
author = "Tuffy"
title = '線性回歸'
date = 2024-12-05T12:32:24+08:00
math = true 
draft = false
comments = true
toc = true
description = "{線性回歸}"

+++

線性回歸理解起來很簡單，當参數样本只有一个特征与一个標籤時就是高中学最小二乘法的那个模型，两个特征及以上時則是大学概率論上学的参數估計，只不過估計参數時選用了梯度下降法。

值得注意的是損失函數的選擇，爲什么是**平方誤差函數**而不是差的絕對值呢？這里其实是因爲我們假定偏差的分布符合高斯分布，其對應的概率密度函數就是平方和形式，而且**差的絕對值**（拉普拉斯分布）所對應的极大似然函數沒法方便尋优。

而說到梯度下降，在工程实現方面其实也不難，具体见后文代碼，主要談談背后的數学。

### 0.梯度是什么

對于一个多元函數，梯度是該函數所有偏導數的向量，指向函數增长最快的方向，大小表示增长的速率。

假设有一个函數 $f(x_1,x_2,…,x_n)$，梯度可以表示爲一个向量：

$\nabla f = \left( \frac{\partial f}{\partial x_1}, \frac{\partial f}{\partial x_2}, \dots, \frac{\partial f}{\partial x_n} \right)$

這里，$\frac{\partial f}{\partial x_i}$ 是函數 $f$ 對每个变量 $x_i$ 的偏導數。

### 1. **梯度下降与泰勒展開的關系**

先解釋下爲什么要有梯度下降法：其实最簡單的二維凸函數是拋物線 $f(x)=x^2$，很容易通過解方程 $f'(x)=0$ 求出最小值在 $x=0$ 处；只是有一些凸函數這样解方程太麻煩，便用梯度下降法來找最值。

最簡單情况 ($f(x)=x^2$) 下，若將給定点 $x_0$ 加上 -$\eta\nabla f(x_0)$，就相當于一个逐漸靠近最低点的物理過程。比如取 $x_0$ = 10，$\eta = 0.2$ ，迭代 10 次左右就是差不多靠近了最低点 $x=0$

>初始值 x=10, f(x)=100<br>
>第 1 次迭代: x=6.0, f(x)=36.0<br>
>第 2 次迭代: x=3.6, f(x)=12.96<br>
>第 3 次迭代: x=2.16, f(x)=4.67<br>
>第 4 次迭代: x=1.296, f(x)=1.68<br>
>第 5 次迭代: x=0.7776, f(x)=0.60<br>
>第 6 次迭代: x=0.46656, f(x)=0.22<br>
>第 7 次迭代: x=0.27994, f(x)=0.078<br>
>第 8 次迭代: x=0.16796, f(x)=0.028<br>
>第 9 次迭代: x=0.10078, f(x)=0.0102<br>
>第 10 次迭代: x=0.06047, f(x)=0.00366<br>

不難理解這个 $\eta$ （取个名字叫学習率）取得太小則迭代次數過多，太大則会越過最低点不斷震蕩。

在实际中的函數沒有這么簡單，更复雜的算式里梯度下降沿着函數的梯度方向進行优化，而梯度本身就是函數在當前位置的泰勒展開的一階導數。此時梯度下降法可以看作是在每一步使用泰勒展開的一階近似來更新参數。具体來說，梯度下降可以看作是泰勒展開的一階近似的迭代應用。

### 2. **梯度下降的解析解**

梯度下降法**沒有解析解**，因爲它是一种數值优化方法。（解析解和數值解的區別不知道的話何承春会头疼呢）

返回的結果就是使目標函數取得最小值的**参數**。

接下來通過一个具体的例子來解釋如何理解**批量梯度下降**、**小批量梯度下降**和**随机梯度下降**。

假设我們有一个三元函數（即目標函數是有三个参數的函數），并且我們有100个样本數据。我們希望通過這些數据來优化函數的参數。

##### 1. 全批量梯度下降（Batch Gradient Descent）

全批量梯度下降会在每次迭代中使用**所有 100 个样本**來計算梯度，并更新参數。

> 1. 从初始参數 $\theta_0$ 開始。<br>
> 2. 計算所有 100 个样本点的梯度，得到梯度平均值：<br>
$\nabla f(\theta) = \frac{1}{100} \sum_{i=1}^{100} \nabla f_i(\theta)$<br>
> 3. 更新参數：<br>
> $\theta_{k+1} = \theta_k - \eta \nabla f(\theta_k)$

##### 2. 随机梯度下降（Stochastic Gradient Descent, SGD）

与全批量梯度下降不同，随机梯度下降每次迭代僅使用**一个样本**來計算梯度并更新参數。對于 100 个样本，随机梯度下降每次只随机選擇一个样本進行参數更新。

> 1. 从初始参數 $\theta_0$  開始。<br>
> 2. 随机選擇一个样本 $i$，計算該样本的梯度 $\nabla f_i(\theta)$。<br>
> 3. 使用該样本的梯度更新参數：<br>
>$\theta_{k+1} = \theta_k - \eta \nabla f_i(\theta_k)$
> 4. 重复這个過程，直到达到停止条件

##### 3. 小批量梯度下降（Mini-batch Gradient Descent）

小批量梯度下降是全批量梯度下降和随机梯度下降的折中方法。每次迭代時，它会从所有 100 个样本中随机選取一个**小批量**（例如，10 个样本），用這些样本來計算梯度并更新参數。

> 1. 从初始参數 $\theta_0$ 開始。<br>
> 2. 随机選擇一个小批量样本，假设這个小批量包含 10 个样本 $i_1, i_2, \dots, i_{10}$。<br>
> 3. 計算這 10 个样本的梯度平均值：<br>
>$\nabla f(\theta) = \frac{1}{10} \sum_{i=1}^{10} \nabla f_{i}(\theta)$
> 4. 使用這个平均梯度更新参數：<br>
>$\theta_{k+1} = \theta_k - \eta \nabla f(\theta_k)θk+1=θk−η∇f(θk)$
> 5. 重复這个過程，直到达到停止条件。



代碼实現：

```python
#%%
import random
import torch
from matplotlib import pyplot as plt
#%%
# 生成随机權重
def get_w_b():
    w = torch.normal(0, 0.01, size=(1, 1), requires_grad=True)
    b = torch.zeros(1, requires_grad=True)
    return w, b
#%%
# 生成訓練數据
def synthetic_data(w, b, num_examples):
    """生成y = Xw + b + 噪声"""
    X = torch.normal(0, 1, (num_examples, len(w)))
    y = torch.matmul(X, w) + b
    y += torch.normal(0, 0.01, y.shape)
    return X, y.reshape((-1, 1))
#%%
# 生成迭代器
def data_iter(batch_size, features, labels):
    num_examples = len(features)
    indices = list(range(num_examples))
    # 這些样本是随机讀取的，沒有特定的順序 
    random.shuffle(indices)
    for i in range(0, num_examples, batch_size):
        batch_indices = torch.tensor(indices[i:min(i + batch_size, num_examples)])
        yield features[batch_indices], labels[batch_indices]
#%% md
## 自己完成
#%%
# 線性回歸模型
def Linear_regression(X, w, b):
    return X*w + b #return torch.matmul(X, w) + b更好
#%%
# 均方損失函數 MSE
def squared_loss(y_hat, y):
    return (y_hat - y) ** 2 / 2
#%%
# 优化器实現
def sgd(params, lr, batch_size):
    """小批量随机梯度下降"""
    with torch.no_grad():
        for param in params:
            param -= lr * param.grad / batch_size
            param.grad.zero_()
#%% md
## 測試
#%%
lr = 0.03
num_epochs = 3
net = Linear_regression
loss = squared_loss
batch_size = 10

# y = -5 * x + 0.1 
true_w = torch.tensor([-5.0])
true_b = 0.1
features, labels = synthetic_data(true_w, true_b, 1000)
w, b = get_w_b()

for epoch in range(num_epochs):
    for X, y in data_iter(batch_size, features, labels):
        l = loss(net(X, w, b), y)
        l.sum().backward()
        sgd([w, b], lr, batch_size)
    with torch.no_grad():
        train_l = loss(net(features, w, b), labels)
        print(f'epoch {epoch + 1}, loss {float(train_l.mean()):f}')
        print(f'w {w}, b {b}')
#%% md
## 繪制圖形
#%%
x = features
y = labels

plt.scatter(x, y, label='Samples')
plt.plot(x.detach().numpy(), w.detach().numpy() * x.detach().numpy() + b.detach().numpy(), c='r', label='True function')
plt.plot(x, -5 * x + 0.1, c='b', label='Trained model')
plt.title("CQUPT2022212062", loc="center")
plt.legend()
plt.show()
#%%
print("權重",w)
#%%

```

