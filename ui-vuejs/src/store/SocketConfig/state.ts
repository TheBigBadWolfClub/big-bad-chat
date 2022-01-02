import {Message} from 'src/store/SocketConfig/message';
import {SocketConfigState} from 'src/store/SocketConfig/model';
import {uuid} from 'vue-uuid';


const socketConfigState: SocketConfigState = {
  uuid: uuid.v4(),
  conn: null,
  isConnected: false,
  state: WebSocket.CLOSED,
  message: {} as Message
}


export default socketConfigState
