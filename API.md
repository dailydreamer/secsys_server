# 概述

## 基础URL（暂时未定）

应将基础URL作为一个config参数，方便日后迁移。

```
https://fakeurl/v1
```

## HTTP动词

| **Verb** | **Description** |
|---|---|
| GET | 获取资源 |
| POST | 新建资源 |
| PATCH | 更新资源（提供部分参数） |
| PUT | 更新资源（提供全部参数） |
| DELETE | 删除资源 |

Request Body和Response Body均为JSON格式。

## 状态码

| **Status Code** | **Description** |
|---|---|
| 200 OK | 请求成功 |
| 400 Bad Request | 解析JSON格式出错 |
| 422 Unprocessable Entity | 提供资源参数不合要求 |
| 401 Unauthorized | 用户未登录或登录错误 |
| 403 Forbidden | 用户权限不够 |
| 404 NOT FOUND | 该资源不存在 |
| 500 INTERNAL SERVER ERROR | 服务器内部错误 |

当返回错误时会同时返回错误信息。

例子

```json
400 Bad Request
{
  "error":"Problems parsing JSON"
}
```

## 分页

如果一个资源需要分页，那么其文档中会注明。默认page为0，per_page为10。

URL参数

| **Name** | **Type** | **Description** |
|---|---|---|
| page | int | 页号 |
| per_page | int | 每页数量 |

例子

```bash
curl 'https://fakeurl/v1/companies?page=2&per_page=100'
```

Response中还会包含Link Header。`rel`指向相关页的资源，其可能值如下：

| **Name** | **Description** |
|---|---|
| next | 下一页 |
| last | 最后一页 |
| first | 第一页 |
| prev | 前一页 |

例子

```
Link: <https://fakeurl/v1/companies?page=3&per_page=100>; rel="next",
  <https://fakeurl/v1/companies?page=50&per_page=100>; rel="last"
```

## 排序

如果一个资源支持排序，那么其文档中会注明。

URL参数

| **Name** | **Type** | **Description** |
|---|---|---|
| sortby | string | 排序的属性 |
| order | ENUM string | 升序（asc）还是降序（desc） |

例子

```bash
curl 'https://fakeurl/v1/companies?sortby=com_name&order=asc'
```

## 时间

返回的时间均采用RFC3339格式，如"2017-01-01T00:00:00Z08:00"，前端需根据客户端所在时区进行转化。

# API

登录后`Header`中要携带`token`

## API目录

- `GET`用来**获取**数据
- `POST`用来**新增**数据
- `PUT`用来**修改**数据
- `DELETE`用来**删除**数据

### 游客可用API

- `GET    /verification`
- `POST   /login`

### 登录用户可用的API

- `GET    /users/{userID}`
- `PUT  /users/{userID}`
- `DELETE /users/{userID}`
- `GET    /users/{userID}/contracts`
- `POST   /users/{userID}/contracts`
- `GET    /users/{userID}/contracts/{contractID}`
- `PUT  /users/{userID}/contracts/{contractID}`
- `DELETE /users/{userID}/contracts/{contractID}`
- `GET    /users/{userID}/scores`

### 仅Admin可用的API

- `GET    /users`
- `POST   /users`
- `GET    /contracts`
- `POST   /contracts`
- `GET    /contracts/{contractID}`
- `PUT  /contracts/{contractID}`
- `DELETE /contracts/{contractID}`
- `GET    /scores`
- `POST   /scores`
- `GET    /scores/{scoreID}`
- `PUT  /scores/{scoreID}`
- `DELETE /scores/{scoreID}`
- `GET    /logs`
- `GET    /messages`

## 登录相关API

### 获取登录验证码（1登录页面.png）

- Request
`GET /verification`

- Response
```javascript
{
    verification: String
}
```

### 登录（1登录页面.png）

- Request
`POST /login`
```javascript
{
    phone: String,
    password: String,
    verification: String
}
```

- Response
```javascript
{
    token: String,
    id: String,
    isAdmin: Boolean,
    comName: '北京保安协会评定办' | '北京安祥致远保安服务有限公司'
}
```

## `GET`类型API

### 获取企业信息列表（2评定办首页.png）

- Request
`GET /users`

- Response
```javascript
[{
    id: String,              // 用户ID
    comName: String,         // 企业名称
    comField: String,        // 企业所在领域
    comMan: String,          // 法人/总经理
    comPhone: String,        // 联系电话
    comRegnum: String,       // 企业注册号
    comRegcap: String,       // 工商注册资本金
    comCapreport: String,    // 验资报告资本金
    comBatch: String,        // 批次
    comLicense: String       // 营业执照扫面件
}]
```

### 获取企业详细信息（9企业首页.png）

- Request
`GET /users/{userID}`

