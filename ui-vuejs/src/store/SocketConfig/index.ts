// store ///////////////////////////////////////
import {Module} from 'vuex';
import {StateInterface} from 'src/store';
import {SocketConfigState} from 'src/store/SocketConfig/model';
import socketConfigState from 'src/store/SocketConfig/state';
import socketConfigGetters from 'src/store/SocketConfig/getters';
import socketConfigMutations from 'src/store/SocketConfig/mutations';
import socketConfigActions from 'src/store/SocketConfig/actions';

const SocketConfigModule: Module<SocketConfigState, StateInterface> = {
  namespaced: true,
  state: socketConfigState,
  getters: socketConfigGetters,
  mutations: socketConfigMutations,
  actions: socketConfigActions,
}
export default SocketConfigModule




