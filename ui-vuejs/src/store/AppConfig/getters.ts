import {AppConfig, AppConfigState, Connection, ServerConfig} from 'src/store/AppConfig/model';
import {GetterTree} from 'vuex';
import {StateInterface} from 'src/store';


export type Getters = {
  [AppConfig.Getter.RestUri](curState: AppConfigState): string
  [AppConfig.Getter.WebsocketUri](curState: AppConfigState): string
  [AppConfig.Getter.ServerConfig](curState: AppConfigState): ServerConfig
}


const appConfigGetters: GetterTree<AppConfigState, StateInterface> & Getters = {

  [AppConfig.Getter.RestUri](curState: AppConfigState): string {
    const addr = curState.server.addr
    const port = curState.server.port
    return baseEndpoint(addr, port, curState.server.rest)
  },

  [AppConfig.Getter.WebsocketUri](curState: AppConfigState): string {
    const addr = curState.server.addr
    const port = curState.server.port
    return baseEndpoint(addr, port, curState.server.webSocket)
  },

  [AppConfig.Getter.ServerConfig](curState: AppConfigState): ServerConfig {
    return curState.server
  }
}

export default appConfigGetters

const baseEndpoint = (addr: string, port: number, conn: Connection): string => {
  return `${conn.protocol}://${addr}:${port}${conn.uri}`
}
