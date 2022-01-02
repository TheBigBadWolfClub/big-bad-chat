import {ChatRoomType} from 'src/store/ChatServer/model';

export enum MsgType {
  NewConnection = 'new-connection',
  NotifCreateUser = 'notif-create-user',
  UpdateChatRooms = 'update-chat-rooms',

  NewTextAction = 'new-text',
  JoinRoomAction = 'join-room',
  LeaveRoomAction = 'leave-room',

  // TODO unk no match
  RoomChange = 'room-change',
  UpdateClient = 'update-client',
}


export interface UserId {
  name: string
  uuid: string
  connection_uuid: string,
}


export interface Message {
  uuid: string,
  connection_uuid: string,
  action: MsgType,
  user_id: UserId,
}


export interface MsgUpdateRooms extends Message {
  chatRooms: ChatRoomType[]
}

export interface MsgChat extends Message {
  chat: ChatMessage
}

export interface ChatMessage {
  uuid: string,
  time: string,
  text: string[],
  sender: UserId,
  room_uuid: string,
}


export const joinRoomMsgString = (roomId: string): string => {
  const msg = {
    action: MsgType.JoinRoomAction,
    target: roomId
  }

  return JSON.stringify(msg);

}
