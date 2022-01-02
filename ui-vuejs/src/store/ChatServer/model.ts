export const ChatServerId = (opType: string): string => `ChatServerModule/${opType}`


export enum RoomType {
  DIRECT = 'Direct',
  MEETING = 'Meeting',
  PUBLIC = 'Public',
  UNKNOWN = 'unknown'
}

export const UndefinedChatRoom: ChatRoomType = {
  name: '',
  type: RoomType.UNKNOWN,
  uuid: ''
}

export interface User {
  uuid: string,
  name: string,
  connectionUuid: string
}

export interface ChatServerState {
  user: User,
  activeRoomId: string,
  chatRooms: ChatRoomType[],
}


export interface ChatRoomType {
  uuid: string
  name: string
  type: RoomType
}


// Mutations
export enum MutationType {
  UpdateChatRooms = 'update-chat-rooms',
  UpdateUser = 'update-socket-connection',
  SetActiveChatRoom = 'set-active-chat-room'
}

// Actions
export enum ActionType {
  ActiveChatRoom = 'set-active-chat-room'
}

// Getters
export enum GetterType {
  DirectRooms = 'get-direct-list',
  PublicRooms = 'get-public-list',
  MeetingRooms = 'get-meeting-list',
  User = 'get-user',
  ActiveChatRoom = 'get-active-chat-room'
}

//
export const ChatServer = {
  Mutation: MutationType,
  Getter: GetterType,
  Action: ActionType,
}
