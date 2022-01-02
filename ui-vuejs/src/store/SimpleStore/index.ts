import {ActionContext, ActionTree, GetterTree, Module, MutationTree} from 'vuex';
import {StateInterface} from 'src/store';


// state ///////////////////////////////////////
export interface ConfigObj {
  id: number
}

const state: ConfigObj = {
  id: 0
}

// mutation ///////////////////////////////////////
export enum MutationType {
  Update = 'UPDATE',
}

export type Mutations = {
  [MutationType.Update](currState: ConfigObj, newConfig: ConfigObj): void
}

const mutations: MutationTree<ConfigObj> & Mutations = {
  [MutationType.Update](curState, newConfig) {
    curState = newConfig
  },
}

// Actions ///////////////////////////////////////
export enum ActionTypes {
  Update = 'UPDATE',
}

export type ActionAugments = Omit<ActionContext<ConfigObj, StateInterface>, 'commit'> & {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>
}

export type Actions = {
  [ActionTypes.Update](context: ActionAugments, config: ConfigObj): void
}

const actions: ActionTree<ConfigObj, StateInterface> & Actions = {
  [ActionTypes.Update]({dispatch, state, commit, rootState, rootGetters}, config: ConfigObj) {
    commit(MutationType.Update, config)
  },
}

// getters ///////////////////////////////////////
export type Getters = {
  getObj(curState: ConfigObj): number
}

const getters: GetterTree<ConfigObj, StateInterface> & Getters = {
  getObj(curState: ConfigObj): number {
    return curState.id;
  }
}

// store ///////////////////////////////////////
const SimpleStore: Module<ConfigObj, StateInterface> = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
export default SimpleStore






