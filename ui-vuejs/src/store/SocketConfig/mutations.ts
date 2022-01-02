import {MutationTree} from 'vuex';
import {Message} from 'src/store/SocketConfig/message';
import {SocketConfig, SocketConfigState} from 'src/store/SocketConfig/model';


export type Mutations<S = SocketConfigState> = {
  [SocketConfig.Mutation.Connect](wsState: SocketConfigState, connection: WebSocket): void
  [SocketConfig.Mutation.Disconnect](wsState: SocketConfigState): void
  [SocketConfig.Mutation.Status](wsState: SocketConfigState): void
  [SocketConfig.Mutation.Nullable](wsState: SocketConfigState): void
  [SocketConfig.Mutation.OnMessage](wsState: SocketConfigState, msg: Message): void
}

const socketConfigMutations: MutationTree<SocketConfigState> & Mutations = {
  [SocketConfig.Mutation.Connect](wsState, connection) {
    wsState.conn = connection
  },
  [SocketConfig.Mutation.Disconnect](wsState) {
    if (wsState.conn != null && wsState.conn.readyState == WebSocket.OPEN)
      wsState.conn.close(1000, `client closed the connection ${wsState.uuid}`)
  },
  [SocketConfig.Mutation.Status](wsState) {
    if (wsState.conn == null) {
      wsState.isConnected = false
      wsState.state = WebSocket.CLOSED
      return
    }
    wsState.isConnected = wsState.conn.readyState == WebSocket.OPEN
    wsState.state = wsState.conn.readyState
  },
  [SocketConfig.Mutation.Nullable](wsState) {
    wsState.conn = null
  },
  [SocketConfig.Mutation.OnMessage](wsState, msg) {
    wsState.message = msg
  }
}

export default socketConfigMutations
