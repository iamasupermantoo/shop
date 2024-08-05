import WebsocketFunc from 'src/utils/websocket';

export const MessageOperateBindUser = 'bindUser'; //  绑定用户
export const MessageOperateAudio = 'audio'  //  提示音消息

export let appWebsocket = null as any
export const initWebsocket = () => {
  appWebsocket = new WebsocketFunc('')
  return appWebsocket
}
