import {Config as NameConfig} from 'unique-names-generator/dist/unique-names-generator.constructor';
import {names, uniqueNamesGenerator} from 'unique-names-generator';
import {ChatRoomType, ChatServerState} from 'src/store/ChatServer/model';

function getRandomName(): string {
  const names = [
    'Julia', 'MarcoPolo', 'Ann', 'Joy'
  ]

  return names[Math.floor(Math.random() * names.length)]
}

function getRandomNameOld(): string {
  const config: NameConfig = {
    dictionaries: [names]
  }
  return uniqueNamesGenerator(config);
}


const chatServerState: ChatServerState = {
  user: {
    name: getRandomName(),
    uuid: '',
    connectionUuid: '',
  },
  activeRoomId: '',
  chatRooms: [] as ChatRoomType[]
}

export default chatServerState
