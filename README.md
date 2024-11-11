# 项目介绍

这是一个基于 Python 和 Go 的本地算法测试小工具。

## 使用方法

1. 将测试的算法与标准答案分别放入 `utils` 目录下。
2. 通过 `cases.py` 生成测试数据（这部分需要一定的自定义）。
3. 运行 `main.go` 生成相应的结果文件。（这里也是，因为不同的算法需要的输入输出数据不一样）
4. 运行`analyse.py`生成时间复杂度分析图表。

## 目录结构

```
/utils
    - algorithm.go
    - standard_answer.go
/cases.py
/analyse.py
/main.go
```

希望这个工具能帮助你更方便地进行算法测试。
