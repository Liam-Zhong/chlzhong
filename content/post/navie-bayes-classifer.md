+++
author = "Tuffy"
title = '朴素贝叶斯算法与垃圾短信识别'
date = 2024-12-03T21:17:52+08:00
math = true 
draft = false
comments = true
toc = false
description = "{朴素贝叶斯算法与垃圾短信识别}"

+++

关于朴素贝叶斯算法的介绍可谓前人之述备矣，在此不多加阐述。只要了解概率论中的全概率公式即可。这里主要结合一个具体例子进行分析，同时也是我的机器学习作业。

<center>$P(B|A)=\frac{P(A|B)⋅P(B)}{P(A)}$</center>

<br>

<center>$P(A|B_1,B_2,...,B_n) = \frac{P(B_1,B_2,...,B_n|A) \cdot P(A)}{P(B_1,B_2,...,B_n)}$</center>

在运用之前要知道我们待处理问题的基本模型：

> 输入：一条样本

> 模型：训练样本（大量带标签的文本信息）

> 输出：布尔型答案（是垃圾短信/不是垃圾短信）

作为输入的样本需进行分词处理，即是把一条“尊敬的客户，您的手机将要爆炸，请充值话费。”变成“`尊敬的` `客户` `您的` `手机` `将要` `爆炸` `请` `充值` `话费`”。

若设 $X$ 为词向量（统计量），$x_i$ 为第 $i$ 个词向量（样本）， $y_j$ 代表分类结果。另设 $X^m$ 代表词向量中的第 $m$ 个元素，不难得：

<center>$P(Y=y_j|X=X_i)=\frac{P(X=x_i|Y=y_j)⋅P(Y=Y_j)}{P(X=X_i)}$</center>

此时假设所有样本互相独立（即所谓“朴素”），又有：

<center>$P(Y=y_j|X=X_i)=\frac{P(X^1=x_i^{(1)}, \cdots ,X^m=x_i^{(m)}|Y=y_j)⋅P(Y=Y_j)}{P(X=X_i)}$</center>
<br>

<center>$\quad = \frac{\prod_{s=1}^m P(X^{(s)}=x_i^{(s)}|Y=y_j)⋅P(Y=Y_j)}{P(X=X_i)}$</center>

此时的分类器显然只用关心那些分母最大的取值（$argmax$ 与 $max$ 区别在于前者关心输入即 $y_j$S）：

<center>$y = f(x_i) = arg \underset{y_i}{max} = \prod_{s=1}^m P(X^{(s)}=x_i^{(s)}|Y=y_j)⋅P(Y=Y_j)$</center>

最终，用极大似然估计思想实现上式关键部分：

<center>$\prod_{s=1}^m P(X^{(s)}=x_i^{(s)}|Y=y_j)⋅P(Y=Y_j)$</center>

<br>

<center>$\quad=\prod_{s=1}^m P(X^{(s)}=\frac{y_j \text{类中第 s 个特征值为}x_i^{s}的样本的个数}{y_j \text{类中样本的个数}} \cdot \frac{\text{训练集中属于} y_j {类的样本个数}}{\text{训练集中样本个数}}$</center>

如何解决某特征值为 0 的问题呢？在分子分母引入一个 $\lambda$ 即可，取个好听的名字叫拉普拉斯平滑因子。

代码实现：

```python
from numpy import *
from functools import reduce

adClass = 1#广告、垃圾标识
def loadDataSet():
    wordsList = [
        ['周六', '公司', '一起', '聚餐', '时间'],
        ['优惠', '返利', '打折', '优惠', '金融', '理财'],
        ['喜欢', '机器学习', '一起', '研究', '欢迎', '贝叶斯', '算法', '公式'],
        ['公司', '发票', '税点', '优惠', '增值税', '打折'],
        ['北京', '今天', '雾霾', '不宜', '外出', '时间', '在家', '讨论', '学习'],
        ['招聘', '兼职', '日薪', '保险', '返利']
    ]
    classVec = [0, 1, 0, 1, 0, 1]
    return wordsList, classVec

def doc2VecList(docList):
    a = list(reduce(lambda x, y: set(x) | set(y), docList))
    return a

def words2Vec(vecList, inputWords):
    resultVec = [0] * len(vecList)
    
    for word in inputWords:
        if word in vecList:
            resultVec[vecList.index(word)] = 1
    return array(resultVec)

def trainNB(trainMatrix, trainClass):
    numTrainClass = len(trainClass)
    numWords = len(trainMatrix[0])
    
    # 初始化概率向量，使用拉普拉斯平滑
    p0Num = ones(numWords)  # 非垃圾邮件类的词频统计
    p1Num = ones(numWords)  # 垃圾邮件类的词频统计
    p0Words = 2.0          # 非垃圾邮件类的总词数
    p1Words = 2.0          # 垃圾邮件类的总词数
    
    # 统计词频
    for i in range(numTrainClass):
        if trainClass[i] == 1:
            p1Num += trainMatrix[i]
            p1Words += sum(trainMatrix[i])
        else:
            p0Num += trainMatrix[i]
            p0Words += sum(trainMatrix[i])
 
    p0Vec = log(p0Num / p0Words)
    p1Vec = log(p1Num / p1Words)

    pClass1 = sum(trainClass) / float(numTrainClass)
    
    return p0Vec, p1Vec, pClass1

def classifyNB(testVec, p0Vec, p1Vec, pClass1) :
    p1 = sum(testVec * p1Vec) + log (pClass1)
    p0 = sum(testVec * p0Vec) + log(1 - pClass1)
    if p0 > p1:
        return 0
    return 1

def printClass(words, testClass):
    if testClass == adClass:
        print(words, 'ad')
    else:
        print(words, 'notad')

def tNB() :
    docList, classVec = loadDataSet()
    allWordsVec = doc2VecList (docList)
    trainMat = list (map(lambda x: words2Vec (allWordsVec,x), docList))
    p0V, p1V, pClass1 = trainNB (trainMat, classVec)
    testWords = ['公司','聚餐','讨论','贝叶斯'] 
    testVec = words2Vec (allWordsVec,testWords)
    testClass = classifyNB(testVec, p0V, p1V, pClass1)
    printClass (testWords, testClass)
    testWords = ['公司','保险','金融']
    testVec = words2Vec (allWordsVec, testWords)
    testClass = classifyNB(testVec, p0V, p1V, pClass1)
    printClass (testWords, testClass)

if __name__ == '__main__':
    tNB()
```

