import {ChatRoomType, ChatServer, ChatServerState, RoomType, UndefinedChatRoom, User} from 'src/store/ChatServer/model';
import {GetterTree} from 'vuex';
import {StateInterface} from 'src/store';

export type Getters = {
  [ChatServer.Getter.DirectRooms](curState: ChatServerState): ChatRoomType[]
  [ChatServer.Getter.PublicRooms](curState: ChatServerState): ChatRoomType[]
  [ChatServer.Getter.MeetingRooms](curState: ChatServerState): ChatRoomType[]
  [ChatServer.Getter.User](curState: ChatServerState): User
  [ChatServer.Getter.ActiveChatRoom](curState: ChatServerState): ChatRoomType
}

const chatServerGetters: GetterTree<ChatServerState, StateInterface> & Getters = {
  [ChatServer.Getter.DirectRooms](curState: ChatServerState): ChatRoomType[] {
    return curState.chatRooms.filter(r => r.type == RoomType.DIRECT && r.uuid != curState.user.uuid);
  },
  [ChatServer.Getter.PublicRooms](curState: ChatServerState): ChatRoomType[] {
    return curState.chatRooms.filter(r => r.type == RoomType.PUBLIC);
  },
  [ChatServer.Getter.MeetingRooms](curState: ChatServerState): ChatRoomType[] {
    return curState.chatRooms.filter(r => r.type == RoomType.MEETING);
  },
  [ChatServer.Getter.User](curState: ChatServerState): User {
    return curState.user
  },
  [ChatServer.Getter.ActiveChatRoom](curState: ChatServerState): ChatRoomType {
    const room = curState.chatRooms.find(r => r.uuid === curState.activeRoomId);
    if (curState.activeRoomId === '' || room === undefined) {
      return UndefinedChatRoom
    }
    return room
  },
}

export default chatServerGetters
