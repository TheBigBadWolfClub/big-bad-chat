import {AppConfigState} from 'src/store/AppConfig/model';


const appConfigState: AppConfigState = {
  server: {
    addr: '127.0.0.1',
    port: 8080,
    rest: {
      protocol: 'http',
      uri: '/api'
    },
    webSocket: {
      protocol: 'ws',
      uri: '/ws'
    }
  }
}


export default appConfigState
