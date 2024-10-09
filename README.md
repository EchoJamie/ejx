<h1 align="center">
  EJX
</h1>

<h3 align="center">
  常用命令整合工具(个人常用)
</h3>

#### 基本用法

```shell
```shell
# 查看所有命令; 在不同模式下, 显示可能不同
$ ejx -h
```

##### 模式命令

```shell
# 模式命令
$ ejx mode -h
# 查看模式
$ ejx mode -v
# 设置模式
$ ejx mode --set <mode>
```

##### hexo命令

```shell
# 切换至 hexo 模式 (默认)
$ ejx mode --set hexo
# hexo 初始化: 设置博客目录
$ ejx init
# 编写新的文章: 交互式操作
$ ejx new
# 启动预览服务: 默认不包含草稿
$ ejx run
# 发布文章: 交互式操作, 输入数字选择文章
$ ejx publish
```
