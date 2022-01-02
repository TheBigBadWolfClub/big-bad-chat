<template>
  <div>
    <q-btn v-if="isConnectRef" color="green" label="Disconnect" @click="disconnect"/>
    <q-btn v-if="!isConnectRef" color="red" label="Connect" @click="connect"/>
  </div>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent, onMounted} from 'vue';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';

export default defineComponent({
  name: 'ConnButton',
  components: {},
  props: {
    title: {type: String, default: 'Socket Status'},
    active: {type: Boolean, default: true}
  },

  setup() {
    const store = useStore();

    const isConnectRef = computed(() => store.getters[SocketConfigId(SocketConfig.Getter.isConnected)])

    const connect = () => {
      void store.dispatch(SocketConfigId(SocketConfig.Action.Connect))
    }


    const disconnect = () => {
      void store.dispatch(SocketConfigId(SocketConfig.Action.Disconnect))
    }

    onMounted(() => connect())
    return {isConnectRef, connect, disconnect}
  }
})
</script>

<style lang="sass" scoped>

</style>

