# protoc-gen-go-errors-code

`protoc-gen-go-errors-code` 是一个 Protobuf 插件，用于生成 Go 语言的错误码定义文件。

## 安装

```sh
go install github.com/costa92/k8s-krm-go/tools/protoc-gen-go-errors-code
```

## 使用
    `path/to/your/protos` 是定义 protoc 插件的 proto 文件的目录，`path/to/your/` 是生成的 Go 语言错误码定义文件的目录。
    `path/to/your/docs` 是生成的错误码文档的目录。

    ```sh
    protoc --go-errors-code_out=paths=source_relative:. path/to/your/protos \
           --go-errors-code_out=paths=source_relative:. path/to/your/docs
    ```
## example

    ```sh
    protoc 	--go-errors_out=paths=source_relative:$(APIROOT) \
            --go-errors-code_out=paths=source_relative:$(KRM_ROOT)/docs/guide/zh-CN/api/errors-code
    ```

## 参考
    https://github.com/lyouthzzz/protoc-gen-go-errors