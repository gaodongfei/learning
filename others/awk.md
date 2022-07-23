# awk

## 字段的引用

- awk中使用$1 $2 ... $n 表示每一个字段
  - awk '{print $1,$2}' filename
- awk 可以使用-F选项改变字段分隔符（默认为空格分隔符）
  - awk -F ',' '{ print $1,$2,$3}' filename
  - 分割符可以使用正则表达式

### 演示

```bash
awk '/^menu/{print $0}' /boot/grub2/grub.cfg
awk -F "'" '/^menu/{print $2}' /boot/grub2/grub.cfg
awk -F "'" '/^menu/{print x++,$2}' /boot/grub2/grub.cfg 
```

## awk表达式

### 赋值操作符

- =是最常用的赋值操作符
  - var1 = "name"
  - var2 = "hello" "world"
  - var3 = $1
- 其他赋值操作符
  - ++ -- += -= *= /= %= ^=

### 算数操作符

- +-*/

### 系统变量

- FS 和OFS 字段分隔符
  - FS 输入字段分隔符
  - OFS 输出字段分隔符
- RS 记录分隔符
- NR 和FNR行数
- NF字段数量，最后一个字段内容可以用$NF

### 演示

```bash
# FS输入字段分隔符，等于-F，BEGIN在读入文件之前做的操作
head -5 /etc/passwd | awk 'BEGIN{FS=":"}{print $1,$2}'
# OFS输出字段分隔符，默认为空格
head -5 /etc/passwd | awk 'BEGIN{FS=":";OFS="-"}{print $1,$2}'
# RS标记输入行的分隔符，默认为/n，改了之后行结束规则改变
head -5 /etc/passwd | awk 'BEGIN{RS=":"}{print $0}'
# NR显示行号，多个文件，行号会累加
head -5 /etc/passwd | awk '{print NR,$0}'
# FNR显示行号，但区分文件，每个文件从0重新开始
head -5 /etc/passwd | awk '{print FNR,$0}'
# NF每行字段数量,$NF显示最后一个字段内容
head -5 /etc/passwd | awk 'BEGIN{FS=":"}{print NF,$NF}'
```



### 关系操作符

- < > <= >= == != ~（匹配） !~（不匹配）

### 布尔操作符

- && || ！

## awk的条件和循环

kpi.txt 

```
user1 70 72 74 76 74 72
user2 80 82 84 82 80 78
user3 60 61 62 63 64 65
user4 90 89 88 87 86 85
user5 45 60 63 62 61 50
```

### 演示

```bash
# 第二列大于等于80，输出第一列。这里if后面没有；print是在if后面执行
awk '{if($2>=80) {print $1}}' kpi.txt 
# if 后面有多行语句，需要大括号
 awk '{if($2>=80) {print $1;print $2}}' kpi.txt
# 求总成绩和平均成绩
 awk '{sum=0;for(i=2;i<=NF;i++){sum+=$i}print $1,sum,sum/(NF-1)}' kpi.txt 
```

## awk 数组

### 演示

```bash
# 使用数组保留平均成绩，用end汇总输出
awk '{sum=0;for(i=2;i<=NF;i++){sum+=$i} average[$1]=sum/(NF-1)}END{for (user in average) sum2+=average[user];print sum2/NR}' kpi.txt 
# 保留awk脚本
vim ave.awk
# 使用脚本
awk -f ave.awk kpi.txt
```

### 命令行参数数组

- ARGC 命令行参数数量

- ARGV 具体命令行参数数组

```awk
BEGIN{
  for (i=0;i<ARGC;i++){
  	print ARGV[i]
  }
  print ARGC
}
```

