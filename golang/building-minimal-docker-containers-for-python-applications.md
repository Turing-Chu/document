# 为 Python 应用构建最小化 Docker 容器

> [Nick Joyce][author] at 2018-03-19
>
> Translated by Turing Zhu
> 
> Source article: [Building Minimal Docker Containers for Python Applications][source]


[author]: https://medium.com/@nick.joyce?source=post_page-----37d0272c52f3--------------------------------
[source]: https://blog.realkinetic.com/building-minimal-docker-containers-for-python-applications-37d0272c52f3

![docker](https://miro.medium.com/max/4400/1*Wu-YcRkS4CQUFQcY_rcdtw.png)

在创建 Docker 容器时，最佳实践时让镜像尽可能的小。分流到网络或存储到磁盘越小越好。保持小尺寸通常意味着更快的构建镜像和部署容器。

每个容器都应包括应用代码，基于语言的依赖，操作系统依赖等。再多就意味着浪费以及潜在的安全性问题。当容器中包含像 gcc 这样的工具被部署到生产时，那么攻击者使用 shell 可以很容易的构建工具来访问其他的内部系统。拥有安全层可以最大程度的减少攻击者可能引起的破坏。

## Docker 中的 Python

我最近在一个 Python 的 web 项目中工作。requirements.txt 如下: 

```
Flask>=1.1.1,<1.2
flask-restplus>=0.13,<0.14
Flask-SSLify>=0.1.5,<0.2
Flask-Admin>=1.5.3,<1.6
gunicorn>=19,<20
```
## 臃肿的镜像

要是你通过 Google 搜索 Dockerfile 的文件示例的话，应该是这样:

```
FROM python:3.7
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["gunicorn", "-w 4", "main:app"]
```

这个容器镜像可是 **958MB!!**

如果你想我一样的话，就会挠头在想“这只是个 简单的 Python web 程序，为啥那么大？？” 我们看看怎么瘦身。 

## Alpine

少即是多，但太少也可能是有害的。你可以从 [scratch][scratch] 构建所有的容器，但这意味着你必须处理像 shell 这样的低级操作系统原语，如 cat、find 等。这相当的无聊且会分散你想尽快将代码呈现给客户的注意力(我们的诅咒之一)。我发现一个实用的平衡点是使用像  [Alpine][Alpine] 这样的基础镜像。在写这篇文章时，最新的 Alpine 镜像(v3.10) 只有 5.58MB，值得尊敬。同时你也会拥有一个最小的 POSIX 环境来构建应用。

[scratch]: https://docs.docker.com/develop/develop-images/baseimages/#create-a-simple-parent-image-using-scratch
[Alpine]: https://hub.docker.com/_/alpine/

```
FROM python:3.7-alpine
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["gunicorn", "-w 4", "main:app"]
```

构建这个容器会生成 **139MB** 大小的镜像。当然基础镜像只有98.7MB(在写这篇文章时)。这意味着我们的应用占用了另外 40.3MB。

注意， Alpine 默认使用 musl 而不是 glibc。这说明一些 Python wheel 在非强制性重新编译的情况下是无法安装的。

## 开发

那些眼光敏锐以及有 Docker 经验的人会看到上面 Dockerfile 的问题。每当修改源代码然后重新构建容器时，依赖会被重新下载并安装。这可不是什么好事 -- 迭代开发需要太多时间。我们利用层缓存来重写 Dockerfile。


## 层缓存(layer caching)

```
FROM python:3.7-alpine
COPY requirements.txt /
RUN pip install -r /requirements.txt
COPY src/ /app
WORKDIR /app
CMD ["gunicorn", "-w 4", "main:app"]
```

像这样重写 Dockerfile 可以使用 Docker 的层缓存并且可以在 requiremewnts.txt 文件没有修改的情况下跳过安装 Python 依赖。

这让构建更快，但对镜像大小并没有影响。


## 缓存依赖

如何仔细看上面 Docker 的构建输出，应该是下面这样的：

```
Building wheels for collected packages: Flask-SSLify, Flask-Admin, MarkupSafe, pyrsistent
  Building wheel for Flask-SSLify (setup.py): started
  Building wheel for Flask-SSLify (setup.py): finished with status 'done'
  Created wheel for Flask-SSLify: filename=Flask_SSLify-0.1.5-cp37-none-any.whl size=2439 sha256=97d9f3687a0ead6056c0d5472e506cf01c5bbfa7e688964c96072653aa581ede
  Stored in directory: /root/.cache/pip/wheels/f6/be/7c/b262753258e34b3f07ec47973038f199c34678985b9614a50d
  Building wheel for Flask-Admin (setup.py): started
  Building wheel for Flask-Admin (setup.py): finished with status 'done'
  Created wheel for Flask-Admin: filename=Flask_Admin-1.5.3-cp37-none-any.whl size=1853777 sha256=04068d272f06c802ff0288fa5727ea01a90f5a8ce8d9f545f945c9e24207fc31
  Stored in directory: /root/.cache/pip/wheels/6f/ca/26/3dcc4b3286ed103ef9328b856221a9881188653c5d38ac73db
  Building wheel for MarkupSafe (setup.py): started
  Building wheel for MarkupSafe (setup.py): finished with status 'done'
  Created wheel for MarkupSafe: filename=MarkupSafe-1.1.1-cp37-none-any.whl size=12629 sha256=32853345d5291f8c97218a4ca0474098a69680961306192205366a277fc1141e
  Stored in directory: /root/.cache/pip/wheels/f2/aa/04/0edf07a1b8a5f5f1aed7580fffb69ce8972edc16a505916a77
  Building wheel for pyrsistent (setup.py): started
  Building wheel for pyrsistent (setup.py): finished with status 'done'
  Created wheel for pyrsistent: filename=pyrsistent-0.15.4-cp37-cp37m-linux_x86_64.whl size=56384 sha256=c960e45578b3a33a35094111af9445ebb96287038841438c83e01c9cc1df63d4
  Stored in directory: /root/.cache/pip/wheels/bb/46/00/6d471ef0b813e3621f0abe6cb723c20d529d39a061de3f7c51
Successfully built Flask-SSLify Flask-Admin MarkupSafe pyrsistent
```

当 pip install 运行时，其也复制一份下载好的依赖文件存储到 /root/.cache 下面。当在 Docker 外面做本地化开发时，这非常有用，但却使用了应用无法触及的非必需空间。该目录在 40.3MB 的应用镜像中占了 8.3MB。接下来我们通过 Docker 的另一个特性--多级构建来消除这种情况。

## 多级构建

Docker 17.05 添加了[multistage builds][multistage-build](多级构建) 的支持。这意味着依赖可以被构建进一个镜像然后导入到另一个镜像中。现在让我们重新编写 Dockerfile 文件来使用多级构建：

[multistage-build]: https://docs.docker.com/develop/develop-images/multistage-build

```
FROM python:3.7-alpine as base
FROM base as builder
RUN mkdir /install
WORKDIR /install
COPY requirements.txt /requirements.txt
RUN pip install --install-option="--prefix=/install" -r /requirements.txt
FROM base
COPY --from=builder /install /usr/local
COPY src /app
WORKDIR /app
CMD ["gunicorn", "-w 4", "main:app"]
```

该 Docker 容器只有 125MB， 编译后的 Python 依赖只占 31.0M（我在构建的容器里面运行 du -h /install 获取到这个结果）。

相较于刚开始的 958MB 来说，12MB 是个重大提升。


## 陷进无底洞

进一步减少镜像大小也是可能的 -- 通过切换 Python 的 Alpine 版本然后删除容器构建时的额外文件，如 docs、tests等。我能让镜像大小减少到 70MB 以下。但是，这值得么？如果你每天在同一组虚拟机上进行多次部署，构成镜像的大多数层很有可能会缓存到磁盘上，这意味着随着镜像大小接近于零，回报也会逐级递减。而且，这导致 Dockerfile 变得复杂而非简单就是好的（简单才是王道）。实用主义重要--让其他人来维护基本镜像可以让你专注于业务问题。

## 总结

Docker 是个强大的工具，它可以让我们的应用和语言以及操作系统依赖进行捆绑。当我们把该镜像推上生产时，这是非常有价值的，因为能保证我们测试过的镜像就是生产上运行的镜像。

如果我们天真的写 Dockerfile 文件的话， Docker 的构建系统能让创建的镜像非常大，反过来，如果正确的写的话，就会变得非常的小，轻量并且可缓存。

[Real Kinetic][Real-Kinetic]  可以帮助公司获得容器的最大价值，同时提升云架构。[联系我们][contact-us]了解更多。


[Real-Kinetic]: https://realkinetic.com/
[contact-us]: https://www.realkinetic.com/#contact