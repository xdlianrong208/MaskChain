#### 注册

- 请求路径与方式

  ​	/register	post

- 所需参数

  ```
  Name string `json:"name" form:"name"`
  Id      string `json:"id" form:"id"`
  Str  string `json:"str" form:"str"`
  ```

- 返回

  ​	Pub  PublicKey  `json:"Pub"`



#### 购币

- 请求路径与方式

  ​	/buycoin	post

- 所需参数

  ​	

  ```
  G1     string `json:"g1"`
  G2     string `json:"g2"`
  P      string `json:"p"`
  H      string `json:"h"`
  Account struct {
     Pub  PublicKey  `json:"Pub"`
     Priv PrivateKey `json:"Priv"`
     Info struct {
        Name    string `json:"Name"`
        ID      string `json:"ID"`
        Hashky  string `json:"Hashky"`
        ExtInfo string `json:"ExtInfo"`
     } `json:"Info"`
  }
  ```

- 返回

  ```
  Cmv    string `json:"cmv"`
  Epkrc1 string `json:"epkrc1"`
  Epkrc2 string `json:"epkrc2"`
  Hash   string `json:"hash"` //此次购币交易的交易哈希
  ```



#### 转账

- 请求路径与方式

  ​		/exchangecoin	post

- 所需参数

  ```
  Cmv    string `json:"cmv"`
  Epkrc1 string `json:"epkrc1"`
  Epkrc2 string `json:"epkrc2"`
  Hash   string `json:"hash"` 
  Priv PrivateKey `json:"Priv"`
  ```

- 返回

  ```
  Cmv  string `json:"cmv"`
  Vor  string `json:"vor"`
  Hash string `json:"hash"`
  ```



