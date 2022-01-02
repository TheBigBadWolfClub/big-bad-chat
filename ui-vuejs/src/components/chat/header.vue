<template>
  <q-item>
    <q-item-section avatar>
      <q-avatar>
        <img :src=getImgUrl(room)>
      </q-avatar>
    </q-item-section>

    <q-item-section>
      <q-item-label>{{ room.name }}</q-item-label>
      <q-item-label caption>{{ room.type }}</q-item-label>
    </q-item-section>
    <room-actions/>
  </q-item>
</template>

<script lang="ts">
import {computed, Ref} from 'vue';
import {useStore} from 'src/store';
import {ChatRoomType, ChatServer, ChatServerId} from 'src/store/ChatServer/model';
import RoomActions from 'components/chat/Actions.vue';
import {AvatarTools} from 'components/AvatarTools';

export default {
  name: 'chat-header',
  components: {RoomActions},
  setup() {
    const {getImgUrl} = AvatarTools()
    const store = useStore();
    const room: Ref<ChatRoomType> = computed(() => {
        return store.getters[ChatServerId(ChatServer.Getter.ActiveChatRoom)]
      }
    )
    return {room, getImgUrl}
  }
}

</script>

<style scoped>

</style>
