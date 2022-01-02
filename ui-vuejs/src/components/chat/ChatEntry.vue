<template>
  <q-chat-message v-if="isDirect"
                  :sent="isSender"
                  :stamp="timestampStr()"
                  :text="message.text"
                  bg-color="amber-7"
  />

  <q-chat-message v-if="!isDirect"
                  :name="message.sender.name"
                  :sent="isSender"
                  :stamp="timestampStr()"
                  :text="message.text"
                  bg-color="amber-7"
  />
</template>

<script lang="ts">


import {defineComponent, PropType, ref} from 'vue';
import {ChatRoomType, RoomType, User} from 'src/store/ChatServer/model';
import {ChatMessage} from 'src/store/SocketConfig/message';
import {logger} from 'boot/logger';


export default defineComponent({
  name: 'ChatEntry',
  props: {
    message: {type: Object as PropType<ChatMessage>, required: true},
    room: {type: Object as PropType<ChatRoomType>, required: true},
    user: {type: Object as PropType<User>, required: true}
  },
  setup(props) {
    const isDirect = ref(props.room.type === RoomType.DIRECT)
    const isSender = props.message.sender.uuid === props.user.uuid

    logger.debug('===>', isSender)
    logger.debug('===>', props.message.sender.uuid)
    logger.debug('===>', props.user.uuid)
    logger.debug('===>', props)
    const timestampStr = () => {
      return '7 min ago'
    }

    return {isDirect, isSender, timestampStr}
  }
})
</script>

<style scoped>

</style>
