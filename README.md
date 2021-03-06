## 区块链

- 去中心化
- 可追溯
- 不可篡改


### 知识点

- 区块与区块链
- 共识
- 数据库
- CLI命令行操作
- 交易管理
- 密码学（crypto）
- 数字签名（ecdsa）
- 交易缓存池
- P2P网络管理



- 基本概念
  - 比特币系统没有余额，使用UTXO（unspent transaction output）模型
  - UTXO: 未花费输出
  - 锁定脚本可以理解为**用户名**
  - 在比特币钱包中的余额，其实是一个钱包地址的UTXO
  - UTXO可以理解为一个b,包含金额与拥有者
  - 有效交易条件
    - 交易需要签名：来自UTXO的拥有者
  - 交易过程解析
    - ![image.png](./img/transaction.jpg)
    - Alice接受12.5b奖励
    - 锁定脚本作用：只有alice用这笔钱
    - ALice向Bob进行转账
      - 首先确认alice有足够的钱（将所有的UTXO汇合起来）
      - 创建一笔交易（转账交易）
        - 交易中包含两个输出
          - 0发送给bob
          - 1发给给alice
          - 完成交易之后，还剩下一个UTXO

      - 在一个交易中，如果指定地址的某个未话费输出utxo已经被其他交易的input所引用。那这一个输出，就不能使用第二次了。
      - ![image.png](./img/utxo.png)
      - ```
          Height : 2 
         TimeStamp : 1644913486 
         PrevBlockHash : 00000e4e792fc3b7a50f3d3c86c3cb32229f9d52da91bb7d0150fb2b8f42c9b9 
         Hash : 00000ac3ee00bac7de2bb0e2feadb49400879a0935161afc1bb9c203a47fcaa3 
         Transaction : [c0000d6050] 
                 tx-hash : 40956b11157f060b80a227452a89eaef6f4f2a2f7b75e9ffa833dce56d454a19 
                 输入..
                        vin-txhash: [53 49 53 98 51 102 56 52 53 101 50 100 54 100 51 50 57 52 102 52 52 101 98 54 54 56 101 101 101 102 51 49 99 100 56 55 53 98 48 97 57 102 49 49 53 97 54 99 55 55 102 48 48 101 101 97 51 49 99 49 49 102 50 51] 
                        vin-vout: 0 
                        vin-scripsig: lijia 
                 输出..
                        vout-value: 4 
                        vout-ScriptPubkey: darren 
                        vout-value: 6 
                        vout-ScriptPubkey: lijia 
         Nonce : 505688 

         Height : 1 
         TimeStamp : 1644913419 
         PrevBlockHash :  
         Hash : 00000e4e792fc3b7a50f3d3c86c3cb32229f9d52da91bb7d0150fb2b8f42c9b9 
         Transaction : [c0000d60a0] 
                 tx-hash : 515b3f845e2d6d3294f44eb668eeef31cd875b0a9f115a6c77f00eea31c11f23 
                 输入..
                        vin-txhash: [] 
                        vin-vout: -1 
                        vin-scripsig: Genesis Data 
                 输出..
                        vout-value: 10 
                        vout-ScriptPubkey: lijia 
         Nonce : 1107723 


        ```
      

- 地址
  - ![image.png](./img/gongyaomakebtcaddress.jpg)
  - 地址构造
    - version
    - public key hash
    - checksum
  -  ![image.png](./img/makeaddress.png)


### 已经实现

- 区块基本结构与区块添加
- 区块链基本结构的实现
- pow共识算法
- 数据库实现
- 区块数据的持久化
- 数据迭代
- 实现命令行操作
- 当前命令
  - 创建区块链
  - 添加区块链
  - 打印区块链
- 实现获取区块链对象
- 修改区块结构，并且替换
- 输入输出UTXO
- 挖矿交易
  - MineNewBlock() 挖矿
- 普通转账交易
- 实现CLI查询余额与UTXO函数定义
- UTXO优化查找功能
- 文件分离
- UTXO优化
- 多笔交易
  - 考虑未打包到区块中数据也的考虑到,如果只查询区块中的数据，会出现余额不足
- 实现sha256he ripedmd160 实例问题
- base64
  - 主要用于传输字节码，是一个从二进制到字符的过程，通常可以在http中传递比较昶的表示信息
  - base64说明：是将3个8位字节转化为4个6位字节，在6位前面补两个0
  - base64中，输出0字符使用"="
  - ![image.png](./img/base64编码表.png) 
  - 有一个专门的编码表，以便统一进行转换的
- base58
  - 概念：和base64功能相同，但相对来说，去掉了容易产生混淆的字符
  - 去掉的字符  "+" ，"/"  "0" "O" "l" "1"
- 钱包
  - 比特币钱包本质就是一个公钥-私钥的秘钥对
    - 通过钱包获取地址
  - 过程
    - 1 公钥进行sha256再进ripemd160得到公钥hash
    - 2 组成
      - 1 version 版本前缀，大小是一个字节，用于创造一个易于辨别的格式 "1" 代表比特币的地址
      - ![image.png](./img/addressversion.png)
      - pubkey hash 20个字节 公钥哈希
      - checksum: 校验和，4个字节，添加到正在编码的数据的一端的4个字节检验和是通过pubkey哈希得到，用来检测输入时产生的错误
      - ![image.png](./img/公钥生成比特币地址.jpg)
    - 3 (version + pubkey hash +checkSum) 得到两次hash的地址
    - 4 通过base58编码得到比特币的地址
  - 添加命令行对钱包的操作功能
    - 创建钱包
    - 获取当前钱包集合
    - 将钱包地址持久化
    - 钱包与UTXO结合 


### 将钱包功能嵌入到区块链中

- 









 