- Response
```javascript
{
    // 账户信息
    nickName: String,        // 使用人
    phone: String,           // 手机号
    email: String,           // 邮箱
    avatar: String           // 头像
    // 企业基本信息 
    id: String,              // 用户ID
    comName: String,         // 企业名称
    comField: String,        // 企业所在领域
    comMan: String,          // 法人/总经理
    comPhone: String,        // 联系电话
    comRegnum: String,       // 企业注册号
    comRegcap: String,       // 工商注册资本金
    comCapreport: String,    // 验资报告资本金
    comBatch: String,        // 批次
    // 上传的文件
    comLicense: String       // 营业执照扫面件
    // 企业等级信息
    comLevel: String,        // 企业等级
    appliDate: String,       // 申请日期
    appliLevel: String,      // 申请级别
    appliResult: String,     // 降级/通过
    certfDate: String,       // 发证日期
    certfNum: String,        // 证书号
    verifDate: String,       // 年审日期
    verifResult: String,     // 年审结果
    // 企业规模
    comTurnover: String,     // 上一年度总营业额
    comArea: String,         // 固定办公场所面积
    policeNum: String,       // 公安信息监督人数
    policeDuty: String,      // 公安监管系统正在执行的驻勤点总数
    policeCancel: String,    // 公安监管系统的撤点项目总数
    policeDutycancel: String,// 公安监管系统的驻勤点撤点率
    listDuty: String,        // 列表申报正执行的驻勤点总数
    listDutycancel: String,  // 列表申报驻勤点撤点率
    // 企业用人情况
    empNum: String,          // 聘用人数
    empContract: String,     // 签订劳动合同人数
    empLccr: String,         // 签订率
    contNum: String,         // 与甲方签订合同总人数
    contVac: String,         // 空额率
    contSamptnum: String,    // 抽样5个项目合同总人数
    contSampfnum: String,    // 抽样5个项目实际在岗总人数
    contSampvac: String,     // 抽样5个项目保安员空额率
    empSep: String,          // 上一年离职人员总人数
    empSeprate: String,      // 流失率
    // 员工职业资格
    listCertrate: String,    // 列表保安员证持证率
    listSampcertrate: String,// 抽样项目保安员持证率
    empSsemanum: String,     // 高级保安管理师人数
    empSsemarate: String,    // 高级保安管理师持证率
    empSemanum: String,      // 保安管理师人数
    empSemarate: String,     // 保安管理师持证率
    empJsenum: String,       // 初级保安员人数
    empJserate: String,      // 初级保安证持证率
    // 员工职业资格
    trainPeriod: String,     // 培训课时
    comSalary: String,       // 上一年度工资
    trainFunds: String,      // 培训经费
    trainFundsrate: String,  // 培训经费与全员工资总额比例
    // 员工职业资格
    comComins: String,       // 商业险
    comSosec: String,        // 参险员工数（社保、新农合）
    comSosecrate: String,    // 参险人员与全员比（社保、新农合）
    // 党团工会建设
    comParty: String,        // 党组织
    comYouth: String,        // 团组织
    comUnion: String,        // 工会
    // 社会责任
    comCrime: String,        // 重大刑事案件
    comAcc: String,          // 重大现任事故
    comMwgs: String,         // 保安员工资高于最低工资标准10%
}
```

### 获取合同列表（3评定办合同列表.png & 10企业合同列表.png）

- Request
`GET /contracts` or `GET /users/{userID}/contracts`

- Response
```javascript
[{
    id: String,
    userID: String,
    comName: String,        // 公司名称
    contractNo: String,     // 合同编号
    projectName: String,    // 项目名称
    comField: String,       // 所有项目？（人力防范RF等）
    customerName: String,   // 客户名称
    customerType: String,   // 客户类别
    peopleNum: String,      // 合同人数
    startTime: String,      // 起始时间
    endTime: String,        // 终止时间
    unitPrice: String,      // 合同单价
    totalPrice: String,     // 合同总额
    income: String          // 入账金额（分年度统计，从合同起始年份至结束年份）
}]
```

### 获取打分表（4评定办打分表.png & 11企业评级表页面.png）

- Request
`GET /scores` or `GET /users/{userID}/scores`

- Response
```javascript
[{
    id: String,
    userID: String,
    comName: String,        // 公司名称
    year: String,           // 打分年份
    standard: String,       // 打分标准：全国标准/北京标准
    scoreNo: String,        // 序号
    scoreType: String,      // 打分类别
    satisfied: String,      // 满足评分项
    score: String,          // 得分
    reason: String          // 结果成因
}]
```

### 获取评定办账户登录日志(6评定办登录日志.png & 13企业登录日志.png)
- Request
`GET /logs`

- Response
```javascript
[{
    id: String,
    userID: String,
    created: String,  // 登录日期时间
    comName: String,  // 操作公司
    ip: String,       // 登录ip
    address: String,  // 登陆地点
    status: String    // 是否异常
}]
```

