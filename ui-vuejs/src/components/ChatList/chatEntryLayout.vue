<template>
  <q-item clickable @click="openRoom(model.uuid)">
    <q-item-section avatar>
      <q-avatar>
        <img :src=getImgUrl(model)>
      </q-avatar>
    </q-item-section>
    <q-item-section>
      <q-item-label>{{ model.name }}</q-item-label>
      <q-item-label caption>{{ model.type }}</q-item-label>
    </q-item-section>
  </q-item>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {ChatRoomType, ChatServer, ChatServerId} from 'src/store/ChatServer/model';
import {useStore} from 'src/store';
import {AvatarTools} from 'components/AvatarTools';

export default defineComponent({
  name: 'ChatEntryLayout',
  components: {},
  props: {
    model: {type: Object as PropType<ChatRoomType>, required: true},
  },
  setup() {
    const store = useStore();
    const {getImgUrl} = AvatarTools()


    const openRoom = (uuid: string) => {
      store.commit(ChatServerId(ChatServer.Mutation.SetActiveChatRoom), uuid)
    }

    return {getImgUrl, openRoom}
  },
})
</script>

<style lang="sass" scoped>


</style>

