<template>
  <q-card-section class=" bg-grey-9  full-width">
    <q-form class=" column" @submit="onSubmit">
      <q-input
        v-model="state.message"
        :rows="1"
        autogrow
        bg-color="grey-3"
        color="grey-10"
        dense
        input-style="max-height: 6em"
        outlined
        type="textarea"
      >

        <template v-slot:after>
          <q-btn
            color="brown-5"
            dense
            icon="send"
            square
            type="submit"
            @click="onSubmit"
          />
        </template>
      </q-input>
    </q-form>
  </q-card-section>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent, reactive, Ref} from 'vue';
import {ChatMessage, MsgChat, MsgType} from 'src/store/SocketConfig/message';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';
import {ChatRoomType, ChatServer, ChatServerId, User} from 'src/store/ChatServer/model';
import {uuid} from 'vue-uuid';
import {logger} from 'boot/logger';

export default defineComponent({
  name: 'chat-input',
  components: {},
  setup() {
    const store = useStore();
    const state = reactive({message: ''})


    const room: Ref<ChatRoomType> = computed(() => store.getters[ChatServerId(ChatServer.Getter.ActiveChatRoom)])
    const user: Ref<User> = computed(() => store.getters[ChatServerId(ChatServer.Getter.User)])

    const onSubmit = () => {
      const userId = {
        name: user.value.name,
        uuid: user.value.uuid,
        connection_uuid: user.value.connectionUuid,
      }
      const chat: ChatMessage = {
        uuid: uuid.v4(),
        room_uuid: room.value.uuid,
        sender: userId,
        text: [state.message],  // TODO unk
        time: new Date().toISOString(),
      };

      const msg: MsgChat = {
        uuid: uuid.v4(),
        connection_uuid: user.value.connectionUuid,
        user_id: userId,
        action: MsgType.NewTextAction,
        chat: chat
      };

      store.getters[SocketConfigId(SocketConfig.Getter.getSocket)].send(JSON.stringify(msg))
      state.message = '';

      logger.debug(msg)
    }

    return {state, onSubmit}
  }
})
</script>


