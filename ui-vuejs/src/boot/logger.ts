import {boot} from 'quasar/wrappers';
import VueLogger, {Log} from 'vuejs3-logger'

// https://stackoverflow.com/questions/60634393/logging-in-vuejs-without-console-log

const isProduction = process.env.NODE_ENV === 'production';

const vlOptions = {
  isEnabled: true,
  logLevel: isProduction ? 'info' : 'debug',
  stringifyArguments: false,
  showLogLevel: true,
  showMethodName: true,
  separator: '|',
  showConsoleColors: true
};


let logger: Log;
export default boot(({app}) => {
  // Set i18n instance on app
  app.use(VueLogger, vlOptions);
  logger = app.config.globalProperties.$log;
  // ^ ^ ^ this will allow you to use this.$log (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});


export {logger, VueLogger, vlOptions};
