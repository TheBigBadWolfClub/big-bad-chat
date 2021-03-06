import {store} from 'quasar/wrappers'
import {InjectionKey} from 'vue'
import {createStore, Store as VuexStore, useStore as vuexUseStore,} from 'vuex'
import SocketConfigModule from './SocketConfig';
import ChatServerModule from 'src/store/ChatServer';
import AppConfigModule from 'src/store/AppConfig';

// import example from './module-example'
// import { ExampleStateInterface } from './module-example/state';

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Store instance.
 */

export interface StateInterface {
  // Define your own store structure, using submodules if needed
  // example: ExampleStateInterface;
  // Declared as unknown to avoid linting issue. Best to strongly type as per the line above.
  SocketConfigModule: any;
  ChatServerModule: any;
  AppConfigModule: any;
}

// provide typings for `this.$store`
declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $store: VuexStore<StateInterface>
  }
}

// provide typings for `useStore` helper
export const storeKey: InjectionKey<VuexStore<StateInterface>> = Symbol('vuex-key')

export default store(function (/* { ssrContext } */) {
  return createStore<StateInterface>({
    modules: {
      ChatServerModule: ChatServerModule,
      AppConfigModule: AppConfigModule,
      SocketConfigModule: SocketConfigModule,
    },

    // enable strict mode (adds overhead!)
    // for dev mode and --debug builds only
    strict: !!process.env.DEBUGGING
  });
})

export function useStore() {
  return vuexUseStore(storeKey)
}
