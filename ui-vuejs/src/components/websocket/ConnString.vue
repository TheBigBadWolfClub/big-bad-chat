<template>
  <q-banner :class="{'bg-red': !isConnectRef, 'bg-green': isConnectRef}"
            class="text-white"
            dense inline-actions>
    Websocket Connection to server: {{ serverUri }}
  </q-banner>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent} from 'vue';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';
import {AppConfig, AppConfigId} from 'src/store/AppConfig/model';

export default defineComponent({
  name: 'ConnString',
  setup() {
    const store = useStore();

    const isConnectRef = computed(() => store.getters[SocketConfigId(SocketConfig.Getter.isConnected)])
    const serverUri = computed(() => store.getters[AppConfigId(AppConfig.Getter.WebsocketUri)])
    return {isConnectRef, serverUri}
  }
})
</script>

<style lang="sass" scoped>

</style>

