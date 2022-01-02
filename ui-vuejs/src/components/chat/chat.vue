<template>
  <q-page class="column">
    <q-card v-if="room.isValid"
            class="col column bg-grey-9 full-height full-width"
            dark
            square>
      <chat-header></chat-header>
      <q-separator dark inset/>
      <chat-output></chat-output>
      <q-separator dark inset/>
      <chat-input></chat-input>

    </q-card>

    <q-card v-if="!room.isValid"
            class="col column bg-grey-9 full-height full-width">
      <q-img fit="scale-down" src="../../assets/clipart2137325.png">
        <div class=" absolute-bottom text-subtitle2 text-center">
          <h3>Best Chat App</h3>
        </div>
      </q-img>

    </q-card>

  </q-page>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent, reactive, Ref, watchEffect} from 'vue';
import {MsgChat, MsgType} from 'src/store/SocketConfig/message';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';
import ChatHeader from 'components/chat/header.vue';
import ChatOutput from 'components/chat/output.vue';
import ChatInput from 'components/chat/input.vue';
import {ChatRoomType, ChatServer, ChatServerId, RoomType} from 'src/store/ChatServer/model';
import {logger} from 'boot/logger';

interface ActiveState {
  isValid: boolean;
  room: ChatRoomType;
}

export default defineComponent({
  name: 'Chat',
  components: {ChatInput, ChatOutput, ChatHeader},
  setup() {
    const store = useStore();
    const room: Ref<ActiveState> = computed(() => {
      const room: ChatRoomType = store.getters[ChatServerId(ChatServer.Getter.ActiveChatRoom)]
      return {
        isValid: room.type != RoomType.UNKNOWN,
        room: room
      };
    })



    return {room}
  }
})
</script>

<style lang="sass">

</style>


