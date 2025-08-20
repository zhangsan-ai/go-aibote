# Go-Aibote模块路径更新指南

## 重要提示

在将代码推送到GitHub后，您**必须**更新项目中的模块路径，使其与您的实际GitHub用户名匹配。这是确保Go模块系统正常工作的关键步骤。

## 更新步骤

请按照以下步骤操作，将占位符模块路径更新为您的实际GitHub用户名：

### 步骤1：编辑go.mod文件

1. 打开`d:\aibote\PyAibote\go-aibote\go.mod`文件
2. 找到第一行：
   ```
   module github.com/example/go-aibote
   ```
3. 将其替换为：
   ```
   module github.com/zhangsan-ai/go-aibote
   ```

### 步骤2：更新所有代码文件中的import路径

您需要更新以下文件中的import路径：

#### 1. pkg/windowsbot/windowsbot.go
- 找到：
  ```go
  import "github.com/example/go-aibote/pkg/common"
  ```
- 替换为：
  ```go
  import "github.com/zhangsan-ai/go-aibote/pkg/common"
  ```

#### 2. pkg/webbot/webbot.go
- 找到：
  ```go
  import "github.com/example/go-aibote/pkg/common"
  ```
- 替换为：
  ```go
  import "github.com/zhangsan-ai/go-aibote/pkg/common"
  ```

#### 3. pkg/androidbot/androidbot.go
- 找到：
  ```go
  import "github.com/example/go-aibote/pkg/common"
  ```
- 替换为：
  ```go
  import "github.com/zhangsan-ai/go-aibote/pkg/common"
  ```

#### 4. cmd/windows-bot/main.go
- 找到：
  ```go
  import "github.com/example/go-aibote/pkg/windowsbot"
  ```
- 替换为：
  ```go
  import "github.com/zhangsan-ai/go-aibote/pkg/windowsbot"
  ```

### 步骤3：提交并推送更改

完成上述修改后，执行以下命令提交并推送更改：

```bash
cd d:\aibote\PyAibote\go-aibote
git add .
git commit -m "Update module path to match GitHub username"
git push
```

## 验证

更新完成后，您可以通过运行以下命令验证更改是否生效：

```bash
go mod tidy
go build ./cmd/windows-bot
```

如果这些命令成功执行且没有报错，则说明模块路径更新成功。

## 注意事项

- 请确保将所有文件中的`github.com/example/go-aibote`替换为`github.com/zhangsan-ai/go-aibote`
- 模块路径必须与您在GitHub上创建的仓库路径完全匹配
- 如果您在将来更改了GitHub用户名，需要再次执行类似的步骤更新模块路径

完成这些步骤后，其他开发者就可以通过以下命令轻松地使用您的库：

```bash
go get github.com/zhangsan-ai/go-aibote
```