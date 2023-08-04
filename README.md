# blob

## 项目介绍

本项目来源于知了传课：<https://www.bilibili.com/video/BV1YV411b7JJ/?spm_id_from=333.999.0.0&vd_source=84fc27804252448ba51ef3b6abfd5d36>

课程资料和源码均已上传Github。

笔者对源代码有不少修改,比如代码的优化更新。通过学习这个项目，你可以对用beego进行全栈开发有一定了解。

后台由管理员登录，主要负责上传修改删除帖子。

前台由用户登录，主要负责查看帖子，以及评论。本项目支持二级评论。

主要功能如下：

![image](https://github.com/uestc-wxy/blob/assets/120303802/d04b4a62-8ed7-4afd-90f9-954065fc639b)

![image](https://github.com/uestc-wxy/blob/assets/120303802/da6fd30a-7334-4956-8e08-186f7d135d83)

![image](https://github.com/uestc-wxy/blob/assets/120303802/1773ad4e-64d4-4577-b8be-e7403b96c6fe)

![image](https://github.com/uestc-wxy/blob/assets/120303802/e2e4192f-dc50-441a-98e1-69f8f5c7c056)

![image](https://github.com/uestc-wxy/blob/assets/120303802/8d2176a5-dfbe-4e46-9a6a-29054986eaef)

## 关于使用

克隆仓库到本地：
```bash
git clone https://github.com/uestc-wxy/blob.git
```

配置`mysql.conf`文件：
> username = yourusername
> 
> password = yourpwd
> 
> host = 127.0.0.1
> 
> port = 3306
> 
> database = blob

在本地创建名为`blob`的数据库，然后在项目根目录运行`bee run`，前提是你有`bee`工具。当然也可以运行：
```bash
go build
./blob
```

后台管理登录界面：http://127.0.0.1:8080/cms

前端界面：http://127.0.0.1:8080

由于后端管理不支持管理员注册，需要自行在数据库插入一条管理员数据，注意密码使用MD5加密，请将加密后的密码填入数据库中，加密函数在tests/test里~












