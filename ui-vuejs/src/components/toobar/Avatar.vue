<template>
  <q-chip>
    <q-avatar>
      <img :src="getImage">
    </q-avatar>
    {{ username }}
  </q-chip>
</template>

<script lang="ts">

import {computed, defineComponent} from 'vue';
import {useStore} from 'src/store';
import {ChatServer, ChatServerId} from 'src/store/ChatServer/model';

export default defineComponent({
  name: 'Avatar',
  components: {},
  setup() {
    const store = useStore();
    const username = computed(() => {
      const user = store.getters[ChatServerId(ChatServer.Getter.User)];
      return user.name
    });
    const getImage = computed(() => {
      const user = store.getters[ChatServerId(ChatServer.Getter.User)];
      return `https://i.pravatar.cc/150?u=${user.uuid}`
    });
    return {username, getImage}
  }
})
</script>

<style scoped>

</style>
