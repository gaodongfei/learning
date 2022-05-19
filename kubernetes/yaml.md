## 申明多行字符串

1. 多行字符串表示|

   ```yaml
   # Latte\nCappuccino\nEspresso\n
   coffee: |
     Latte
     Cappuccino
     Espresso  
   ```

2. 控制多行字符串中的空格 |-

   ```yaml
   # Latte\nCappuccino\nEspresso
   coffee: |-
     Latte
     Cappuccino
     Espresso 
     
   # Latte\n 12 oz\n 16 oz\nCappuccino\nEspresso
   coffee: |-
     Latte
       12 oz
       16 oz
     Cappuccino
     Espresso  
   ```

3. 保留尾随空格 |+

   ```yaml
   # Latte\nCappuccino\nEspresso\n\n\n
   coffee: |+
     Latte
     Cappuccino
     Espresso  
   
   
   another: value
   ```

4. 折叠多行字符串

   ```yaml
   # Latte Cappuccino Espresso\n
   coffee: >
     Latte
     Cappuccino
     Espresso  
   # Latte\n 12 oz\n 16 oz\nCappuccino Espresso
   coffee: >-
     Latte
       12 oz
       16 oz
     Cappuccino
     Espresso 
     
   ```

## 在一个文件中嵌入多个文档

可以将多个YAML文档放在单个文件中。文档前使用---，文档后使用...

```yaml
---
document:1
...
---
document:2
...
```

## YAML 锚点

YAML规范存储了一种引用值的方法，然后通过引用指向该值。YAML称之为“锚定”：

```yaml
coffee: "yes, please"
favorite: &favoriteCoffee "Cappucino"
coffees:
  - Latte
  - *favoriteCoffee
  - Espresso
```

