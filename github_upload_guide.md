# Go-Aibote GitHub上传指南

## 准备工作

在开始上传之前，请确保您已经完成以下准备工作：

1. 您已经拥有一个GitHub账号
2. 您的本地环境已经安装了Git
3. 您已经完成了项目中的所有Python相关元素的清理（本指南中的项目已完成清理）

## 步骤1：初始化Git仓库（如果尚未初始化）

打开命令行工具（如PowerShell、Git Bash等），进入go-aibote项目目录：

```bash
cd d:\aibote\PyAibote\go-aibote
```

执行以下命令初始化Git仓库：

```bash
git init
```

## 步骤2：配置Git用户信息

设置您的用户名和电子邮件地址（这些信息将显示在您的GitHub提交记录中）：

```bash
git config --global user.name "zhangsan-ai"
git config --global user.email "3258856837@qq.com"
```

## 步骤3：添加和提交文件

将所有文件添加到暂存区：

```bash
git add .
```

提交文件到本地仓库：

```bash
git commit -m "Initial commit: Go-Aibote RPA framework"
```

## 步骤4：在GitHub上创建新仓库

1. 登录您的GitHub账号
2. 点击页面右上角的"+"号，选择"New repository"
3. 在"Repository name"字段中输入"go-aibote"
4. 选择仓库可见性（公开或私有）
5. 不要勾选"Initialize this repository with a README"、"Add .gitignore"或"Add a license"选项，因为我们已经在本地仓库中设置了这些内容
6. 点击"Create repository"

## 步骤5：设置远程仓库地址

在GitHub上创建仓库后，您会看到一个页面，其中包含了仓库的URL。复制这个URL，然后在命令行中执行以下命令：

```bash
git remote add origin https://github.com/zhangsan-ai/go-aibote.git
```

## 步骤6：推送代码到GitHub

执行以下命令将本地仓库的代码推送到GitHub：

```bash
git push -u origin master
```

或者，如果您使用的是main分支作为默认分支（GitHub现在默认使用main分支）：

```bash
git push -u origin main
```

## 步骤7：更新模块路径（重要）

在将代码推送到GitHub后，您需要更新项目中的模块路径，使其与您的实际GitHub用户名匹配：

1. 编辑go.mod文件，将
   ```
   module github.com/example/go-aibote
   ```
   改为
   ```
   module github.com/zhangsan-ai/go-aibote
   ```

2. 更新所有代码文件中的import路径，将
   ```go
   import "github.com/example/go-aibote/..."
   ```
   改为
   ```go
   import "github.com/zhangsan-ai/go-aibote/..."
   ```

3. 提交并推送这些更改：
   ```bash
   git add .
   git commit -m "Update module path to match GitHub username"
   git push
   ```

## 步骤8：验证上传结果

打开您的GitHub仓库页面，检查所有文件是否已成功上传。您应该能够看到：
- README.md文件
- go.mod文件
- .gitignore文件
- pkg目录及其所有内容
- cmd目录及其所有内容

## 额外提示

1. **添加许可证**：如果您想为项目添加许可证，请在GitHub仓库页面的"Settings" > "Licenses"中选择一个合适的许可证。

2. **启用GitHub Actions**：如果您想为项目添加自动化测试或持续集成，请在GitHub仓库页面的"Actions"标签中设置。

3. **创建Releases**：当您的项目达到一定的里程碑时，可以创建Releases来标记不同的版本。

4. **邀请协作者**：如果您想让其他人参与项目开发，可以在GitHub仓库页面的"Settings" > "Manage access"中邀请协作者。

恭喜您！您已成功将Go-Aibote项目上传到GitHub。现在您可以与他人共享这个项目，或者继续开发和完善它。