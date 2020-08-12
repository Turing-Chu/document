# 为 Fork 的项目配置远程

必须在 git 中配置一个指向上游仓库的远程，以将[在 fork 中所做的更改][ref1][1]与原始存储库同步。 这还允许使用 fork 同步在原始存储库中所做的更改。

[ref1]:https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork

## MAC

- 1. 打开终端。
- 2. 列出所 fork 的远程仓库的当前配置

```
$ git remote -v
> origin	https://github.com/YourUserName/YourFork.git (fetch)
> origin	https://github.com/YourUserName/YourFork.git (push)
```

- 3. 指定一个新的上游远程仓库，其会与fork进行同步。

```
$ git remote add upstream https://github.com/Original_Owner/Original_Repository.
```

- 4. 验证你所指定的 fork 的新上游仓库。

```
$ origin	      https://github.com/YourUserName/YourFork.git (fetch)
> origin	      https://github.com/T YourUserName/YourFork.git (push)
> upstream	https://github.com/Original_Owner/Original_Repository (fetch)
> upstream	https://github.com/Original_Owner/Original_Repository (push)
```

# 同步一个 fork

同步一个 repository 的 fork 以与上游 repository 保持最新。

在将 fork 与上游 repository  同步之前，必须在 git 中[配置远程指向上游 repository][ref2][2]。

[ref2]: https://docs.github.com/en/enterprise/2.15/user/articles/configuring-a-remote-for-a-fork


## MAC

- 1. 打开终端。
- 2. 切换到当前本地项目的工作目录。
- 3. 从上游 repository 拉取(fetch) 分支与各自 commit。`master` 上的 commit 会存储到本地 `upstream/master` 分支中。

```
$ git fetch upstream
> remote: Enumerating objects: 166, done.
> remote: Counting objects: 100% (166/166), done.
> remote: Compressing objects: 100% (30/30), done.
> remote: Total 314 (delta 133), reused 153 (delta 129), pack-reused 148
> Receiving objects: 100% (314/314), 139.21 KiB | 30.00 KiB/s, done.
> Resolving deltas: 100% (183/183), completed with 57 local objects.
> From https://github.com/Original_Owner/Original_Repository
> * [new branch]      BRANCH-1 -> upstream/BRANCH-1
> * [new branch]      BRANCH-2 -> upstream/BRANCH-2
> * [new branch]      master         -> upstream/master
> * [new tag]           TAG-1         -> TAG-1
> * [new tag]           TAG-2         -> TAG-2
```

- 4. checkout 你所 fork 的本地 `master` branch

```
$ git checkout master
> Switched to branch 'master'
``` 

- 5. 将 `upstream/master` merge 到你本地的 `master` branch 上. 这会将上游 repository 同步到你  frok 的 `master` branch 上，且没有丢失本地修改。

```
$ git merge upstream/master
> Updating c0c2ff8..b28321d
> Fast-forward
>  .github/workflows/ci.yml                     |    4 +-
>  .travis.yml                                          |    2 -
>  README.md                                      |    8 +-
```

如果本地分支没有其他特殊的提交，git 会执行 `fast-forward`:

```
$ git merge upstream/master
> Updating c0c2ff8..b28321d
> Fast-forward
>  .github/workflows/ci.yml                     |    4 +-
>  .travis.yml                                          |    2 -
>  README.md                                      |    8 +-
```

> 同步(sync) 只更新你所 fork 的 repository 的本地副本。如果要将你的 fork 更新到 GitHub 企业服务器实例上，必须[推送你的变更][refc3][3]。

[refc3]: https://docs.github.com/en/enterprise/2.15/user/articles/pushing-commits-to-a-remote-repository

[1]\: https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork

[2]\: https://docs.github.com/en/enterprise/2.15/user/articles/configuring-a-remote-for-a-fork

[3]\: https://docs.github.com/en/enterprise/2.15/user/articles/pushing-commits-to-a-remote-repository