### 获取评定办账户消息(7评定办新消息处理页面.png)
- Request
`GET /messages`

- Response
```javascript
[{
    id: String,
    userID: String,
    created: String,  // 操作时间
    comName: String,  // 操作公司
    message: String   // 操作内容
}]
```

## `POST`类型API

用来**新增**数据，参数与对应的`GET`类型API相同。这里只列出URL列表和其中一个例子：

- `POST /users`

| **Name** | **Type** |
|---|---|
| id | AutoGenerate |
| phone | Required |
| password | Required |
| comName | Required |

- `POST /contracts` or `POST /users/{userID}/contracts`

| **Name** | **Type** |
|---|---|
| id | AutoGenerate |
| userID | AutoGenerate |
| comName | Required |
| startTime | Required |
| endTime | Required |

- `POST /scores`

| **Name** | **Type** |
|---|---|
| id | AutoGenerate |
| userID | AutoGenerate |
| comName | Required |

### 例：新增企业信息（8新增企业信息.png）

- Request
`POST /users`
```javascript
{
    // 账户信息
    nickName: String,        // 使用人
    phone: String,           // 手机号
    email: String,           // 邮箱
    avatar: String           // 头像
    // 企业基本信息
    id: String,              // 用户ID
    comName: String,         // 企业名称
    comField: String,        // 企业所在领域
    comMan: String,          // 法人/总经理
    comPhone: String,        // 联系电话ID
    comRegnum: String,       // 企业注册号
    comRegcap: String,       // 工商注册资本金
    comCapreport: String,    // 验资报告资本金
    comBatch: String,        // 批次
    // 上传的文件
    comLicense: String       // 营业执照扫面件
    // 企业等级信息
    comLevel: String,        // 企业等级
    appliDate: String,       // 申请日期
    appliLevel: String,      // 申请级别
    appliResult: String,     // 降级/通过
    certfDate: String,       // 发证日期
    certfNum: String,        // 证书号
    verifDate: String,       // 年审日期
    verifResult: String,     // 年审结果
    // 企业规模
    comTurnover: String,     // 上一年度总营业额
    comArea: String,         // 固定办公场所面积
    policeNum: String,       // 公安信息监督人数
    policeDuty: String,      // 公安监管系统正在执行的驻勤点总数
    policeCancel: String,    // 公安监管系统的撤点项目总数
    policeDutycancel: String,// 公安监管系统的驻勤点撤点率
    listDuty: String,        // 列表申报正执行的驻勤点总数
    listDutycancel: String,  // 列表申报驻勤点撤点率
    // 企业用人情况
    empNum: String,          // 聘用人数
    empContract: String,     // 签订劳动合同人数
    empLccr: String,         // 签订率
    contNum: String,         // 与甲方签订合同总人数
    contVac: String,         // 空额率
    contSamptnum: String,    // 抽样5个项目合同总人数
    contSampfnum: String,    // 抽样5个项目实际在岗总人数
    contSampvac: String,     // 抽样5个项目保安员空额率
    empSep: String,          // 上一年离职人员总人数
    empSeprate: String,      // 流失率
    // 员工职业资格
    listCertrate: String,    // 列表保安员证持证率
    listSampcertrate: String,// 抽样项目保安员持证率
    empSsemanum: String,     // 高级保安管理师人数
    empSsemarate: String,    // 高级保安管理师持证率
    empSemanum: String,      // 保安管理师人数
    empSemarate: String,     // 保安管理师持证率
    empJsenum: String,       // 初级保安员人数
    empJserate: String,      // 初级保安证持证率
    // 员工职业资格
    trainPeriod: String,     // 培训课时
    comSalary: String,       // 上一年度工资
    trainFunds: String,      // 培训经费
    trainFundsrate: String,  // 培训经费与全员工资总额比例
    // 员工职业资格
    comComins: String,       // 商业险
    comSosec: String,        // 参险员工数（社保、新农合）
    comSosecrate: String,    // 参险人员与全员比（社保、新农合）
    // 党团工会建设
    comParty: String,        // 党组织
    comYouth: String,        // 团组织
    comUnion: String,        // 工会
    // 社会责任
    comCrime: String,        // 重大刑事案件
    comAcc: String,          // 重大现任事故
    comMwgs: String,         // 保安员工资高于最低工资标准10%
}
```

- Response
```javascript
    success: true
```

## `PUT`类型API

用来**修改**数据，参数与对应的`GET`类型API相同，这里只列出URL列表：
- `PUT /users/{userID}`
- `PUT /contracts/{contractID}` or `PUT /users/{userID}/contracts/{contractID}`
- `PUT /scores/{scoreID}`

## `DELETE`类型API

用来**删除**数据
- `DELETE /users/{userID}`
- `DELETE /contracts/{contractID}` or `DELETE /users/{userID}/contracts/{contractID}`
- `DELETE /scores/{scoreID}`
