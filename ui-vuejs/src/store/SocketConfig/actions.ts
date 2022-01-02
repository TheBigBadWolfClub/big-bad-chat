import {ActionContext, ActionTree} from 'vuex';
import {StateInterface} from 'src/store';
import {Message, MsgType, MsgUpdateRooms} from 'src/store/SocketConfig/message';
import {ChatRoomType, ChatServer, ChatServerId} from 'src/store/ChatServer/model';
import {Notify} from 'quasar';
import {AppConfig, AppConfigId} from 'src/store/AppConfig/model';
import {MutationType, SocketConfig, SocketConfigState} from 'src/store/SocketConfig/model';
import {Mutations} from 'src/store/SocketConfig/mutations';
import {api} from 'boot/axios';
import {logger} from 'boot/logger';


export type ActionAugments = Omit<ActionContext<SocketConfigState, StateInterface>, 'commit'> & {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>
}

export type Actions = {
  [SocketConfig.Action.Connect](context: ActionAugments): void
  [SocketConfig.Action.Disconnect](context: ActionAugments): void
  [SocketConfig.Action.Send](context: ActionAugments, msg: Message): void
}

const socketConfigActions: ActionTree<SocketConfigState, StateInterface> & Actions = {

  [SocketConfig.Action.Connect]({dispatch, state, commit, rootState, rootGetters}) {

    const name = rootGetters[ChatServerId(ChatServer.Getter.User)].name
    const wsUrl = rootGetters[AppConfigId(AppConfig.Getter.WebsocketUri)]
    const conn = new WebSocket(`${wsUrl}?user=${name}&uuid=${state.uuid}`);

    const OnOpen = (event: Event) => {
      logger.debug('OnOpen', event)
      commit(MutationType.Status, undefined)
    }

    const OnClose = (event: CloseEvent) => {
      logger.debug('closed', event)
      commit(MutationType.Status, undefined)
      Notify.create('Websocket disconnected')
    }

    const OnError = (event: Event) => {
      logger.debug('OnError', event)
      commit(MutationType.Status, undefined)
    }

    const OnMessage = (event: MessageEvent) => {
      logger.debug('OnMessage', event)
      const split = event.data.split('\n');
      split.forEach((str: string) => {
        const message = JSON.parse(str);
        commit(MutationType.OnMessage, message)
        messageHandler(state, message, commit, rootGetters)
      })

    }

    conn.addEventListener('open', OnOpen)
    conn.addEventListener('close', OnClose)
    conn.addEventListener('error', OnError)
    conn.addEventListener('message', OnMessage)

    commit(MutationType.Connect, conn)
  },
  [SocketConfig.Action.Disconnect]({commit}) {
    commit(MutationType.Disconnect, undefined)
  },
  [SocketConfig.Action.Send]({state, commit, rootState}, msg: Message) {
    const text = JSON.stringify(msg);
    const socket = rootState.SocketConfigModule.conn
    socket.send(text)
  },

}

export default socketConfigActions


const messageHandler = (state: SocketConfigState, msg: Message, commit: any, rootGetters: any) => {

  const curUser: ChatRoomType = rootGetters[ChatServerId(ChatServer.Getter.User)];
  if (msg.connection_uuid != state.uuid) {
    logger.debug(`skipping msg: ${msg.action}`)
    logger.debug(`-- skipping msg FROM: `, msg.user_id)
    logger.debug(`-- skipping msg TO:`, curUser)
    return
  }

  logger.debug(`handling msg: ACTION[${msg.action}] => `, msg)
  switch (msg.action) {
    case MsgType.UpdateClient:
      logger.debug('message handler not implemented 1')
      break;
    case MsgType.UpdateChatRooms:
      const spMsgClient: MsgUpdateRooms = msg as MsgUpdateRooms
      commit(ChatServerId(ChatServer.Mutation.UpdateChatRooms), spMsgClient.chatRooms, {root: true})
      break;
    case MsgType.NotifCreateUser:
      logger.debug('message handler not implemented 2')
      break;
    case MsgType.NewConnection:
      const clientId = msg.user_id;
      commit(ChatServerId(ChatServer.Mutation.UpdateUser), clientId, {root: true})
      newLoginMsg(clientId.name)
      break;
    default:
      logger.error('unknown message type ', msg)
  }
}

function newLoginMsg(username: string) {
  void api.get('/chat').then(r => {

    logger.debug(`
- - - - - - - - - - - - - - - - - - - - - - - -
${r.data}
- - - - - - - - - - - - - - - - - - - - - - - -
user: ${username}
- - - - - - - - - - - - - - - - - - - - - - - -
        `)
  })
}
