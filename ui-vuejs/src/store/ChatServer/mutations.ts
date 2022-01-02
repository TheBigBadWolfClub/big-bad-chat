// mutation ///////////////////////////////////////
import {ChatRoomType, ChatServer, ChatServerState, User} from 'src/store/ChatServer/model';
import {MutationTree} from 'vuex';


export type Mutations = {
  [ChatServer.Mutation.UpdateChatRooms](curState: ChatServerState, update: ChatRoomType[]): void
  [ChatServer.Mutation.UpdateUser](curState: ChatServerState, update: User): void
  [ChatServer.Mutation.SetActiveChatRoom](curState: ChatServerState, uuid: string): void
}

const chatServerMutations: MutationTree<ChatServerState> & Mutations = {
  [ChatServer.Mutation.UpdateChatRooms](curState: ChatServerState, update: ChatRoomType[]) {
    curState.chatRooms = curState.chatRooms
      .filter(it => update.find(b => it.uuid == b.uuid) == undefined)
      .concat(update)
      .sort(sortByName)
  },
  [ChatServer.Mutation.UpdateUser](curState: ChatServerState, update: User) {
    curState.user = update
  },
  [ChatServer.Mutation.SetActiveChatRoom](curState: ChatServerState, uuid: string) {
    curState.activeRoomId = uuid
  }
}

const sortByName = (str1: any, str2: any) => {
  if (str1.name < str2.name) {
    return -1;
  }
  if (str1.name > str2.name) {
    return 1;
  }
  return 0
}

export default chatServerMutations
