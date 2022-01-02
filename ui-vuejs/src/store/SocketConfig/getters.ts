import {Message} from 'src/store/SocketConfig/message';
import {GetterTree} from 'vuex';
import {StateInterface} from 'src/store';
import {SocketConfig, SocketConfigState} from 'src/store/SocketConfig/model';

export type Getters = {
  [SocketConfig.Getter.getSocket](wsState: SocketConfigState): WebSocket | null
  [SocketConfig.Getter.isConnected](wsState: SocketConfigState): boolean
  [SocketConfig.Getter.stateCode](wsState: SocketConfigState): number
  [SocketConfig.Getter.onMessage](wsState: SocketConfigState): Message | undefined
}

const socketConfigGetters: GetterTree<SocketConfigState, StateInterface> & Getters = {
  [SocketConfig.Getter.getSocket](wsState: SocketConfigState): WebSocket | null {
    return wsState.conn;
  },

  [SocketConfig.Getter.isConnected](wsState: SocketConfigState): boolean {
    return wsState.isConnected
  },

  [SocketConfig.Getter.stateCode](wsState: SocketConfigState): number {
    return wsState.state
  },

  [SocketConfig.Getter.onMessage](wsState: SocketConfigState): Message | undefined {
    return wsState.message
  },

}

export default socketConfigGetters
