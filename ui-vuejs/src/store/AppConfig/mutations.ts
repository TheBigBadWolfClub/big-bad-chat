import {MutationTree} from 'vuex';
import {AppConfig, AppConfigState, ServerConfig} from 'src/store/AppConfig/model';


export type Mutations<S = AppConfigState> = {
  [AppConfig.Mutation.SetServerConfig](currState: AppConfigState, config: ServerConfig): void
}


const appConfigMutations: MutationTree<AppConfigState> & Mutations = {
  [AppConfig.Mutation.SetServerConfig](curState: AppConfigState, config: ServerConfig) {
    curState.server = config
  }
}

export default appConfigMutations


