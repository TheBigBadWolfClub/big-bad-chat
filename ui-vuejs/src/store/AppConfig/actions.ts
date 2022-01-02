import {AppConfig, AppConfigState, ServerConfig} from 'src/store/AppConfig/model';
import {ActionContext, ActionTree} from 'vuex';
import {StateInterface} from 'src/store';
import {Mutations} from 'src/store/AppConfig/mutations';


export type ActionAugments = Omit<ActionContext<AppConfigState, StateInterface>, 'commit'> & {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>
}

export type Actions = {
  [AppConfig.Action.UpdateServer](context: ActionAugments, config: ServerConfig): void
}

const appConfigActions: ActionTree<AppConfigState, StateInterface> & Actions = {
  [AppConfig.Action.UpdateServer]({state, commit, rootState}, config: ServerConfig) {
    void commit(AppConfig.Mutation.SetServerConfig, config)
  },
}

export default appConfigActions
