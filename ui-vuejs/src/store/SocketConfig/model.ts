import {Message} from 'src/store/SocketConfig/message';

export const SocketConfigId = (opType: string): string => `SocketConfigModule/${opType}`


// State
export interface SocketConfigState {
  uuid: string
  conn: WebSocket | null
  isConnected: boolean
  state: number
  message: Message
}


// Mutations
export enum MutationType {
  Connect = 'CONNECT',
  Disconnect = 'DISCONNECT',
  Status = 'STATUS',
  Nullable = 'NULLABLE',
  OnMessage = 'ON_MESSAGE',
}

// Getters
enum GetterType {
  getSocket = 'GET-SOCKET',
  isConnected = 'IS-CONNECTED',
  stateCode = 'STATE-CODE',
  onMessage = 'ON-MESSAGE',
}


// Actions
export enum ActionType {
  Connect = 'CONNECT',
  Disconnect = 'DISCONNECT',
  Send = 'SEND'
}


export const SocketConfig = {
  Mutation: MutationType,
  Getter: GetterType,
  Action: ActionType,
}
