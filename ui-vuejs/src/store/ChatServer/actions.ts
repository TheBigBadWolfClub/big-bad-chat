import {ActionContext, ActionTree} from 'vuex';
import {StateInterface} from 'src/store';
import {Mutations} from 'src/store/ChatServer/mutations';
import {ChatServer, ChatServerState} from 'src/store/ChatServer/model';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';


export type ActionAugments = Omit<ActionContext<ChatServerState, StateInterface>, 'commit'> & {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>
}

export type Actions = {
  [ChatServer.Action.ActiveChatRoom](context: ActionAugments, uuid: string): void
}

const chatServerActions: ActionTree<ChatServerState, StateInterface> & Actions = {
  [ChatServer.Action.ActiveChatRoom]({dispatch, state, commit, rootState, rootGetters}, uuid: string) {

    // TODO unk
    rootGetters[SocketConfigId(SocketConfig.Getter.getSocket)].send(JSON.stringify('msg missing'))
  }
}


export default chatServerActions


