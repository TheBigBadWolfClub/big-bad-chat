// Module
import {Module} from 'vuex';
import {AppConfigState} from 'src/store/AppConfig/model';
import {StateInterface} from 'src/store';
import appConfigState from 'src/store/AppConfig/state';
import appConfigActions from 'src/store/AppConfig/actions';
import appConfigGetters from 'src/store/AppConfig/getters';
import appConfigMutations from 'src/store/AppConfig/mutations';


const AppConfigModule: Module<AppConfigState, StateInterface> = {
  namespaced: true,
  state: appConfigState,
  getters: appConfigGetters,
  mutations: appConfigMutations,
  actions: appConfigActions,
}


export default AppConfigModule
