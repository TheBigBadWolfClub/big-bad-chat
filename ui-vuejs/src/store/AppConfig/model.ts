export const AppConfigId = (opType: string): string => `AppConfigModule/${opType}`

// State
export interface Connection {
  protocol: string
  uri: string
}

export interface ServerConfig {
  addr: string
  port: number
  rest: Connection
  webSocket: Connection
}


export interface AppConfigState {
  server: ServerConfig
}

// Mutations
enum MutationType {
  SetServerConfig = 'SET-SERVER-CONFIG',
}

// Getters
enum GetterType {
  RestUri = 'REST-URI',
  WebsocketUri = 'WEBSOCKET-URI',
  ServerConfig = 'SERVER-CONFIG'
}

// Actions
export enum ActionType {
  UpdateServer = 'UPDATE-SERVER'
}

export const AppConfig = {
  Mutation: MutationType,
  Getter: GetterType,
  Action: ActionType,
}
