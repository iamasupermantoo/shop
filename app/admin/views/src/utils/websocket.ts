class WebsocketFunc {
  isConnect: boolean = false  //  是否连接
  url: string;                    //  连接路由
  conn: null | any = null;        //  连接对象
  restTime = 5000;      //  重连时间
  onMessageFuncList = [] as any //  消息回调
  onMessageOpenFunc: () => void = (() => {/**/})  //  消息打开事件

  // 构造函数
  constructor(url: string) {
    if (url == '' || url == undefined) {
      const protocol = document.location.protocol === 'https:' ? 'wss:' : 'ws:';
      const domain = process.env.baseURL?.substring(process.env.baseURL.indexOf('//'))
      this.url = protocol + domain + '/ws';
    } else {
      this.url = url
    }
  }

  connect() {
    if (this.conn != null) {
      return;
    }

    // 连接websocket
    this.conn = new WebSocket(this.url);

    // 关闭websocket
    this.conn.onclose = () => {
      this.isConnect = false
      this.conn = null;

      // 等待时间进行重连
      setTimeout(() => {
        this.connect();
      }, this.restTime);
    };

    // 连接成功, 绑定
    this.conn.onopen = () => {
      this.isConnect = true
      this.onMessageOpenFunc()
    }

    // 消息回调
    this.conn.onmessage = (ev: any) => {
      this.onMessageFuncList.forEach((callbackFunc: any) => {
        callbackFunc(JSON.parse(ev.data));
      });
    };
    return this
  }

  // 发送Json消息
  sendMessageJsonFunc(data: object) {
    if (this.conn == null || !this.isConnect) {
      setTimeout(() => {
        this.sendMessageJsonFunc(data)
      }, 1000)
      return
    }

    this.conn.send(JSON.stringify(data));
  }

  // 发送Text消息
  sendMessageTextFunc(data: string) {
    if (this.conn) {
      this.conn.send(data)
    }
  }

  // 设置消息回调
  setOnMessageFunc(fn: (msg: any) => void) {
    this.onMessageFuncList.push(fn);
    return this
  }

  // 设置连接成功事件
  setOnMessageOpenFunc(fn: () => void) {
    this.onMessageOpenFunc = fn
    return this
  }
}

export default WebsocketFunc;
