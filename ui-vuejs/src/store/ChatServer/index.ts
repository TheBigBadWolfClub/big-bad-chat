import {Module} from 'vuex';
import {StateInterface} from 'src/store';
import {ChatServerState} from 'src/store/ChatServer/model';
import chatServerState from 'src/store/ChatServer/state';
import chatServerGetters from 'src/store/ChatServer/getters';
import chatServerMutations from 'src/store/ChatServer/mutations';
import chatServerActions from 'src/store/ChatServer/actions';


const ChatServerModule: Module<ChatServerState, StateInterface> = {
  namespaced: true,
  state: chatServerState,
  getters: chatServerGetters,
  mutations: chatServerMutations,
  actions: chatServerActions
}
export default ChatServerModule




