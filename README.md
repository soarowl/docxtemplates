# docxtemplates使用说明

## 功能

将`template.docx`模板中的字段用`data.txt`中的数据进行替换。

## 说明

1. 模板中需要被替换的地方用`{学号}`，`{姓名}`等标识，这些字段放在`data.txt`中第一行，所有内容不能有空格。
1. `data.txt`采用`utf-8`格式保存，形式如下：

```txt
学号 姓名 总分 评价
211124010104 张三 95 项目做的很不错，希望百尺竿头更进一步
211124010130 胡四 80 项目基本完成，加油！
211124010135 李五 65 项目勉强能运行
211124010137 江六 50 仅仅入门
211124010144 鲁七 10 不知你在学什么
```

1. [源码及可执行程序](docxtemplates.7z)
